package cfmanifest

// NewSSHManifest prepares for a new cf-ssh.yml
func NewSSHManifest(appName string) (manifest *Manifest) {
	manifest = NewManifest()
	cfssh := manifest.AddApplication(appName)
	cfssh.Command = "curl http://tmate-bootstrap.cfapps.io | sh"
	cfssh.NoRoute = true
	return
}

// NewSSHManifestFromManifestPath prepares for a new cf-ssh.yml based on existing manifest.yml
func NewSSHManifestFromManifestPath(appName string, manifestPath string) (manifest *Manifest, err error) {
	manifest, err = NewManifestFromPath(manifestPath)
	if err != nil {
		return
	}
	cfssh := manifest.Apps[0]
	cfssh.Command = "curl http://tmate-bootstrap.cfapps.io | sh"
	cfssh.NoRoute = true

	manifest.Apps = []*ManifestApp{cfssh}
	return
}
