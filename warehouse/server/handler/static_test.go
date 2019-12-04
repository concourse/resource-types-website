package handler_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/concourse/dutyfree/server/handler"
)

var _ = Describe("static server", func() {
	Context("initializing handler", func() {
		It("initializes", func() {
			path := "./static-test"
			handler.NewstaticHandler(path)
		})
	})
})
