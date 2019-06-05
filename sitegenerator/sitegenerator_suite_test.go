package sitegenerator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSitegenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sitegenerator Suite")
}
