package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}
	var s bytes.Buffer
	log.SetOutput(&s)
	log.SetFlags(0)
	main()

	if s.String() != "500000000" {
		t.Error(s.String())
	}
}
