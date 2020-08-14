package persistence

import (
	"fmt"
	"net/url"
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

	resourcesMap := make(map[string]bool)

	for _, fileBytes := range files {
		if strings.Contains(fileBytes.Name, ".yml") || strings.Contains(fileBytes.Name, ".yaml") {
			var currResource resource.Resource
			err = yaml.UnmarshalStrict(fileBytes.Contents, &currResource)
			//TODO: do we exit if one file is corrupt of just skip it??
			if err != nil {
				return err
			}

			parsedRepo, err := parseRepoURL(currResource.URL)
			if err != nil {
				return err
			}

			if len(parsedRepo.Owner) > 0 {
				currResource.Owner = "@"
			}
			currResource.Owner += parsedRepo.Owner

			currResource.Host = parsedRepo.Host

			currResource.NameWithOwner = parsedRepo.Owner + "/" + parsedRepo.Name
			resourcesMap[currResource.NameWithOwner] = parsedRepo.IsGithub

			fs.resources = append(fs.resources, currResource)
		}
	}

	//TODO: doesn't look like the best way is to call github gql here :thinking:
	resourcesStars, err := fs.GhGqlWrapper.GetStars(resourcesMap)
	if err != nil {
		return err
	}

	for i, curr := range fs.resources {
		fs.resources[i].StarsCount = resourcesStars[curr.NameWithOwner]
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

type Repo struct {
	Host     string
	Owner    string
	Name     string
	IsGithub bool
}

func parseRepoURL(sourceURL string) (Repo, error) {
	var parsedRepo Repo
	u, err := url.Parse(strings.ToLower(sourceURL))
	if err != nil {
		return Repo{}, err
	}
	parts := strings.Split(u.Path, "/")

	if strings.Contains(u.Host, "github") {
		parsedRepo.IsGithub = true
	}

	parsedRepo.Host = strings.Split(u.Host, ".")[0]

	if len(parts) == 2 {
		parsedRepo.Name = parts[1]
		parsedRepo.Owner = ""
		return parsedRepo, nil
	}

	parsedRepo.Name = parts[2]
	parsedRepo.Owner = parts[1]
	return parsedRepo, nil
}
