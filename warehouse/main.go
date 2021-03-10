package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/concourse/dutyfree/githubwrapper"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/server"
)

func main() {

	publicFetcher := fetcher.Fetcher{Box: os.DirFS("../web/public")}
	resourcesFetcher := fetcher.Fetcher{Box: os.DirFS("../resource-types")}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil || port == 0 {
		port = 9090
	}

	token := os.Getenv("GH_TOKEN")
	if token == "" {
		panic("GH_TOKEN environment variable is not set")
	}

	ghURL := os.Getenv("GH_URL")
	if ghURL == "" {
		ghURL = "https://api.github.com/graphql"
	}

	s := server.Server{
		Port:                     port,
		PublicFilesFetcher:       publicFetcher,
		ResourceTypesFileFetcher: resourcesFetcher,
		GithubGraphqlWrapper:     githubwrapper.NewWrapper(ghURL, token),
	}

	s.Start()
	fmt.Println("Dutyfree server started on port " + strconv.Itoa(port))
	<-s.Exited
}
