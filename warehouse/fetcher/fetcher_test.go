package fetcher_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/dutyfree/fetcher"
)

var _ = Describe("fetcher", func() {
	Context("Getting files", func() {
		It("returns the requested file", func() {
			fetchr := fetcher.Fetcher{Box: os.DirFS("testData")}
			fileByte, err := fetchr.GetFile("file1.yml")
			Expect(err).NotTo(HaveOccurred())
			Expect(string(fileByte)).To(ContainSubstring("name: name"))
		})

		It("returns all the files with their names as requested", func() {
			fetchr := fetcher.Fetcher{Box: os.DirFS("testData")}
			files, err := fetchr.GetAll()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(files)).To(Equal(2))
			Expect(files[0].Name).To(Equal("file1.yml"))
		})
	})
})
