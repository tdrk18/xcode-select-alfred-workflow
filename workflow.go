package main

func addItem(path string) {
	wf.NewItem(getAppName(path)).
		Valid(true)
}

func sendFeedback() {
	wf.SendFeedback()
}