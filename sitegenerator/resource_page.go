package sitegenerator

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

type ResourcePage struct {
	templateBasePath string
	resource         Resource
}

func NewResourcePage(templateBasePath string, resource Resource) ResourcePage {
	return ResourcePage{templateBasePath, resource}
}

func (i *ResourcePage) Generate(w io.Writer) error {
	var tmpl = template.Must(template.ParseFiles(path.Join(i.templateBasePath, "templates/resource.html")))
	err := tmpl.Execute(w, i.resource)
	if err != nil {
		return fmt.Errorf("cannot write resource %s: %s", i.resource.Repository, err)
	}
	return nil
}
