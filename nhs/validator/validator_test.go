package validator

import (
	"testing"
)


func TestInputValidation(t *testing.T) {
	inputs := []struct {
		name    string
		number   string
		target	 error
	}{
		{"10 digits", "5990128088", nil},
		{"9 digits", "127598811", ErrBadInput},
		{"mixed chars", "AC3602666F", ErrBadInput},
	}
	
	vC, err := NewValidationChain()
	if err != nil {
		t.Errorf("Could not init validation chain")
	}

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			err := vC.inputValidattion(tt.number)
			if err != tt.target {
				t.Errorf("NHS Number(%s); want %v; error - %v", tt.number, tt.target, err)
			}
		})
	}

}

func TestGetRemainder(t *testing.T) {
	inputs := []struct {
		name    string
		number   int
		target	 error
	}{
		{"0", 0, ErrDivisionNotAllowed},
		{"< 0", -293134, ErrDivisionNotAllowed},
		{"valid", 70273820, nil},
	}
	
	vC, err := NewValidationChain()
	if err != nil {
		t.Errorf("Could not init validation chain")
	}

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vC.GetRemainder(tt.number)
			if err != tt.target {
				t.Errorf("Number(%d); want %v; error - %v", tt.number, tt.target, err)
			}
		})
	}

}

func TestProcessMap(t *testing.T) {
	validInputs := []struct {
		name     string
		number   string
	}{
		{"Valid 1", "70273822"},
		{"Valid 2", "70273820"},
	}
	
	INValidInputs := []struct {
		name     string
		number   string
	}{
		{"INVALID 1", "ABcdEf"},
		{"INVALID 2", "&0273AR"},
	}
	
	vC, err := NewValidationChain()
	if err != nil {
		t.Errorf("Could not init validation chain")
	}

	for _, tt := range validInputs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vC.ProcessMap(tt.number)
			if err != nil {
				t.Errorf("Number(%s); want %v; error - %v", tt.number, nil, err)
			}
		})
	}

	for _, tt := range INValidInputs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vC.ProcessMap(tt.number)
			if err == nil {
				t.Errorf("Number(%s); want %v; error - %v", tt.number, err, nil)
			}
		})
	}

}
