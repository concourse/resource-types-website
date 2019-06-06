package sitegenerator

import (
	"fmt"
	"strings"
)

type Resource struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

func (r Resource) ExtractIdentifier() (string, error) {
	identifiers := strings.Split(r.Repository, "/")
	if len(identifiers) >= 5 && identifiers[0] == "https:" && identifiers[2] == "github.com" {
		identifiers = identifiers[3:]
		return strings.Join(identifiers, "-"), nil
	}
	return "", fmt.Errorf("invalid repository for the resource (%s)", r.Repository)
}
