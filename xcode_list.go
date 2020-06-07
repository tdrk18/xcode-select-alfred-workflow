package main

import (
	"os/exec"
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