package sitegenerator

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

type IndexPage struct {
	templateBasePath string
	ResourceModels   []ResourceModel
}

func NewIndexPage(templateBasePath string, resourceModels []ResourceModel) IndexPage {
	return IndexPage{templateBasePath, resourceModels}
}

func (i *IndexPage) Generate(w io.Writer) error {
	var tmpl = template.Must(template.ParseFiles(path.Join(i.templateBasePath, "templates/index.html")))
	err := tmpl.Execute(w, i.ResourceModels)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	return nil
}
