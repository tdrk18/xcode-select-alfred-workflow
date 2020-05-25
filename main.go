package main

import (
	aw "github.com/deanishe/awgo"
	"log"
	"os"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	wf.Run(run)
}
