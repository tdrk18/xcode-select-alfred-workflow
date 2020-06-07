package main

import (
	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			for _, repo := range filterXcodeApp(execMDFind()) {
				addItem(repo)
			}
			sendFeedback()
			return nil
		},
	}
	return app
}
