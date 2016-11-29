package cfmanifest_test

import (
	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/cloudfoundry-community/cf-ssh/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
	Describe("NewManifestFromPath", func() {
		testWithFixture := func(path string) {
			path, err := fixtures.FixturePath(path)
			Expect(err).NotTo(HaveOccurred())
			manifest, err := cfmanifest.NewManifestFromPath(path)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest.Applications())).To(Equal(1))
			app := manifest.FirstApplication()
			Expect(app["name"]).To(Equal("oneapp"))
		}

		Context("with 'applications' root key", func() {
			It("loads manifest", func() {
				testWithFixture("manifest-oneapp.yml")
			})
		})

		Context("without 'applications' root key", func() {
			It("loads manifest", func() {
				testWithFixture("manifest-oneapp-without-root-node.yml")
			})
		})
	})

	Describe("AddApplication", func() {
		It("adds first app", func() {
			manifest := cfmanifest.NewManifest()
			Expect(len(manifest.Applications())).To(Equal(0))
			manifest.AddApplication("first")
			Expect(len(manifest.Applications())).To(Equal(1))
			first := manifest.FirstApplication()
			Expect(first["name"]).To(Equal("first"))
		})

	})
})
