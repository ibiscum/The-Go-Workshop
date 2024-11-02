package main

import (
	"fmt"
	"log"
	"testing"
)

func TestNewRecord(t *testing.T) {
	testCases := []struct {
		name           string
		inputKey       string
		inputInterface interface{}
		wanted         record
	}{

		{
			name:           "Int Scenario",
			inputKey:       "intValue",
			inputInterface: 100,
			wanted:         record{key: "intValue", valueType: "int", data: 100},
		},
		{
			name:           "Bool Scenario",
			inputKey:       "boolValue",
			inputInterface: true,
			wanted:         record{key: "boolValue", valueType: "bool", data: true},
		},
		{
			name:           "String Scenario",
			inputKey:       "stringValue",
			inputInterface: 100,
			wanted:         record{key: "stringValue", valueType: "string", data: 100},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := newRecord(tc.inputKey, tc.inputInterface)
			if got.key != tc.wanted.key {
				err := fmt.Errorf("Got: %v wanted %v", got, tc.wanted.key)
				if err != nil {
					log.Fatal(err)
				}

			}
		})
	}
}
