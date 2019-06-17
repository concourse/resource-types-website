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
	Describe("when a directory and a resources file are provided", func() {
		var (
			outputDir string
			resources *os.File
			err       error
			server    *ghttp.Server
		)

		BeforeEach(func() {
			outputDir, err = ioutil.TempDir("", "dutyfree")
			Expect(err).ToNot(HaveOccurred())

			resources, err = ioutil.TempFile("", "resources.yml")
			Expect(err).ToNot(HaveOccurred())

			_, err = fmt.Fprint(resources, `---
- repository: https://github.com/concourse/git-resource
  name: git resource
- repository: https://github.com/concourse/hg-resource
  name: hg resource
`)
			Expect(err).ToNot(HaveOccurred())

			server = ghttp.NewServer()

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/repos/concourse/git-resource/readme"),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(http.StatusOK, `<div id="readme">foo git</div>`),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/repos/concourse/hg-resource/readme"),
					ghttp.VerifyHeaderKV("Accept", "application/vnd.github.VERSION.html"),
					ghttp.VerifyHeaderKV("Authorization", "token SOMEGITHUBTOKEN"),
					ghttp.RespondWith(http.StatusOK, `<div id="readme">foo hg</div>`),
				))

			fmt.Println(server.URL())
		})

		It("generates an index.html in the directory", func() {
			cmd := exec.Command(pathToBin, outputDir, resources.Name())
			cmd.Env = append(cmd.Env, "GITHUB_API_ENDPOINT="+server.URL(), "GITHUB_TOKEN=SOMEGITHUBTOKEN")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			indexHTML, err := os.Open(filepath.Join(outputDir, "index.html"))

			Expect(err).ToNot(HaveOccurred())

			doc, err := goquery.NewDocumentFromReader(indexHTML)

			Expect(err).ToNot(HaveOccurred())

			Expect(doc).To(
				SatisfyAll(
					ContainSelectorWithText("title", Equal("Duty Free")),
					ContainSelector(`a[href="resources/concourse-git-resource.html"]`),
					ContainSelector(`a[href="resources/concourse-hg-resource.html"]`)))

			By("copying the static folder")
			staticSrcDir, err := ioutil.ReadDir(filepath.Join(outputDir, "static"))
			Expect(err).ToNot(HaveOccurred())

			staticDstDir, err := ioutil.ReadDir("static")
			Expect(err).ToNot(HaveOccurred())

			for i := range staticSrcDir {
				Expect(staticDstDir[i].Name()).To(Equal(staticSrcDir[i].Name()))
				Expect(staticDstDir[i].Size()).To(Equal(staticSrcDir[i].Size()))
			}

			Expect(server.ReceivedRequests()).Should(HaveLen(2))

			resourceHTML, err := os.Open(filepath.Join(outputDir, "resources/concourse-git-resource.html"))
			Expect(err).ToNot(HaveOccurred())

			doc, err = goquery.NewDocumentFromReader(resourceHTML)
			Expect(err).ToNot(HaveOccurred())

			Expect(doc).To(SatisfyAll(
				ContainSelectorWithText("title", Equal("Duty Free - git resource")),
				ContainSelectorWithText("body", ContainSubstring("https://github.com/concourse/git-resource")),
				ContainSelectorWithText("#github-readme #readme", Equal("foo git"))))
		})

		AfterEach(func() {
			os.RemoveAll(outputDir)
			os.Remove(resources.Name())
			server.Close()
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
})
