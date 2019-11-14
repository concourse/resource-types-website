package publichandler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/concourse/dutyfree/server/publichandler"
)

var _ = Describe("public server", func() {
	Context("handler", func() {
		var server *httptest.Server
		BeforeEach(func() {
			path := "./test_public"
			sHandler := publichandler.Handler{Path: path}

			server = httptest.NewServer(sHandler)
		})
		It("returns files if they exist", func() {
			defer server.Close()

			res, err := http.Get(server.URL + "/public/index.js")
			Expect(err).NotTo(HaveOccurred())
			var body []byte
			body, err = ioutil.ReadAll(res.Body)
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			Expect(err).NotTo(HaveOccurred())

			Expect(string(body)).To(ContainSubstring("this is a test file"))
		})
		It("returns 404 for files that don't exist", func() {
			defer server.Close()

			res, err := http.Get(server.URL + "/public/index2.js")
			Expect(err).NotTo(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusNotFound))
		})
	})
})
