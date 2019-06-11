package sitegenerator_test

import (
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
			"Name":              Equal("time resource"),
			"Repository":        Equal("https://github.com/concourse/time-resource"),
			"Identifier":        Equal("concourse-time-resource"),
			"AuthorHandle":      Equal("concourse"),
			"AuthorProfileLink": Equal("https://github.com/concourse"),
		}))
	})

	It("extracts all the values, even when multiple resources are deeper in the repository", func() {
		var r sitegenerator.Resource
		err := yaml.Unmarshal([]byte(`---
name: gerrit
repository: https://github.com/google/concourse-resources/tree/master/gerrit
`), &r)

		Expect(err).ToNot(HaveOccurred())
		Expect(r).To(MatchAllFields(Fields{
			"Name":              Equal("gerrit"),
			"Repository":        Equal("https://github.com/google/concourse-resources/tree/master/gerrit"),
			"Identifier":        Equal("google-concourse-resources-tree-master-gerrit"),
			"AuthorHandle":      Equal("google"),
			"AuthorProfileLink": Equal("https://github.com/google"),
		}))
	})

	It("fails if repository is not part of github and has less than 5 components", func() {
		var r sitegenerator.Resource
		err := yaml.Unmarshal([]byte(`---
name: time resource
repository: https://bitbucket.com/concourse/time-resource
`), &r)
		Expect(err).To(HaveOccurred())

		err = yaml.Unmarshal([]byte(`---
name: time resource
repository: https://github.com/concourse
`), &r)
		Expect(err).To(HaveOccurred())
	})
})
