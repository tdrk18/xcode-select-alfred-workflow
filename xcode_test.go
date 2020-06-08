package main

import (
	"fmt"
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

func TestGetAppName(t *testing.T) {
	result1 := getAppName("/Applications/Xcode11.5.app")
	if result1 != "Xcode11.5.app" {
		t.Fatal(fmt.Sprintf("failed: getAppName() returns %s", result1))
	}
	result2 := getAppName("/Applications/xcode11.5.app")
	if result2 != "xcode11.5.app" {
		t.Fatal(fmt.Sprintf("failed: getAppName() returns %s", result2))
	}
	result3 := getAppName("/Applications/Xcode_11_5.app")
	if result3 != "Xcode_11_5.app" {
		t.Fatal(fmt.Sprintf("failed: getAppName() returns %s", result3))
	}
	result4 := getAppName("/Applications/hoge.app")
	if result4 != "" {
		t.Fatal(fmt.Sprintf("failed: getAppName() returns %s", result4))
	}
}

func TestGetAppContentPath(t *testing.T) {
	result := getAppContentPath("/Applications/Xcode11.5.app")
	if result != "/Applications/Xcode11.5.app/Contents/Developer" {
		t.Fatal(fmt.Sprintf("failed: getAppContentPath() returns %s", result))
	}
}

func TestGetAppIconPath(t *testing.T) {
	result := getAppIconPath("/Applications/Xcode11.5.app")
	if result != "/Applications/Xcode11.5.app/Contents/Resources/Xcode.icns" {
		t.Fatal(fmt.Sprintf("failed: getAppIconPath() returns %s", result))
	}
}
