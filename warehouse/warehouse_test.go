package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/dutyfree/resource"
)

var _ = Describe("Warehouse main", func() {
	Context("Resources Endpoint", func() {
		It("returns JSON on API call", func() {
			response, err := http.Get("http://localhost:9090/api/v1/resources")

			Expect(err).ToNot(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			defer response.Body.Close()

			responseString, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())

			var resources []resource.Resource
			err = json.Unmarshal(responseString, &resources)
			Expect(err).ToNot(HaveOccurred())

			dir, err := ioutil.ReadDir("../resource-types")
			Expect(err).NotTo(HaveOccurred())

			// length  - 3 to remove the licence, readme and .git entries.
			Expect(len(resources)).To(Equal(len(dir) - 3))
		})
	})
})
