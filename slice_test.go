package main

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {
	result1 := isEqual([]string{}, []string{})
	if !result1 {
		t.Fatal(fmt.Sprintf("failed: isEqual() returns %s", toString(result1)))
	}
	result2 := isEqual([]string{"hoge"}, []string{"hoge"})
	if !result2 {
		t.Fatal(fmt.Sprintf("failed: isEqual() returns %s", toString(result2)))
	}
	result3 := isEqual([]string{"hoge"}, []string{"fuga"})
	if result3 {
		t.Fatal(fmt.Sprintf("failed: isEqual() returns %s", toString(result3)))
	}
	result4 := isEqual([]string{"hoge"}, []string{"hoge", "fuga"})
	if result4 {
		t.Fatal(fmt.Sprintf("failed: isEqual() returns %s", toString(result4)))
	}
}

func toString(value bool) string {
	if value {
		return "true"
	} else {
		return "false"
	}
}
