package sitegenerator

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

type IndexPage struct {
	templateBasePath string
	Resources        []Resource
}

func NewIndexPage(templateBasePath string, resources []Resource) IndexPage {
	return IndexPage{templateBasePath, resources}
}

func (i *IndexPage) Generate(w io.Writer) error {
	var tmpl = template.Must(template.ParseFiles(path.Join(i.templateBasePath, "templates/index.html")))
	err := tmpl.Execute(w, i.Resources)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	return nil
}
