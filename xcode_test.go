package main

import (
	"testing"
)

func TestFilterXcodeApp(t *testing.T) {
	list := []string{
		"/Applications/Xcode11.5.app",
		"/Applications/Xcode11.4.1.app",
		"/Applications/Xcode11.5.app/Contents/Developer/Applications/Xcode Server Builder.app",
		"/Applications/Xcode11.4.1.app/Contents/Developer/Applications/Xcode Server Builder.app",
	}
	result := filterXcodeApp(list)
	if !isEqual(result, []string{"/Applications/Xcode11.5.app", "/Applications/Xcode11.4.1.app"}) {
		t.Fatal("failed: filterXcodeApp() returns wrong value")
	}
}