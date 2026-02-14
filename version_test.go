package main

import "testing"

func TestGetVersion(t *testing.T) {
	expected := "0.1.0-alpha"
	actual := GetVersion()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
