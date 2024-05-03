package main

import (
	"os"
	"testing"
)

// TestPostDataAndReturnResponse requires the server to be running to succeed
func TestPostFileAndReturnResponse(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}
	data := postFileAndReturnResponse("./test.txt")
	if data != "./test.txt Uploaded!" {
		t.Errorf("Expected string './test.txt Uploaded!' but received: '%s'", data)
	}
}
