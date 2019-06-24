package sitegenerator

import (
	"fmt"
	"io"
)

type ResourcePage struct {
	ResourceModel
	Path []string
}

func NewResourcePage(resourceModel ResourceModel) ResourcePage {
	return ResourcePage{resourceModel, []string{"Home", resourceModel.Name}}
}

func (rp *ResourcePage) Generate(w io.Writer) error {
	err := load("resource.html").ExecuteTemplate(w, "resource.html", rp)

	if err != nil {
		return fmt.Errorf("cannot write resource %s: %s", rp.Repository, err)
	}
	return nil
}
