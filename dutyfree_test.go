package main_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Dutyfree", func() {
	It("runs successfully", func() {
		fmt.Println(pathToBin)
		cmd := exec.Command(pathToBin)

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		Eventually(session).Should(gexec.Exit())
		Eventually(session.Out).Should(gbytes.Say("dutyfree"))
	})
})
