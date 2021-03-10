FROM concourse/golang-builder as builder
COPY . /src
RUN apt-get update \
     && apt-get install -y --no-install-recommends curl ca-certificates apt-transport-https software-properties-common gpg-agent

# install Node 12.x
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get update && apt-get install -y nodejs

# install Yarn for web UI tests
RUN curl -fsSL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN add-apt-repository "deb https://dl.yarnpkg.com/debian/ stable main"
RUN apt-get update && apt-get -y install yarn

WORKDIR /src/web
RUN yarn install && yarn build

WORKDIR /src/warehouse
ENV CGO_ENABLED 0
RUN go build ./main.go

FROM ubuntu:bionic AS dutyfree
EXPOSE 9090
COPY --from=builder src/warehouse/dutyfree /usr/local/bin/
RUN apt-get update \
      && apt-get install -y --no-install-recommends ca-certificates
RUN chmod +x /usr/local/bin/dutyfree

FROM dutyfree
ENTRYPOINT ["dutyfree"]
