package payroll

import (
	"errors"
	"os"
	"testing"
)

func TestPay(t *testing.T) {
	testCases := []struct {
		name                   string
		inputFirst             string
		inputLast              string
		inputHourlyRate        float64
		inputHoursWorkedInYear float64
		wantedFullName         string
		wantedPay              float64
	}{
		{
			name:                   "Test Basic Pay",
			inputFirst:             "Celina",
			inputLast:              "Smith",
			inputHourlyRate:        40.00,
			inputHoursWorkedInYear: 2000,
			wantedFullName:         "Celina Smith",
			wantedPay:              80_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{Individual: Employee{Id: 1, FirstName: tc.inputFirst, LastName: tc.inputLast}, HourlyRate: tc.inputHourlyRate, HoursWorkedInYear: tc.inputHoursWorkedInYear, Review: nil}
			gotName, gotSalary := d.Pay()
			if gotName != tc.wantedFullName || gotSalary != tc.wantedPay {
				t.Errorf("Got name: %v wanted name: %v Got salary: %v wanted salary %v", gotName, tc.wantedFullName, gotSalary, tc.wantedPay)
			}
		})
	}
}

func TestFullName(t *testing.T) {
	testCases := []struct {
		name   string
		fname  string
		lname  string
		wanted string
	}{
		{
			name:   "First test",
			fname:  "Cailyn",
			lname:  "Jones",
			wanted: "Cailyn Jones",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{Individual: Employee{Id: 1, FirstName: tc.fname, LastName: tc.lname}}
			gotName := d.fullName()
			if gotName != tc.wanted {
				t.Errorf("Got name: %v wanted name: %v ", gotName, tc.wanted)
			}
		})
	}
}

func TestConvertReviewToInt(t *testing.T) {
	if os.Getenv("TEST_NO_CI") != "" {
		t.Skip("Skipping, not yet prepared for CI")
	}

	testCases := []struct {
		name      string
		review    string
		wantedInt int
		wantedErr error
	}{
		{
			name:      "Test-Excellent",
			review:    "Excellent",
			wantedInt: 5,
			wantedErr: nil,
		},
		{
			name:      "Test-Good",
			review:    "Good",
			wantedInt: 4,
			wantedErr: nil,
		},
		{
			name:      "Test-Fair",
			review:    "Fair",
			wantedInt: 3,
			wantedErr: nil,
		},
		{
			name:      "Test-Poor",
			review:    "Poor",
			wantedInt: 2,
			wantedErr: nil,
		},
		{
			name:      "Test-Unsatisfactory",
			review:    "Unsatisfactory",
			wantedInt: 1,
			wantedErr: nil,
		},
		{
			name:      "Test-Error",
			review:    "great",
			wantedInt: 0,
			wantedErr: errors.New("invalid rating: great"),
		},
	}

	for _, tc := range testCases {
		gotrating, goterr := convertReviewToInt(tc.review)

		if gotrating != tc.wantedInt || goterr != tc.wantedErr {
			t.Errorf("Got rating: %v wanted: %v Got error: %v wanted %v", gotrating, tc.wantedInt, goterr, tc.wantedErr)
		}
	}
}
