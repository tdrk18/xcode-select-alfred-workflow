package main

func addItem(path string) {
	wf.NewItem(getAppName(path)).
		Arg(getAppContentPath(path)).
		Valid(true)
}

func sendFeedback() {
	wf.SendFeedback()
}
