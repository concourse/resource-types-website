package sitegenerator

import (
	"fmt"
	"io"
)

type ResourcePage struct {
	PageName string
	ResourceModel
	Path []string
}

func NewResourcePage(resourceModel ResourceModel) ResourcePage {
	return ResourcePage{"detail-page",resourceModel, append(IndexPagePath, resourceModel.Name)}
}

func (rp *ResourcePage) Generate(w io.Writer) error {
	err := load("resource.html").ExecuteTemplate(w, "resource.html", rp)

	if err != nil {
		return fmt.Errorf("cannot write resource %s: %s", rp.Repository, err)
	}
	return nil
}
