package main

import (
	"github.com/urfave/cli"
)

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Supervisor"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		run(c.String("listen"))
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "Host & port to listen to",
			Value: ":4445",
		},
	}
	return
}
