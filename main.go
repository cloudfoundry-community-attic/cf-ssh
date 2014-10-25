package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func cmdSSH(c *cli.Context) {
	appName := c.Args().First()
	if appName == "" {
		fmt.Println("USAGE: cf-ssh APPNAME")
		return
	}

}

func main() {
	app := cli.NewApp()
	app.Name = "cf-ssh"
	app.Usage = "SSH into a Cloud Foundry app container"
	app.Action = cmdSSH

	app.Run(os.Args)
}
