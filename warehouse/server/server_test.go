package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/githubwrapper/githubwrapperfakes"
	"github.com/concourse/dutyfree/resource"
	"github.com/gobuffalo/packr/v2"

	"github.com/concourse/dutyfree/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server Test", func() {
	var (
		srv        server.Server
		port       int
		serverAddr string
	)

	BeforeEach(func() {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		Expect(err).NotTo(HaveOccurred())

		l, err := net.ListenTCP("tcp", addr)
		Expect(err).NotTo(HaveOccurred())

		port = l.Addr().(*net.TCPAddr).Port
		l.Close()
		time.Sleep(1 * time.Second)

		fakeWrapper := &githubwrapperfakes.FakeWrapper{}
		retMap := make(map[string]int)
		retMap["concourse/git-resource"] = 10
		fakeWrapper.GetStarsReturns(retMap, nil)

		//TODO: counterfeiter
		srv = server.Server{
			Port:                     port,
			PublicFilesFetcher:       fetcher.Fetcher{Box: packr.New("publicTestBox", "./testdata/public")},
			ResourceTypesFileFetcher: fetcher.Fetcher{Box: packr.New("resourcesTestBox", "./testdata/resource-types")},
			GithubGraphqlWrapper:     fakeWrapper,
		}
		srv.Start()
		time.Sleep(1 * time.Second)

		serverAddr = net.JoinHostPort("localhost", strconv.Itoa(port))
	})

	AfterEach(func() {
		err := srv.Close()
		Expect(err).NotTo(HaveOccurred())
	})

	Context("server initialization", func() {
		It("server runs and listens on port", func() {
			conn, err := net.DialTimeout("tcp", serverAddr, time.Second)
			Expect(err).NotTo(HaveOccurred())
			Expect(conn).NotTo(Equal(nil))
			err = conn.Close()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("API", func() {
		It("returns the resources on calls to /api/v1/resources", func() {
			resp, err := http.Get("http://" + serverAddr + "/api/v1/resources")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			var reses []resource.Resource
			err = json.Unmarshal(body, &reses)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(reses)).To(Equal(1))
			Expect(reses[0].Name).To(Equal("git"))
			Expect(reses[0].Stars).To(Equal("10"))
		})
		It("returns 404 on calls to unknown api /api/v1/res", func() {
			resp, err := http.Get("http://" + serverAddr + "/api/v1/res")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("serving public files", func() {
		It("returns index file on calls to /", func() {
			resp, err := http.Get("http://" + serverAddr + "/")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(body)).To(ContainSubstring("<html>"))
		})

		It("returns index file on calls to /public", func() {
			resp, err := http.Get("http://" + serverAddr + "/public/elm.js")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
