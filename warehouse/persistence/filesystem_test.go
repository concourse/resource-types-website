package persistence_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/githubwrapper/githubwrapperfakes"
	"github.com/concourse/dutyfree/persistence"
	"github.com/gobuffalo/packr/v2"
)

var _ = Describe("filesystem persistence", func() {
	Context("fetching", func() {
		It("returns all the resources", func() {
			fakeWrapper := &githubwrapperfakes.FakeWrapper{}
			retMap := make(map[string]int)
			retMap["concourse/test"] = 0
			retMap["concourse/test1"] = 10
			retMap["concourse/test2"] = 100
			retMap["concourse/test3"] = 1051
			retMap["concourse/test4"] = 1040
			retMap["concourse/test5"] = 9999
			retMap["concourse/test6"] = 100000
			retMap["concourse/test7"] = 106999
			fakeWrapper.GetStarsReturns(retMap, nil)

			box := packr.New("test", "./sample_resource_types")
			fs := persistence.Filesystem{
				Fetcher:      fetcher.Fetcher{*box},
				GhGqlWrapper: fakeWrapper,
			}

			err := fs.LoadResources()
			Expect(err).NotTo(HaveOccurred())

			resources := fs.GetAllResources()
			Expect(len(resources)).To(Equal(8))

			Expect(resources[0].Name).To(Equal("test"))
			Expect(resources[0].URL).To(Equal("https://github.com/concourse/test"))
			Expect(resources[0].Description).To(ContainSubstring("line1"))
			Expect(resources[0].Description).To(ContainSubstring("line2"))
			Expect(resources[0].Owner).To(Equal("@concourse"))

			Expect(resources[0].StarsCount).To(Equal(0))

			Expect(resources[0].Stars).To(Equal("0"))
			Expect(resources[1].Stars).To(Equal("10"))
			Expect(resources[2].Stars).To(Equal("100"))
			Expect(resources[3].Stars).To(Equal("1.1k"))
			Expect(resources[4].Stars).To(Equal("1k"))
			Expect(resources[5].Stars).To(Equal("10k"))
			Expect(resources[6].Stars).To(Equal("100k"))
			Expect(resources[7].Stars).To(Equal("107k"))

			Expect(fakeWrapper.GetStarsCallCount()).To(Equal(1))
		})
	})

})
