package server_test

import (
	"net"
	"net/http"
	"strconv"
	"time"

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
		port = 9000
		srv = server.Server{
			Port: port,
		}
		serverAddr = net.JoinHostPort("localhost", strconv.Itoa(port))
		srv.Start()
	})

	Context("server initialization", func() {
		It("server runs and listens on port", func() {
			conn, err := net.DialTimeout("tcp", serverAddr, time.Second)
			Expect(err).NotTo(HaveOccurred())
			Expect(conn).NotTo(Equal(nil))
			conn.Close()
		})
	})

	Context("serving public files", func() {
		//It("returns index file on calls to /", func() {
		//	resp, err := http.Get("http://" + serverAddr + "/")
		//	Expect(err).NotTo(HaveOccurred())
		//	Expect(resp.Status).To(Equal(http.StatusOK))
		//})

		It("returns index file on calls to /public", func() {
			resp, err := http.Get("http://" + serverAddr + "/public/elm.js")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
