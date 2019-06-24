package sitegenerator

import (
	"fmt"
	"io"
)

type IndexPage struct {
	ResourceModels []ResourceModel
	Path           []string
}

func NewIndexPage(resourceModels []ResourceModel) IndexPage {
	return IndexPage{resourceModels, []string{"Home"}}
}

func (i *IndexPage) Generate(w io.Writer) error {
	err := load("index.html").ExecuteTemplate(w, "index.html", i)
	if err != nil {
		return fmt.Errorf("cannot write index.html: %s", err)
	}
	return nil
}
