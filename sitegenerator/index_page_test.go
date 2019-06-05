package sitegenerator_test

import (
	"bytes"

	. "github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var resources = []Resource{
	{"git", "https://github.com/concourse/git-resource"},
	{"hg", "https://github.com/concourse/hg-resource"},
}

var _ = Describe("IndexPage", func() {
	It("renders the template", func() {
		b := bytes.Buffer{}

		ip := NewIndexPage("", resources)
		ip.Generate(&b)

		Expect(b.String()).To(ContainSubstring("Duty Free"))
		Expect(b.String()).To(ContainSubstring("https://github.com/concourse/git-resource"))
		Expect(b.String()).To(ContainSubstring("https://github.com/concourse/hg-resource"))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := NewIndexPage("", []Resource{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
