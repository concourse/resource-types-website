package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	. "github.com/concourse/dutyfree/matchers"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Dutyfree", func() {
	var (
		server    *ghttp.Server
		outputDir string
		resources *os.File
	)

	BeforeEach(func() {
		var err error
		server = ghttp.NewServer()

		outputDir, err = ioutil.TempDir("", "dutyfree")
		Expect(err).ToNot(HaveOccurred())

		resources, err = ioutil.TempFile("", "resources.yml")
		Expect(err).ToNot(HaveOccurred())

		_, err = fmt.Fprint(resources, `---
- repository: https://github.com/concourse/git-resource
  name: git resource
  desc: git resource description
- repository: https://github.com/concourse/hg-resource
  name: hg resource
  desc: 
- repository: https://github.com/concourse/foo-resource
  name: foo resource
`)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("when a directory and a resources file are provided", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/repos/concourse/git-resource/readme"),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(http.StatusOK, `<div id="readme">git foo</div>`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/repos/concourse/hg-resource/readme"),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(http.StatusOK, `<div id="readme">hg foo</div>`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/repos/concourse/foo-resource/readme"),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(http.StatusOK, `<div id="readme">foo foo</div>`),
				))
		})

		It("generates all the website in the output folder", func() {
			cmd := exec.Command(pathToBin, outputDir, resources.Name())
			cmd.Env = append(cmd.Env, "GITHUB_API_ENDPOINT="+server.URL(), "GITHUB_TOKEN=SOMEGITHUBTOKEN")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			By("creating the index.html page")
			indexHTML, err := os.Open(filepath.Join(outputDir, "index.html"))

			Expect(err).ToNot(HaveOccurred())

			doc, err := goquery.NewDocumentFromReader(indexHTML)

			Expect(err).ToNot(HaveOccurred())

			Expect(doc).To(
				SatisfyAll(
					ContainSelectorWithText("title", Equal("Duty Free")),
					ContainSelector(`a[href="resources/concourse-git-resource.html"]`),
					ContainSelector(`a[href="resources/concourse-hg-resource.html"]`)))

			Expect(doc).To(SatisfyAll(
				ContainSelectorWithText("#concourse-git-resource .desc", Equal("git resource description")),
				ContainSelectorWithText("#concourse-hg-resource .desc", BeZero()),
				ContainSelectorWithText("#concourse-foo-resource .desc", BeZero()),
			))

			By("copying the static folder")
			staticSrcDir, err := ioutil.ReadDir(filepath.Join(outputDir, "static"))
			Expect(err).ToNot(HaveOccurred())

			staticDstDir, err := ioutil.ReadDir("static")
			Expect(err).ToNot(HaveOccurred())

			for i := range staticSrcDir {
				Expect(staticDstDir[i].Name()).To(Equal(staticSrcDir[i].Name()))
				Expect(staticDstDir[i].Size()).To(Equal(staticSrcDir[i].Size()))
			}

			By("creating a html page for each resource")
			Expect(server.ReceivedRequests()).Should(HaveLen(3))

			for _, resource := range []string{"git", "hg"} {
				resourceHTML, err := os.Open(filepath.Join(outputDir, fmt.Sprintf("resources/concourse-%s-resource.html", resource)))
				Expect(err).ToNot(HaveOccurred())

				doc, err = goquery.NewDocumentFromReader(resourceHTML)
				Expect(err).ToNot(HaveOccurred())

				Expect(doc).To(SatisfyAll(
					ContainSelectorWithText("title", ContainSubstring("Duty Free - %s resource", resource)),
					ContainSelectorWithText("#github-readme #readme", ContainSubstring("%s foo", resource))))
			}

		})

		AfterEach(func() {
			os.RemoveAll(outputDir)
			os.Remove(resources.Name())
		})

	})

	Describe("when github returns a non 200 status", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", MatchRegexp("/repos/concourse/git-resource/readme")),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(500, "error message"),
				))
		})

		It("exits 1 and print the error cause", func() {
			cmd := exec.Command(pathToBin, outputDir, resources.Name())
			cmd.Env = append(cmd.Env, "GITHUB_API_ENDPOINT="+server.URL(), "GITHUB_TOKEN=SOMEGITHUBTOKEN")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))

			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Eventually(session.Err).Should(gbytes.Say("Unable to access readme for concourse/git-resource due to error: 500, reason: error message"))
		})
	})

	Describe("when no parameter is passed", func() {
		It("exits 1 and prints usage message", func() {
			cmd := exec.Command(pathToBin)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Err).Should(gbytes.Say("undefined output directory"))
			Eventually(session.Err).Should(gbytes.Say("usage: %s <output-directory> <resource-file>", pathToBin))
		})
	})

	Describe("when the directory does not exits", func() {
		It("exits 1 and prints usage message", func() {
			cmd := exec.Command(pathToBin, "/a/folder/that/does/not/exists", "a-resources.yaml")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Err).Should(gbytes.Say("output directory cannot be found"))
			Eventually(session.Err).Should(gbytes.Say("usage: %s <output-directory> <resource-file>", pathToBin))
		})
	})

	Describe("when the resources file does not exists", func() {
		It("exits 1 and prints usage message", func() {
			outputDir, err := ioutil.TempDir("", "dutyfree")
			Expect(err).ToNot(HaveOccurred())
			cmd := exec.Command(pathToBin, outputDir, "a-resources-that-does-not-exists.yaml")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Err).Should(gbytes.Say("cannot read resources file"))
			Eventually(session.Err).Should(gbytes.Say("usage: %s <output-directory> <resource-file>", pathToBin))
		})
	})

	AfterEach(func() {
		server.Close()
	})
})
