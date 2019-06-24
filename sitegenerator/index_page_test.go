package sitegenerator_test

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	. "github.com/concourse/dutyfree/matchers"
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

		ip := sitegenerator.NewIndexPage(resourceModels)
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b.Bytes()))

		Expect(err).ToNot(HaveOccurred())

		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText(".breadcrumb span:last-child", Equal("Home")),
				ContainSelector(`a[href="resources/concourse-git-resource.html"]`),
				ContainSelector(`a[href="resources/concourse-hg-resource.html"]`),
				ContainSelector(`a[href="https://github.com/concourse"]`),
				ContainSelector(`a[href="https://github.com/concourse/git-resource"]`),
				ContainSelector(`img[title="Resource Source on Github"]`),
				ContainSelectorWithText("title", Equal("Duty Free"))))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := sitegenerator.NewIndexPage([]sitegenerator.ResourceModel{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
