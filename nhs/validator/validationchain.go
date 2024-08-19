package validator

import (
	"strconv"
	"unicode"
)

type ValidationChain struct{
	factorsPerDigits map[int]int
}

func NewValidationChain() (*ValidationChain, error){
	factorsPerDigits := map[int]int{
									1: 10,
									2: 9,
									3: 8,
									4: 7,
									5: 6,
									6: 5,
									7: 4,
									8: 3,
									9: 2,
								}

	return &ValidationChain{factorsPerDigits: factorsPerDigits}, nil
}

func (vc *ValidationChain) inputValidattion(nhs string) error{
	// Check length
	if len(nhs) != 10 {
		return ErrBadInput 
	}

	// Check if all chars are digits
	for _, char := range nhs {
		if !unicode.IsDigit(char) {
			return ErrBadInput
		}
	}
	return nil
}

func (vc *ValidationChain) Validate(nhs string)(bool, error){
	// Input validation
	var err error

	err = vc.inputValidattion(nhs)
	if err != nil {
		return false, err
	}

	// var isNHSValid bool
	var validityDigit int
	var checkDigit int

	validityDigit, err = strconv.Atoi(nhs[len(nhs)-1:])
	if err != nil {
		return false, err
	}

	// Multiply(stage 1) + sum(stage 2)
	weightedSum, err := vc.ProcessMap(nhs)
	if err != nil {
		return false, err
	}

	// Getting the reminder using modulo 11
	remainder, err := vc.GetRemainder(weightedSum)
	if err != nil {
		return false, err
	}

	checkDigit = 11 - remainder

	// fmt.Printf("CheckDigit after substraction is %d\n", checkDigit)
	if checkDigit == 11 {
		checkDigit = 0
	} else if checkDigit == 10 {
		return false, ErrInvalidNHS
	}

	if checkDigit != validityDigit {
		return false, ErrInvalidNHS
	} else {
		return true, nil
	}	
}


