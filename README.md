# DutyFree
Duty Free is a website for showcasing re-usable components of Concourse. With DutyFree, Concourse users can explore and discover resources written by the community 

## Duty Free website

- Duty Free is hosted on github pages https://concourse.github.io/dutyfree
- To generate the Duty Free html files locally, the main.go file is executed, passing a destination folder for the html files and [`resources.yml`](https://github.com/concourse/dutyfree/blob/master/resources.yml) as arguments, e.g. `go run main.go /tmp/folder resources.yml`
- a Concourse pipeline for running tests and deploying to Github pages is available in the [`ci` folder](https://github.com/concourse/dutyfree/tree/master/ci); it requires 
  - ssh key (as `ci.private`) for pushing to the `gh-pages` branch
  - Github token (as `github_token`) to retrieve the ReadMe of the resources from Github 
  - Slack hook, to notify about a new push deployment of the website
- The website is rebuilt every day (9am UTC) to refresh the content of ReadMe of the resources

## Run a development server

`go run dev/dev.go`

Mock resources are defined in [`dev.go`](https://github.com/concourse/dutyfree/blob/master/dev/dev.go). A mock server is used instead of connecting to Github to download the ReadMes.

## Tests

`ginkgo -r` or `go test ./...`
