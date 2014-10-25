package cfmanifest

// Manifest models a manifest.yml
// See http://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html
type Manifest struct {
	Apps []*ManifestApp `yaml:"applications"`
}

// ManifestApp describes an individual app as part of the manifest
type ManifestApp struct {
	Name      string            `yaml:"name"`
	Buildpack string            `yaml:"buildpack"`
	Command   string            `yaml:"command"`
	Domain    string            `yaml:"domain"`
	Instances int               `yaml:"instances"`
	Memory    string            `yaml:"memory"`
	Host      string            `yaml:"host"`
	Path      string            `yaml:"path"`
	Timeout   int               `yaml:"timeout"`
	NoRoute   bool              `yaml:"no-route"`
	EnvVars   map[string]string `yaml:"env"`
	Services  []string          `yaml:"services"`
}

// NewManifest creates a Manifest
func NewManifest() (manifest *Manifest) {
	return &Manifest{}
}

// AddApplication adds a default manifestApp
func (manifest *Manifest) AddApplication(appName string) (app *ManifestApp) {
	app = &ManifestApp{
		Name:    appName,
		Memory:  "1024M",
		Host:    appName,
		Timeout: 60,
	}
	manifest.Apps = append(manifest.Apps, app)
	return
}
