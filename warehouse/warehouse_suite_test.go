package main_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var (
	dutyfreePath          string
	session               *gexec.Session
	DutyfreeServerCommand *exec.Cmd
)

var _ = SynchronizedBeforeSuite(func() []byte {
	binPath, err := gexec.Build("github.com/concourse/dutyfree")
	Expect(err).NotTo(HaveOccurred())

	return []byte(binPath)
}, func(data []byte) {
	dutyfreePath = string(data)

	DutyfreeServerCommand = exec.Command(dutyfreePath)
})

func TestWarehouse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Warehouse Suite")
}
