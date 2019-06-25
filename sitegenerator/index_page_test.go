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
					Categories: []string{"Category1", "Category2"},
				},
				Identifier:        "concourse-git-resource",
				AuthorHandle:      "concourse",
				AuthorProfileLink: "https://github.com/concourse",
			},
			{
				Resource: sitegenerator.Resource{
					Name:       "hg resource",
					Repository: "https://github.com/concourse/hg-resource",
					Categories: []string{},
				},
				Identifier:        "concourse-hg-resource",
				AuthorHandle:      "concourse",
				AuthorProfileLink: "https://github.com/concourse",
			},
			{
				Resource: sitegenerator.Resource{
					Name:       "bosh resource",
					Repository: "https://github.com/pivotal-cf/bosh-resource",
					Categories: []string{"Category2", "Category3", "Category4"},
				},
				Identifier:        "pivotal-bosh-resource",
				AuthorHandle:      "bosh",
				AuthorProfileLink: "https://github.com/bosh",
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
				ContainSelectorWithText(".breadcrumb span:last-child", Equal("All Resources")),
				ContainSelector(`a[href="resources/concourse-git-resource.html"]`),
				ContainSelector(`a[href="resources/concourse-hg-resource.html"]`),
				ContainSelector(`a[href="https://github.com/concourse"]`),
				ContainSelector(`a[href="https://github.com/concourse/git-resource"]`),
				ContainSelector(`img[title="Resource Source on Github"]`),
				ContainSelectorWithText("title", Equal("Duty Free"))))

		Expect(doc).To(
			SatisfyAll(
				ContainSelector("#concourse-git-resource .official"),
				ContainSelector("#concourse-hg-resource .official"),
				And(
					ContainSelector("#pivotal-bosh-resource"),
					Not(ContainSelector("#pivotal-bosh-resource .official")))))

		Expect(doc).To(
			SatisfyAll(
				ContainSelector(".categories"),
			))

		Expect(doc.Find(".categories li").Length()).To(Equal(4))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := sitegenerator.NewIndexPage([]sitegenerator.ResourceModel{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
