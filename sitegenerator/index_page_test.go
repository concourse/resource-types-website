package sitegenerator_test

import (
	"bytes"

	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IndexPage", func() {
	It("renders the template", func() {
		resourceModels := []sitegenerator.ResourceModel{
			{
				Resource: sitegenerator.Resource{
					Name:       "git resource",
					Repository: "https://github.com/concourse/git-resource",
				},
				Identifier:        "concourse-git-resource",
				AuthorHandle:      "concourse",
				AuthorProfileLink: "https://github.com/concourse",
			},
			{
				Resource: sitegenerator.Resource{
					Name:       "hg resource",
					Repository: "https://github.com/concourse/hg-resource",
				},
				Identifier:        "concourse-hg-resource",
				AuthorHandle:      "concourse",
				AuthorProfileLink: "https://github.com/concourse",
			},
		}

		b := bytes.Buffer{}

		ip := sitegenerator.NewIndexPage("", resourceModels)
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())

		Expect(b.String()).To(ContainSubstring("Duty Free"))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-git-resource.html"`))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-hg-resource.html"`))
		Expect(b.String()).To(ContainSubstring(`href="https://github.com/concourse"`))
		Expect(b.String()).To(ContainSubstring(`href="https://github.com/concourse/git-resource"`))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := sitegenerator.NewIndexPage("", []sitegenerator.ResourceModel{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
