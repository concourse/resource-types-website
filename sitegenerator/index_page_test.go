package sitegenerator_test

import (
	"bytes"

	. "github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IndexPage", func() {
	It("renders the template", func() {
		var resources = []Resource{
			{"git", "https://github.com/concourse/git-resource"},
			{"hg", "https://github.com/concourse/hg-resource"},
		}

		b := bytes.Buffer{}

		ip := NewIndexPage("", resources)
		ip.Generate(&b)

		Expect(b.String()).To(ContainSubstring("Duty Free"))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-git-resource.html"`))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-hg-resource.html"`))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := NewIndexPage("", []Resource{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
