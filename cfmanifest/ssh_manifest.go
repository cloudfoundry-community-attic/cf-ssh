package cfmanifest

// NewSSHManifest prepares for a new cf-ssh.yml
func NewSSHManifest(appName string) (manifest *Manifest) {
	manifest = NewManifest()
	cfssh := manifest.AddApplication(appName)
	cfssh["command"] = "curl http://tmate-bootstrap.cfapps.io | sh"
	cfssh["no-route"] = true
	cfssh["instances"] = 1
	return
}

// NewSSHManifestFromManifestPath prepares for a new cf-ssh.yml based on existing manifest.yml
func NewSSHManifestFromManifestPath(manifestPath string) (manifest *Manifest, err error) {
	manifest, err = NewManifestFromPath(manifestPath)
	if err != nil {
		return
	}
	cfssh := manifest.FirstApplication()
	name := cfssh["name"].(string)
	cfssh["name"] = name + "-ssh"
	cfssh["command"] = "curl http://tmate-bootstrap.cfapps.io | sh"
	cfssh["no-route"] = true
	cfssh["instances"] = 1

	manifest.RemoveAllButFirstApplication()
	return
}
