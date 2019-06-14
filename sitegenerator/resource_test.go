package sitegenerator_test

import (
	"fmt"

	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Resources", func() {
	It("extracts all the values", func() {
		var r sitegenerator.Resource
		err := yaml.Unmarshal([]byte(`---
name: time resource
repository: https://github.com/concourse/time-resource
`), &r)

		Expect(err).ToNot(HaveOccurred())
		Expect(r).To(MatchAllFields(Fields{
			"Name":       Equal("time resource"),
			"Repository": Equal("https://github.com/concourse/time-resource"),
		}))
	})

	It("fails if repository is not part of github and has less than 5 components", func() {
		resources := []sitegenerator.Resource{
			{Name: "time resource", Repository: "https://github.com/concourse"},
		}
		_, err := sitegenerator.Enrich(resources)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("invalid repository for the resource (%s)", "time resource")))

	})

	It("calculates the additional fields", func() {
		resources := []sitegenerator.Resource{
			{Name: "time resource", Repository: "https://github.com/concourse/time-resource"},
		}

		i, err := sitegenerator.Enrich(resources)

		Expect(err).ToNot(HaveOccurred())

		Expect(i[0]).To(Equal(sitegenerator.ResourceModel{
			Resource: sitegenerator.Resource{
				Name:       "time resource",
				Repository: "https://github.com/concourse/time-resource",
			},
			Identifier:        "concourse-time-resource",
			AuthorHandle:      "concourse",
			AuthorProfileLink: "https://github.com/concourse",
		}))
	})
})
