package persistence_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/persistence"
	"github.com/gobuffalo/packr/v2"
)

var _ = Describe("filesystem persistence", func() {
	Context("fetching", func() {
		It("is able to load resources from filesystem", func() {
			//TODO: counterfeiter
			box := packr.New("test", "./sample_resource_types")
			fs := persistence.Filesystem{
				Fetcher: fetcher.Fetcher{*box},
			}
			err := fs.LoadResources()
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns all the resources", func() {
			//TODO: counterfeiter
			box := packr.New("test", "./sample_resource_types")
			fs := persistence.Filesystem{
				Fetcher: fetcher.Fetcher{*box},
			}

			err := fs.LoadResources()
			Expect(err).NotTo(HaveOccurred())

			resources := fs.GetAllResources()
			Expect(len(resources)).To(Equal(2))

			Expect(resources[0].Name).To(Equal("test"))
			Expect(resources[0].URL).To(Equal("https://github.com/concourse/test"))
			Expect(resources[0].Description).To(ContainSubstring("line1"))
			Expect(resources[0].Description).To(ContainSubstring("line2"))
			Expect(resources[0].Owner).To(Equal("concourse"))
		})
	})

})
