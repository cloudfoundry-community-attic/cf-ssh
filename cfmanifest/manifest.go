package cfmanifest

import (
	"io/ioutil"
	"os"

	"launchpad.net/goyaml"
)

// Manifest models a manifest.yml
// See http://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html
type Manifest map[string]interface{}

// NewManifest creates a Manifest
func NewManifest() (manifest *Manifest) {
	return &Manifest{}
}

// NewManifestFromPath creates a Manifest from a manifest.yml file
func NewManifestFromPath(manifestPath string) (manifest *Manifest, err error) {
	manifest = &Manifest{}
	file, err := os.Open(manifestPath)
	if err != nil {
		return
	}
	yml, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = goyaml.Unmarshal(yml, manifest)
	return
}

// Applications returns the full list of applications
func (manifest Manifest) Applications() (apps []interface{}) {
	if manifest["applications"] == nil {
		return []interface{}{}
	}
	return manifest["applications"].([]interface{})
}

// FirstApplication returns the first application in the manifest
func (manifest Manifest) FirstApplication() map[interface{}]interface{} {
	app := manifest.Applications()[0]
	return app.(map[interface{}]interface{})
}

// ApplicationName returns the "name" of the first application in the manifest
func (manifest Manifest) ApplicationName() string {
	app := manifest.FirstApplication()
	return app["name"].(string)
}

// AddApplication adds a default manifestApp
func (manifest Manifest) AddApplication(appName string) (app map[interface{}]interface{}) {
	app = map[interface{}]interface{}{"name": appName}
	apps := manifest.Applications()
	apps = append(apps, app)
	manifest["applications"] = apps
	return
}

// RemoveAllButFirstApplication removes all applications but the first
func (manifest Manifest) RemoveAllButFirstApplication() {
	firstApp := manifest.Applications()[0]
	apps := []interface{}{firstApp}
	manifest["applications"] = apps
	return
}

// Save the Manifest to a file in YAML format
func (manifest Manifest) Save(path string) (err error) {
	data, err := goyaml.Marshal(manifest)
	if err != nil {
		return
	}
	ioutil.WriteFile(path, data, 0644)
	return
}
