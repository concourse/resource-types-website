package persistence

import (
	"fmt"
	"strconv"
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

			currResource.NameWithOwner = owner + "/" + repoName
			resourcesMap[currResource.NameWithOwner] = 0

			fs.resources = append(fs.resources, currResource)
		}
	}

	//TODO: doesn't look like the best way is to call github gql here :thinking:
	resourcesMap, err = fs.GhGqlWrapper.GetStars(resourcesMap)
	if err != nil {
		return err
	}

	for i, curr := range fs.resources {
		fs.resources[i].StarsCount = resourcesMap[curr.NameWithOwner]
		fs.resources[i].Stars = formatStars(fs.resources[i].StarsCount)
	}

	return nil
}

func formatStars(stars int) string {
	switch {
	case stars < 1000:
		return strconv.Itoa(stars)
	case stars >= 1000:
		displayStars := float32(stars) / 1000
		ghstars := fmt.Sprintf("%.1fk", displayStars)
		return strings.Replace(ghstars, ".0", "", 1)
	case stars >= 100000:
		return strconv.Itoa(stars/1000) + "k"
	default:
		return ""
	}
}

func extractOwnerAndRepo(url string) (string, string) {
	urlNoGit := strings.Replace(url, ".git", "", 1)
	parts := strings.Split(urlNoGit, "/")
	return parts[len(parts)-2], parts[len(parts)-1]
}
