package sitegenerator

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

type ResourcePage struct {
	ResourceModel
	templateBasePath string
}

func NewResourcePage(templateBasePath string, resourceModel ResourceModel) ResourcePage {
	return ResourcePage{resourceModel, templateBasePath}
}

func (rp *ResourcePage) Generate(w io.Writer) error {
	var tmpl = template.Must(template.ParseFiles(path.Join(rp.templateBasePath, "templates/resource.html")))

	err := tmpl.Execute(w, rp)
	if err != nil {
		return fmt.Errorf("cannot write resource %s: %s", rp.Repository, err)
	}
	return nil
}
