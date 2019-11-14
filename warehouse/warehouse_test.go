package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/dutyfree/infoData"
)

var _ = Describe("Warehouse main", func() {

	Context("Bootstrapping Server", func() {
		It("returns JSON on API call", func() {
			response, err := http.Get("http://localhost:9090/info")

			Expect(err).ToNot(HaveOccurred())

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			defer response.Body.Close()

			responseString, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())

			var responseData infoData.Info
			err = json.Unmarshal(responseString, &responseData)
			Expect(err).ToNot(HaveOccurred())

			Expect(string(responseData.Data)).To(Equal("this is dutyfree"))
		})
	})
})
