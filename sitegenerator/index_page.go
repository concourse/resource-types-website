package sitegenerator

import (
	"fmt"
	"io"
)

type IndexPage struct {
	ResourceModels []ResourceModel
	Path           []string
	CategoryList   []string
}

var IndexPagePath = []string{"All Resources"}

func NewIndexPage(resourceModels []ResourceModel) IndexPage {
	return IndexPage{resourceModels, IndexPagePath, createCategoryList(resourceModels)}
}

func (i *IndexPage) Generate(w io.Writer) error {
	err := load("index.html").ExecuteTemplate(w, "index.html", i)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	return nil
}

func createCategoryList(resources []ResourceModel) []string {
	var categoryList []string
	for _, resource := range resources {
		categoryList = append(categoryList, resource.Categories...)
	}

	foundCategories := make(map[string]bool)
	var uniqueCategoryList []string

	for _, category := range categoryList {
		found := foundCategories[category]
		if !found {
			foundCategories[category] = true
			uniqueCategoryList = append(uniqueCategoryList, category)
		}
	}

	return uniqueCategoryList
}
