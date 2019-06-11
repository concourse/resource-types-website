package sitegenerator

import (
	"fmt"
	"strings"
)

type Resource struct {
	Name              string `yaml:"name"`
	Repository        string `yaml:"repository"`
	Identifier        string
	AuthorHandle      string
	AuthorProfileLink string
}

func (r *Resource) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]string{}
	unmarshal(&m)

	r.Name = m["name"]
	r.Repository = m["repository"]

	segmentsAll := strings.Split(r.Repository, "/")

	if len(segmentsAll) < 5 || segmentsAll[0] != "https:" || segmentsAll[2] != "github.com" {
		return fmt.Errorf("invalid repository for the resource (%s)", r.Repository)
	}

	segments := segmentsAll[3:]

	r.Identifier = strings.Join(segments, "-")
	r.AuthorHandle = segmentsAll[3]
	r.AuthorProfileLink = strings.Join(segmentsAll[:4], "/")

	return nil
}
