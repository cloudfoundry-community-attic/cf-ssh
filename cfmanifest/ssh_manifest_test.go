package cfmanifest_test

import (
	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/cloudfoundry-community/cf-ssh/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
	Describe("NewSSHManifestFromManifestPath", func() {
		It("keeps the one app in manifest", func() {
			path, err := fixtures.FixturePath("manifest-oneapp.yml")
			Expect(err).NotTo(HaveOccurred())

			manifest, err := cfmanifest.NewSSHManifestFromManifestPath(path)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest.Applications())).To(Equal(1))
			app := manifest.FirstApplication()
			Expect(app["name"]).To(Equal("oneapp-ssh"))
			Expect(app["command"]).To(Equal("curl http://tmate-bootstrap.cfapps.io | sh"))
			Expect(app["no-route"]).To(Equal(true))
			Expect(app["instances"]).To(Equal(1))
		})

		It("keeps the first app in manifest", func() {
			path, err := fixtures.FixturePath("manifest-twoapps.yml")
			Expect(err).NotTo(HaveOccurred())

			manifest, err := cfmanifest.NewSSHManifestFromManifestPath(path)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest.Applications())).To(Equal(1))
			app := manifest.FirstApplication()
			Expect(app["name"]).To(Equal("first-ssh"))
			Expect(app["command"]).To(Equal("curl http://tmate-bootstrap.cfapps.io | sh"))
			Expect(app["no-route"]).To(Equal(true))
			Expect(app["instances"]).To(Equal(1))
		})
	})
})
