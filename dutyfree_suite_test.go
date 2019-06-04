package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	pathToBin string
)

func TestDutyfree(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		var (
			err error
		)

		pathToBin, err = gexec.Build("github.com/concourse/dutyfree")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Dutyfree Suite")
}
