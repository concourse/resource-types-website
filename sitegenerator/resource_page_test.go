package sitegenerator_test

import (
	"bytes"

	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourcePage", func() {
	It("renders the template", func() {
		resourceModel := sitegenerator.ResourceModel{
			Resource: sitegenerator.Resource{
				Name:       "git resource",
				Repository: "https://github.com/concourse/git-resource",
			},
			Identifier:        "concourse-git-resource",
			AuthorHandle:      "concourse",
			AuthorProfileLink: "https://github.com/concourse",
			Readme:            "<div>foobar readme</div>",
		}

		b := bytes.Buffer{}

		ip := sitegenerator.NewResourcePage("", resourceModel)
		err := ip.Generate(&b)

		Expect(err).ToNot(HaveOccurred())
		Expect(b.String()).To(ContainSubstring("https://github.com/concourse/git-resource"))
		Expect(b.String()).To(ContainSubstring("git resource"))
		Expect(b.String()).To(ContainSubstring(`<div id="github-readme">`))
		Expect(b.String()).To(ContainSubstring("<div>foobar readme</div>"))
	})
})
