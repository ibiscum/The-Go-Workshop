package main

import (
	"os"
	"testing"
)

// TestGetDataAndParseResponse requires the server to be running to succeed
func TestGetDataAndParseResponse(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}

	err := addNameAndParseResponse("Electric")
	if err != nil {
		t.Fatal(err)
	}
	err = addNameAndParseResponse("Boogaloo")
	if err != nil {
		t.Fatal(err)
	}
	names := getDataAndParseResponse()
	if len(names) != 2 {
		t.Error("values don't match")
	}
}
