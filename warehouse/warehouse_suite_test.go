package main_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var dutyfreePath string
var session *gexec.Session

var _ = SynchronizedBeforeSuite(func() []byte {
	binPath, err := gexec.Build("github.com/concourse/dutyfree")
	Expect(err).NotTo(HaveOccurred())

	return []byte(binPath)
}, func(data []byte) {
	dutyfreePath = string(data)

	dutyfreeServerCommand := exec.Command(dutyfreePath)

	var err error
	session, err = gexec.Start(dutyfreeServerCommand, GinkgoWriter, GinkgoWriter)

	Expect(err).NotTo(HaveOccurred())

	Eventually(func() *gbytes.Buffer {
		return session.Out
	}).Should(gbytes.Say("Dutyfree server started"))

})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
	session.Kill()
})

func TestWarehouse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Warehouse Suite")
}
