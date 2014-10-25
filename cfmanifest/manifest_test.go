package cfmanifest_test

import (
	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/cloudfoundry-community/cf-ssh/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
	Describe("NewManifestFromPath", func() {
		It("loads manifest", func() {
			path, err := fixtures.FixturePath("manifest-oneapp.yml")
			Expect(err).NotTo(HaveOccurred())
			manifest, err := cfmanifest.NewManifestFromPath(path)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest.Apps)).To(Equal(1))
			app := manifest.Apps[0]
			Expect(app.Name).To(Equal("oneapp"))
		})
	})

	Describe("AddApplication", func() {
		It("adds first app", func() {
			manifest := cfmanifest.NewManifest()
			Expect(len(manifest.Apps)).To(Equal(0))
			app := manifest.AddApplication("first")
			Expect(len(manifest.Apps)).To(Equal(1))
			Expect(app.Name).To(Equal("first"))
		})

	})
})
