package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/codegangsta/cli"
)

func cmdSSH(c *cli.Context) {
	fmt.Println(c.String("manifest"))
	manifestPath, err := filepath.Abs(c.String("manifest"))
	if err != nil {
		log.Fatal(err)
	}
	var manifest *cfmanifest.Manifest
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		appName := c.Args().First()
		if appName == "" {
			fmt.Println("USAGE: cf-ssh [APPNAME] [-f manifest.yml]")
			return
		}

		manifest = cfmanifest.NewSSHManifest(appName)
	} else {
		manifest, err = cfmanifest.NewSSHManifestFromManifestPath(manifestPath)
		if err != nil {
			log.Fatalf("Manifest %s exists but failed to load: %s", manifestPath, err)
		}
	}

	cfSSHYAML, err := filepath.Abs("cf-ssh.yml")
	if err != nil {
		log.Fatalf("Could not create absolute file path: %s", err)
	}

	manifest.Save(cfSSHYAML)
}

func main() {
	app := cli.NewApp()
	app.Name = "cf-ssh"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "manifest, f",
			Value: "manifest.yml",
			Usage: "Path to application manifest",
		},
	}

	app.Usage = "SSH into a Cloud Foundry app container"
	app.Action = cmdSSH

	app.Run(os.Args)
}
