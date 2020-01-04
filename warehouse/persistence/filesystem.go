package persistence

import (
	"gopkg.in/yaml.v2"
	"strings"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/resource"
)

type Filesystem struct {
	Fetcher   fetcher.Fetcher
	resources []resource.Resource
}

func (fs *Filesystem) GetAllResources() []resource.Resource {
	return fs.resources
}

func (fs *Filesystem) LoadResources() error {
	files, err := fs.Fetcher.GetAll()
	if err != nil {
		return err
	}
	for _, fileBytes := range files {
		if strings.Contains(fileBytes.Name, ".yml") {
			var currResource resource.Resource
			err = yaml.UnmarshalStrict(fileBytes.Contents, &currResource)
			//TODO: do we exit if one file is corrupt of just skip it??
			if err != nil {
				return err
			}
			fs.resources = append(fs.resources, currResource)
		}
	}
	return nil
}
