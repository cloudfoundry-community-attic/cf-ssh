package cfmanifest_test

import (
	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
	Describe("AddApplication", func() {
		var manifest *cfmanifest.Manifest

		BeforeEach(func() {
			manifest = cfmanifest.NewManifest()
		})

		It("adds first app", func() {
			app := manifest.AddApplication("first")
			Expect(app.Name).To(Equal("first"))

		})

	})
})
