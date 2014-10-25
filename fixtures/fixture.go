package fixtures

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadFixture loads one of the fixtures in this folder
func LoadFixture(name string) (contents []byte, err error) {
	path, err := FixturePath(name)
	if err != nil {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		return
	}
	return ioutil.ReadAll(file)
}

// FixturePath returns the absolute path to a fixture file
func FixturePath(name string) (string, error) {
	return filepath.Abs(filepath.Join("../fixtures", name))
}
