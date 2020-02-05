package persistence

import (
	"strings"

	"github.com/concourse/dutyfree/githubwrapper"
	"gopkg.in/yaml.v2"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/resource"
)

type Filesystem struct {
	Fetcher      fetcher.Fetcher
	GhGqlWrapper githubwrapper.Wrapper
	resources    []resource.Resource
}

func (fs *Filesystem) GetAllResources() []resource.Resource {
	return fs.resources
}

func (fs *Filesystem) LoadResources() error {
	files, err := fs.Fetcher.GetAll()
	if err != nil {
		return err
	}

	resourcesMap := make(map[string]int)

	for _, fileBytes := range files {
		if strings.Contains(fileBytes.Name, ".yml") {
			var currResource resource.Resource
			err = yaml.UnmarshalStrict(fileBytes.Contents, &currResource)
			//TODO: do we exit if one file is corrupt of just skip it??
			if err != nil {
				return err
			}

			var repoName, owner string
			owner, repoName = extractOwnerAndRepo(currResource.URL)

			currResource.Owner = "@" + owner
			fs.resources = append(fs.resources, currResource)

			currResource.NameWithOwner = owner + "/" + repoName
			resourcesMap[currResource.NameWithOwner] = 0
		}
	}

	//TODO: doesn't look like the best way is to call github gql here :thinking:
	resourcesMap, err = fs.GhGqlWrapper.GetStars(resourcesMap)
	if err != nil {
		return err
	}

	for i, curr := range fs.resources {
		fs.resources[i].Stars = resourcesMap[curr.NameWithOwner]
	}

	return nil
}

func extractOwnerAndRepo(url string) (string, string) {
	parts := strings.Split(url, "/")
	return parts[len(parts)-2], parts[len(parts)-1]
}
