package sitegenerator_test

import (
	"bytes"

	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourcePage", func() {
	It("renders the template", func() {
		resource := sitegenerator.Resource{Name: "foobar resource", Repository: "https://github.com/foo/foobar-resource"}

		b := bytes.Buffer{}

		ip := sitegenerator.NewResourcePage("", resource)
		ip.Generate(&b)

		Expect(b.String()).To(ContainSubstring("https://github.com/foo/foobar-resource"))
		Expect(b.String()).To(ContainSubstring("foobar resource"))
	})
})
