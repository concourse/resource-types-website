package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/concourse/dutyfree/resource"
)

var _ = Describe("Warehouse main", func() {
	Context("Resources Endpoint", func() {
		It("returns JSON on API call", func() {
			token := os.Getenv("TOKEN")
			DutyfreeServerCommand.Env = []string{"GH_TOKEN=" + token}

			var err error
			session, err = gexec.Start(DutyfreeServerCommand, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())

			Eventually(func() *gbytes.Buffer {
				return session.Out
			}, 10*time.Second).Should(gbytes.Say("Dutyfree server started"))

			response, err := http.Get("http://localhost:9090/api/v1/resources")

			Expect(err).ToNot(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			defer response.Body.Close()

			responseString, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())

			var resources []resource.Resource
			err = json.Unmarshal(responseString, &resources)
			Expect(err).ToNot(HaveOccurred())

			dir, err := ioutil.ReadDir("resource-types")
			Expect(err).NotTo(HaveOccurred())

			// length  - 3 to remove the licence, readme and .git entries.
			Expect(len(resources)).To(Equal(len(dir) - 3))
		})

		AfterEach(func() {
			session.Kill()
			<-session.Exited
		})
	})
})
