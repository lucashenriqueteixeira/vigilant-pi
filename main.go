package main

import (
	"log"
	"os"

	"github.com/petersondmg/vigilant-pi/cmd/admin"
	"github.com/petersondmg/vigilant-pi/cmd/watchdog"
	"github.com/petersondmg/vigilant-pi/lib/config"
	"github.com/urfave/cli"
)

const (
	// Name ...
	Name = "Vigilant PI"
	// Usage ...
	Usage = "Record videos from survillence cameras and upload them to the cloud"
	// Version ...
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Usage
	app.Version = Version

	app.Commands = []cli.Command{
		admin.Command(),
		watchdog.Command(),
	}

	config.Update()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
