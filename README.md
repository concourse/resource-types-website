# Resource Types Website for Concourse (aka Dutyfree)
Welcome to the development repository for the resource types website. Here you can find the source code for the website as well as different options for packaging. The website is currently linked from the Concourse docs website (resource types tab) and it is a [standalone website](https://resource-types.concourse-ci.org).

## Architecture:

- The code is mainly a Golang back-end with an elm front-end.
- We also rely on the `resource-types` directory as the persistence of the website. To fetch the contents of the `resource-types` directory, run:

```
git submodule init
git submodule update
```

## Run Locally:
There are multiple ways to run the code locally:

- to run directly from the code, you can run the following commands:
  - if you want to use a different port, the app uses the environment variable `PORT` to select its port or else it will use port `9090`.
  ```bash
  cd warehouse/web
  yarn install && yarn build
  cd ..
  go run ./main.go
  ```
- To run using docker:
  ```bash
  docker build -t dutyfree
  docker run dutyfree -p 9090:9090
  ```
- To run using docker compose: `docker-compose up --build` that will build the image locally and run the website on port `9090`.

### Notes:

- Using the previous commands you will be able to access the application through: `http://localhost:9090`.
- To use a different set of resource types, you can change the resource types under the directory `resource-types`, this would allow running different versions of the website.

## Tests:

- To test the back-end server:
  ```bash
  cd warehouse

  go install github.com/onsi/ginkgo/ginkgo

  ginkgo -r -keepGoing .
  ```

- To test the front-end code: `cd warehouse && yarn install && yarn test`.
