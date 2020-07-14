package githubwrapper_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/concourse/dutyfree/githubwrapper"
)

var _ = Describe("github wrapper", func() {
	Context("stars", func() {
		var (
			testServer *ghttp.Server
			wrapper    githubwrapper.Wrapper
			testStars  map[string]bool
		)
		BeforeEach(func() {
			testServer = ghttp.NewServer()

			wrapper = githubwrapper.NewWrapper(testServer.URL(), "token")
			testStars = make(map[string]bool)
			testStars["concourse/concourse"] = true
		})

		It("the underlying module does the appropriate call to the graphql api", func() {
			testServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/"),
					ghttp.VerifyHeader(http.Header(map[string][]string{"Authorization": {"Bearer token"}})),
					ghttp.RespondWith(http.StatusOK,
						`{"data":{"i0":{"nameWithOwner":"concourse/concourse","stargazers":{"totalCount":9}}}}`,
						http.Header(map[string][]string{"Content-Type": {"application/json"}}),
					),
				),
			)

			returnStars, err := wrapper.GetStars(testStars)

			Expect(err).ToNot(HaveOccurred())
			Expect(returnStars["concourse/concourse"]).To(Equal(9))
		})

		It("returns the appropriate error in case of server failure", func() {
			testServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/"),
					ghttp.VerifyHeader(http.Header(map[string][]string{"Authorization": {"Bearer token"}})),
					ghttp.RespondWith(http.StatusInternalServerError,
						``,
						http.Header(map[string][]string{"Content-Type": {"application/json"}}),
					),
				),
			)

			_, err := wrapper.GetStars(testStars)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("500"))
		})

		AfterEach(func() {
			testServer.Close()
		})
	})
})
