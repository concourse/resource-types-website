package githubwrapper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGithubwrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Githubwrapper Suite")
}
