package main

import (
	aw "github.com/deanishe/awgo"
)

func addItem(path string) {
	wf.NewItem(getAppName(path)).
		Arg(getAppContentPath(path)).
		Icon(getIcon(path)).
		Valid(true)
}

func getIcon(path string) *aw.Icon {
	return &aw.Icon{Value: getAppIconPath(path)}
}

func filter(query string) {
	wf.Filter(query)
}

func sendFeedback() {
	wf.SendFeedback()
}

func warnEmpty() {
	wf.WarnEmpty("No matched app", "Please try new query")
}
