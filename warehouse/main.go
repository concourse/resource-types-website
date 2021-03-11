package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strconv"

	"github.com/concourse/dutyfree/githubwrapper"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/server"
)

//go:embed web/public
var webFS embed.FS

//go:embed resource-types
var resourceTypesFS embed.FS

func main() {

	webFS, err := fs.Sub(webFS, "web/public")
	if err != nil {
		panic(err)
	}
	resourceTypesFS, err := fs.Sub(resourceTypesFS, "resource-types")
	if err != nil {
		panic(err)
	}

	publicFetcher := fetcher.Fetcher{Box: webFS}
	resourcesFetcher := fetcher.Fetcher{Box: resourceTypesFS}

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
