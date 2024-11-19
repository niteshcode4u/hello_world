package main

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Hello, World!"
	result := GetMessage()

	if result != expected {
		t.Errorf("GetMessage() = %q; want %q", result, expected)
	}
}
