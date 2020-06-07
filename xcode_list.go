package main

import (
	"strings"
)

func filterXcodeApp(list []string) []string {
	var values []string
	for _, path := range list {
		if !strings.Contains(path, "/Contents/Developer") {
			values = append(values, path)
		}
	}
	return values
}