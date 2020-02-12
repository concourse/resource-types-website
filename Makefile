DEPLOY_DATE = $(shell date +%s)

Default:
	@echo "Usage: make [TASK]";
	@echo "Available tasks:";
	@echo "update-resources: \tto update the local version of the resource-types (i.e. clone the latest https://github.com/concourse/resource-types)"
	@echo "build-docker: \t\tbuild the Dockerfile locally and tag it concourse/dutyfree"
	@echo "publish-docker: \tbuilds concourse/dutyfree and then publishes it, needs appropriate permissions to be able to push"
	@echo "helm-diff: \t\truns a helm diff between the deployment and the local files"
	@echo "helm-deploy: \t\truns a helm diff, then attempts to deploy the local chart"

update-resources:
	git submodule update
	  cd resource-types && \
	  git checkout master && \
	  git pull
	$
	git add resource-types && \
	git commit -m "update resource-types @"$(shell git --git-dir ./resource-types/.git log --format=format:%H -1 --pretty=format:%h)
	@echo "\n\n\n\n\n"
	@echo "*****************************************"
	@echo "* Resource Types are now at latest,\t*\n* Please remember to push to a branch.\t*"
	@echo "*****************************************"


build-docker:
	docker build . -t concourse/dutyfree


publish-docker: | build-docker
	docker push concourse/dutyfree


helm-deploy: | helm-diff
	cd dutyfree-chart && \
	  helm upgrade \
	    --wait \
	    --install \
	    --namespace=dutyfree \
	    --set=annotations.rollingUpdate=\"$(DEPLOY_DATE)\" \
	    dutyfree \
	    .

	kubectl \
	  --namespace "dutyfree" \
	  rollout status deployment \
	  "dutyfree"


helm-diff:
	cd dutyfree-chart && \
	  helm diff \
	    upgrade \
	    --namespace=dutyfree \
	    --set=annotations.rollingUpdate=\"$(DEPLOY_DATE)\" \
	    dutyfree .
