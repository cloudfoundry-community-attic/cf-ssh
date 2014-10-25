package cfmanifest_test

import (
	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/cloudfoundry-community/cf-ssh/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
	Describe("NewSSHManifestFromManifestPath", func() {
		It("loads manifest", func() {
			path, err := fixtures.FixturePath("manifest-oneapp.yml")
			Expect(err).NotTo(HaveOccurred())

			manifest, err := cfmanifest.NewSSHManifestFromManifestPath(path)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest.Apps)).To(Equal(1))
			app := manifest.Apps[0]
			Expect(app.Name).To(Equal("oneapp-ssh"))
			Expect(app.Command).To(Equal("curl http://tmate-bootstrap.cfapps.io | sh"))
			Expect(app.NoRoute).To(Equal(true))
		})
	})
})
