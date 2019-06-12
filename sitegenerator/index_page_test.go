package sitegenerator_test

import (
	"bytes"

	. "github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("IndexPage", func() {
	It("renders the template", func() {
		var resources []Resource

		err := yaml.Unmarshal([]byte(`---
- repository: https://github.com/concourse/git-resource
- repository: https://github.com/concourse/hg-resource
`), &resources)

		Expect(err).ToNot(HaveOccurred())

		b := bytes.Buffer{}

		ip := NewIndexPage("", resources)
		ip.Generate(&b)

		Expect(b.String()).To(ContainSubstring("Duty Free"))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-git-resource.html"`))
		Expect(b.String()).To(ContainSubstring(`href="resources/concourse-hg-resource.html"`))
		Expect(b.String()).To(ContainSubstring(`href="https://github.com/concourse"`))
		Expect(b.String()).To(ContainSubstring(`href="https://github.com/concourse/git-resource"`))
	})

	It("handles no resources", func() {
		b := bytes.Buffer{}

		ip := NewIndexPage("", []Resource{})
		err := ip.Generate(&b)
		Expect(err).ToNot(HaveOccurred())
	})
})
