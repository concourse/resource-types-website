package sitegenerator_test

import (
	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resources", func() {
	Describe("ExtractIdentifier", func() {
		It("returns an identifier that can be used in URLs", func() {
			r := sitegenerator.Resource{Name: "time", Repository: "https://github.com/concourse/time-resource"}

			Expect(r.ExtractIdentifier()).To(Equal("concourse-time-resource"))
		})

		DescribeTable("valid repositories", func(repo, expected string) {
			r := sitegenerator.Resource{Repository: repo}

			Expect(r.ExtractIdentifier()).To(Equal(expected))
		},
			Entry("basic", "https://github.com/concourse/time-resource", "concourse-time-resource"),
			Entry("google style", "https://github.com/google/concourse-resources/tree/master/gerrit", "google-concourse-resources-tree-master-gerrit"))

		It("returns an error if the repository is not defined", func() {
			r := sitegenerator.Resource{Name: "time"}
			_, err := r.ExtractIdentifier()

			Expect(err).To(HaveOccurred())
		})

		It("returns an error if the repository URL does not start with https://github.com", func() {
			r := sitegenerator.Resource{Repository: "https://git.ie/concourse/time-resource"}
			_, err := r.ExtractIdentifier()

			Expect(err).To(MatchError(ContainSubstring("invalid repository for the resource")))
		})
	})
})
