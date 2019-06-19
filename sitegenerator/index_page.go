package sitegenerator

import (
	"fmt"
	"io"
)

type IndexPage struct {
	PageName       string
	ResourceModels []ResourceModel
}

func NewIndexPage(resourceModels []ResourceModel) IndexPage {
	return IndexPage{"index", resourceModels}
}

func (i *IndexPage) Generate(w io.Writer) error {
	err := load("index.html").ExecuteTemplate(w, "index.html", i.ResourceModels)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	return nil
}
