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
					Get:        true,
					Put:        true,
					Verified:   true,
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
					Get:        false,
					Put:        false,
					Verified:   true,
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
					Get:        true,
					Verified:   false,
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

		Expect(doc).To(ContainSelector(".container.index-page"))

		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText(".breadcrumb span:last-child", Equal("All Resources")),
				ContainSelector(`a[href="resources/concourse-git-resource.html"]`),
				ContainSelector(`a[href="resources/concourse-hg-resource.html"]`),
				ContainSelector(`a[href="https://github.com/concourse/git-resource"]`),
				ContainSelector(`img[title="Resource Source on Github"]`),
				ContainSelectorWithText("title", Equal("Duty Free"))))

		By("creating a tile for each resource")
		Expect(doc).To(
			SatisfyAll(
				ContainSelector("#concourse-git-resource"),
				ContainSelector("#concourse-hg-resource"),
				ContainSelector("#pivotal-bosh-resource")))

		By("creating appropriate get and put tags to resources")
		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText("#concourse-git-resource .get", Equal("GET")),
				ContainSelectorWithText("#concourse-git-resource .put", Equal("PUT")),
				Not(ContainSelectorWithText("#concourse-hg-resource .get", Equal("GET"))),
				Not(ContainSelectorWithText("#concourse-hg-resource .put", Equal("PUT"))),
				ContainSelectorWithText("#pivotal-bosh-resource .get", Equal("GET")),
				Not(ContainSelectorWithText("#pivotal-bosh-resource .put", Equal("PUT")))))

		By("generating categories")
		Expect(doc).To(
			SatisfyAll(
				ContainSelector(".categories"),
				ContainSelectorWithText(`.categories li a[data-category="Category1"]`, Equal("Category1 (1)")),
				ContainSelectorWithText(`.categories li a[data-category="Category2"]`, Equal("Category2 (2)"))))

		Expect(doc.Find(".categories li a.category-filter").Length()).To(Equal(4))

		By("adding verified tag to resources verified by concourse")
		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText("#concourse-git-resource .verified", Equal("Verified")),
				Not(ContainSelector("#pivotal-bosh-resource .verified"))))

	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := sitegenerator.NewIndexPage([]sitegenerator.ResourceModel{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
