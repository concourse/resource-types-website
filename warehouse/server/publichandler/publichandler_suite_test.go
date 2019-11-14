package publichandler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPublichandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Publichandler Suite")
}
