# DutyFree
Duty Free is a website for showcasing re-usable components of Concourse. With DutyFree, Concourse users can explore and discover resources written by the community 

## Duty Free website

- Duty Free is hosted on github pages, https://concourse.github.io/dutyfree
- To generate locally the website execute the main file passing a destination folder and a [`resources.yml`](https://github.com/concourse/dutyfree/blob/master/resources.yml) as argument. eg `go run main.go /tmp/folder resources.yml`
- a Concourse pipeline is available in the [`ci` folder](https://github.com/concourse/dutyfree/tree/master/ci), it requires 
  - ssh key (as `ci.private`) for pushing to the `gh-pages` branch
  - Github token (as `github_token`) to retrieve the readme of the resources from Github 
  - slack hook, to notify about a new push deployment of the website

## Run a development server

`go run dev/dev.go`
mock resources are defined in [`dev.go`](https://github.com/concourse/dutyfree/blob/master/dev/dev.go) and it uses a mock server instead of connecting to Github to download the readmes.

## Tests

`ginkgo -r` or `go test ./...`
