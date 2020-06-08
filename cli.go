package main

import (
	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			for _, path := range filterXcodeApp(execMDFind()) {
				addItem(path)
			}
			sendFeedback()
			return nil
		},
	}
	return app
}
