package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func execMDFind() []string {
	out, err := exec.Command("mdfind", "-onlyin", "/Applications", "-name", "Xcode").Output()
	if err != nil {
		return []string{}
	}
	trim := strings.Trim(string(out), "\n")
	return strings.Split(trim, "\n")
}

func filterXcodeApp(list []string) []string {
	var values []string
	for _, path := range list {
		if !strings.Contains(path, "/Contents/Developer") {
			values = append(values, path)
		}
	}
	return values
}

func getAppName(path string) string {
	r := regexp.MustCompile(`[xX]code.*\.app`)
	for _, name := range strings.Split(path, "/") {
		if r.MatchString(name) {
			return name
		}
	}
	return ""
}

func getAppContentPath(path string) string {
	path = strings.TrimRight(path, "/")
	return fmt.Sprintf("%s/Contents/Developer", path)
}

func getAppIconPath(path string) string {
	path = strings.TrimRight(path, "/")
	return fmt.Sprintf("%s/Contents/Resources/Xcode.icns", path)
}
