package main

import (
	nhs "nhsvalidator/nhs"
	"testing"
)

func TestGenerate(t *testing.T) {
	NHS, err := nhs.NewNHSManager()
	if err != nil {
		t.FailNow()
	}

	for i := 0; i < 5; i++ { // Test multiple times to ensure reliability
		generatedNumber, _ := NHS.Generate()
		if err != nil {
			t.Fatalf("Generation of NHS number failed")
		}

		// Validate the generated number
		isValid, _ := NHS.Validate(generatedNumber)
		// if err != nil {
		// 	t.FailNow()
		// }
		if !isValid{
			t.Errorf("Generated number %s is invalid", generatedNumber)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		index    string
		number   string
		target	 bool
	}{
		{"valid 1", "5990128088", true},
		{"valid 2", "1275988113", true},
		{"valid 3", "4536026665", true},
		{"INVALID 1", "5990128087", false},
		{"INVALID 2", "4536016660", false},
	}

	NHS, err := nhs.NewNHSManager()
	if err != nil {
		t.FailNow()
	}

	for _, tt := range tests {
		t.Run(tt.index, func(t *testing.T) {
			isValid, err := NHS.Validate(tt.number)
			if isValid != tt.target {
				t.Errorf("NHS Number(%s) = %v; want %v; error - %v", tt.number, isValid, tt.target, err)
			}
		})
	}
}