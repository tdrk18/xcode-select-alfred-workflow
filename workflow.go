package main

func addItem(app string) {
	wf.NewItem(app).
		Valid(true)
}

func sendFeedback() {
	wf.SendFeedback()
}