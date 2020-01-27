DEPLOY_DATE = $(shell date +%s)

update-resources:
	git submodule update
	  cd resource-types && \
	  git checkout master && \
	  git pull
	git add resource-types && \
	git ci -m "update resource-types"
	echo "resource types are now at latest, \
	  Please remember to push to a branch"


build-docker:
	docker build . -t concourse/dutyfree


publish-docker: | build-docker
	docker push concourse/dutyfree


helm-deploy: | helm-diff
	cd dutyfree-chart && \
	  helm upgrade \
	    --install \
	    --namespace=dutyfree \
	    --set=annotations.rollingUpdate=\"$(DEPLOY_DATE)\" \
	    dutyfree \
	    .
	kubectl get --namespace dutyfree pods -w


helm-diff:
	cd dutyfree-chart && \
	  helm diff \
	    upgrade \
	    --namespace=dutyfree \
	    --set=annotations.rollingUpdate=\"$(DEPLOY_DATE)\" \
	    dutyfree .
