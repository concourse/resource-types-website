package persistence

import (
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"strings"

	"github.com/concourse/dutyfree/resource"
	"gopkg.in/yaml.v2"
)

type Filesystem struct {
	Location  string
	resources []resource.Resource
}

func (fs *Filesystem) GetAllResources() []resource.Resource {
	return fs.resources
}

func (fs *Filesystem) LoadResources() error {
	//_, err := ioutil.ReadDir(fs.Location)
	//
	//if err != nil {
	//	return err
	//}

	filesNew := packr.NewBox(fs.Location)

	err := filesNew.Walk(func(s string, f packd.File) error {
		inf, err := f.FileInfo()
		if err != nil {
			return err
		}
		if !inf.IsDir() && strings.Contains(inf.Name(), ".yml") {
			fileBytes, err := ioutil.ReadAll(f)
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
		return nil

	})
	//if err != nil {
	//	return err
	//}
	//for _, currFile := range filesNew. {
	//	if !file.IsDir() && strings.Contains(file.Name(), ".yml") {
	//		fileBytes, err := ioutil.ReadFile(fs.Location + "/" + file.Name())
	//		if err != nil {
	//			return err
	//		}
	//		var currResource resource.Resource
	//		//fmt.Println("parsing: " + file.Name())
	//		err = yaml.UnmarshalStrict(fileBytes, &currResource)
	//		if err != nil {
	//			return err
	//		}
	//		fs.resources = append(fs.resources, currResource)
	//	}
	//}
	return err
}
