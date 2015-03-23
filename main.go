package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/cloudfoundry-community/cf-ssh/cfmanifest"
	"github.com/codegangsta/cli"
)

func cmdSSH(c *cli.Context) {
	// TODO: confirm that `cf` and `ssh` are in path
	// TODO: Windows: cf.exe and ssh.exe?
	manifestPath, err := filepath.Abs(c.String("manifest"))
	if err != nil {
		log.Fatal(err)
	}
	var manifest *cfmanifest.Manifest
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		log.Fatal("USAGE: cf-ssh -f manifest.yml")

		// appName := c.Args().First()
		//
		// if appName == "" {
		// 	log.Fatal("USAGE: cf-ssh [APPNAME] [-f manifest.yml]")
		// }
		// manifest = cfmanifest.NewSSHManifest(appName)
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
	sshAppname := manifest.ApplicationName()
	fmt.Printf("Deploying SSH container '%s'...\n", sshAppname)

	// TODO: extract the `cf push` & log scraping
	cmd := exec.Command("cf", "push", "-f", cfSSHYAML)
	if c.Bool("verbose") {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	// TODO: defer cf delete
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run SSH container: %s", err)
	}

	var sshUser, sshHost string
	fmt.Print("Initiating tmate connection...")
	time.Sleep(1 * time.Second)
	for counter := 0; counter < 10; counter++ {
		time.Sleep(1 * time.Second)

		// repeat following until it succeeds or times out
		// ssh_host=$(cf logs $ssh_appname --recent | grep tmate.io | tail -n1 | awk '{print $NF }')
		cmd = exec.Command("cf", "logs", sshAppname, "--recent")
		var out bytes.Buffer
		cmd.Stdout = &out

		err = cmd.Run()
		if err != nil {
			log.Fatalf("Failed to get recent logs: %s", err)
		}
		logs := out.String()
		sshHostLine, err := regexp.CompilePOSIX("=====> (.*)@(.*)$")
		if err != nil {
			log.Fatalf("Invalid POSIX regular expression: %s", err)
		}
		sshHostMatches := sshHostLine.FindAllStringSubmatch(logs, -1)
		if sshHostMatches != nil {
			sshHostMatch := sshHostMatches[len(sshHostMatches)-1]
			sshUser = sshHostMatch[1]
			sshHost = sshHostMatch[2]
			break
		} else {
			fmt.Print(".")
		}

	}
	if sshUser == "" {
		fmt.Print("timed out\n")
	}

	fmt.Print("success\n")
	cmd = exec.Command("ssh", "-t", "-t", fmt.Sprintf("%s@%s", sshUser, sshHost))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// Either:
	// cf delete $ssh_appname -f
	// cf stop $ssh_appname

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
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose output",
		},
	}

	app.Usage = "SSH into a Cloud Foundry app container"
	app.Action = cmdSSH

	app.Run(os.Args)
}
