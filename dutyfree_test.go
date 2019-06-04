package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Dutyfree", func() {
	Describe("when a directory is given as parameter", func() {
		var (
			outputDir string
			err       error
		)

		BeforeEach(func() {
			outputDir, err = ioutil.TempDir("", "dutyfree")
			Expect(err).ToNot(HaveOccurred())
		})

		It("generates an index.html in the directory", func() {
			cmd := exec.Command(pathToBin, outputDir)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			content, err := ioutil.ReadFile(filepath.Join(outputDir, "index.html"))
			Expect(err).ToNot(HaveOccurred())
			Expect(content).To(ContainSubstring("Duty Free"))
		})

		AfterEach(func() {
			os.RemoveAll(outputDir)
		})
	})

	Describe("when no parameter is passed", func() {
		It("exits 1 and prints usage message", func() {
			cmd := exec.Command(pathToBin)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Err).Should(gbytes.Say("undefined output directory"))
			Eventually(session.Err).Should(gbytes.Say("usage: %s <output-directory>", pathToBin))
		})
	})

	Describe("when the directory does not exits", func() {
		It("exits 1 and prints usage message", func() {
			cmd := exec.Command(pathToBin, "/a/folder/that/does/not/exists")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Err).Should(gbytes.Say("output directory cannot be found"))
			Eventually(session.Err).Should(gbytes.Say("usage: %s <output-directory>", pathToBin))
		})
	})
})
