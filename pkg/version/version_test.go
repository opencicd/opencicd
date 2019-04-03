package version_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/opencicd/opencicd/pkg/version"
)

var _ = Describe("Version", func() {

	Describe("Version", func() {

		It("Should not be nil", func() {
			Expect(&Version{}).NotTo(BeNil())
		})
	})
})
