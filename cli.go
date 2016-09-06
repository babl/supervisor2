package main

import (
	"github.com/urfave/cli"
)

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Supervisor"
	app.Version = Version
	app.Action = func(c *cli.Context) {
		run(c.String("listen"), c.String("kafka-brokers"), c.GlobalBool("debug"))
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "Host & port to listen to",
			Value: ":4445",
		},
		cli.StringFlag{
			Name:  "kafka-brokers, kb",
			Usage: "Comma separated list of kafka brokers",
			Value: "127.0.0.1:9092",
		},
		cli.StringFlag{
			Name:  "storage",
			Usage: "Endpoint for Babl storage",
			Value: "babl.sh:4443",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Enable debug mode & verbose logging",
			EnvVar: "BABL_DEBUG",
		},
	}
	return
}
