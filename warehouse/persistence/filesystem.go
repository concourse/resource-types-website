package persistence

import (
	"github.com/concourse/dutyfree/resource"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Filesystem struct {
	Location  string
	resources []resource.Resource
}

func (fs *Filesystem) GetAllResources() []resource.Resource {
	return fs.resources
}

func (fs *Filesystem) LoadResources() error {
	files, err := ioutil.ReadDir(fs.Location)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileBytes, err := ioutil.ReadFile(fs.Location + "/" + file.Name())
			if err != nil {
				return err
			}
			var currResource resource.Resource
			err = yaml.UnmarshalStrict(fileBytes, &currResource)
			if err != nil {
				return err
			}
			fs.resources = append(fs.resources, currResource)
		}
	}
	return nil
}
