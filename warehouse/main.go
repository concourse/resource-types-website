package main

import (
	"fmt"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/server"
	"github.com/gobuffalo/packr/v2"
)

func main() {

	publicFetcher := fetcher.Fetcher{Box: *packr.New("publicBox", "../web/public")}
	resourcesFetcher := fetcher.Fetcher{Box: *packr.New("resourcesBox", "../resource-types")}

	s := server.Server{
		Port:                     9090,
		PublicFilesFetcher:       publicFetcher,
		ResourceTypesFileFetcher: resourcesFetcher,
	}

	s.Start()
	fmt.Println("Dutyfree server started")
	<-s.Exited
}
