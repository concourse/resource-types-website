package sitegenerator_test

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	. "github.com/concourse/dutyfree/matchers"
	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourcePage", func() {
	It("renders the template", func() {
		resourceModel := sitegenerator.ResourceModel{
			Resource: sitegenerator.Resource{
				Name:        "git resource",
				Repository:  "https://github.com/concourse/git-resource",
				Description: "git resource description",
				Get:         true,
				Put:         true,
				Example:     "This is an\n example",
			},
			Identifier:        "concourse-git-resource",
			AuthorHandle:      "concourse",
			AuthorProfileLink: "https://github.com/concourse",
			Readme:            "<div>foobar readme</div>",
		}

		b := bytes.Buffer{}

		ip := sitegenerator.NewResourcePage(resourceModel)
		err := ip.Generate(&b)

		Expect(err).ToNot(HaveOccurred())

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b.Bytes()))

		Expect(err).ToNot(HaveOccurred())

		Expect(doc).To(ContainSelector(".container.detail-page"))

		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText("h2", Equal("git resource")),
				ContainSelectorWithText("#github-readme > div", Equal("foobar readme")),
				ContainSelectorWithText(".desc", Equal("git resource description")),
				ContainSelectorWithText(`a[href="https://github.com/concourse"]`, Equal("concourse")),
				ContainSelectorWithText(`.breadcrumb a[href="/dutyfree"]`, Equal("All Resources")),
				ContainSelectorWithText(".breadcrumb span:last-child", Equal("git resource")),
				ContainSelector(`a[href="https://github.com/concourse/git-resource"] img[title="Resource Source on Github"]`),
				ContainSelectorWithText(".example code", Equal("This is an\n example")),
				ContainSelectorWithText(".get", Equal("GET")),
				ContainSelectorWithText(".put", Equal("PUT")),
			))
	})
})
