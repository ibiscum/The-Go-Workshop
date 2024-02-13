package main

import (
	"os"
	"testing"
)

// TestGetDataWithCustomOptionsAndReturnResponse requires the server to be running to succeed
func TestGetDataWithCustomOptionsAndReturnResponse(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}

	data := getDataWithCustomOptionsAndReturnResponse()
	if data != "hello client!" {
		t.Errorf("Expected string 'hello client!' but received: '%s'", data)
	}
}
