package sitegenerator

import (
	"fmt"
	"html/template"
	"strings"
)

type Resource struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

type ResourceModel struct {
	Resource
	Identifier        string
	AuthorHandle      string
	AuthorProfileLink string
	Readme            template.HTML
}

func Enrich(resources []Resource) ([]ResourceModel, error) {
	var resourceModels []ResourceModel

	for _, resource := range resources {
		resourceModel := ResourceModel{Resource: resource}

		// Here happens the Enrichment
		segmentsAll := strings.Split(resource.Repository, "/")

		if len(segmentsAll) < 5 || segmentsAll[0] != "https:" || segmentsAll[2] != "github.com" {
			return resourceModels, fmt.Errorf("invalid repository for the resource (%s)", resource.Name)
		}

		segments := segmentsAll[3:]

		resourceModel.Identifier = strings.Join(segments, "-")
		resourceModel.AuthorHandle = segmentsAll[3]
		resourceModel.AuthorProfileLink = strings.Join(segmentsAll[:4], "/")

		resourceModels = append(resourceModels, resourceModel)
	}
	return resourceModels, nil
}
