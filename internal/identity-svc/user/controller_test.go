package user_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/opencicd/opencicd/internal/identity-svc/user"
)

var _ = Describe("Controller", func() {

	Describe("NewController", func() {

		It("Should create a new instance", func() {
			c := NewController(nil)

			Expect(c).NotTo(BeNil())
		})
	})
})
