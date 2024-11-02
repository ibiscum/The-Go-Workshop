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

	electricCount, boogalooCount := getDataAndParseResponse()
	if electricCount < 1 {
		t.Errorf("expected more than one name 'Electric', recieved: %d", electricCount)
	}
	if boogalooCount < 1 {
		t.Errorf("expected more than one name 'Boogaloo', recieved: %d", boogalooCount)
	}
}
