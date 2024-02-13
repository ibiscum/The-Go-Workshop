package main

import (
	"os"
	"testing"
)

// TestGetDataAndReturnResponse requires the server to be running to succeed
func TestGetDataAndReturnResponse(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}

	data := getDataAndReturnResponse()
	if data.Message != "hello world" {
		t.Errorf("Expected string 'hello world' but received: '%s'", data)
	}
}
