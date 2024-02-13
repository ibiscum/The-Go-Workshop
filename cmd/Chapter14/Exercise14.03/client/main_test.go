package main

import (
	"os"
	"testing"
)

// TestPostDataAndReturnResponse requires the server to be running to succeed
func TestPostDataAndReturnResponse(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}
	msg := messageData{Message: "Testing 123"}
	data := postDataAndReturnResponse(msg)
	if data.Message != "Testing 123" {
		t.Errorf("Expected string 'Testing 123' but received: '%s'", data)
	}
}
