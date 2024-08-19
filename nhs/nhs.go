package nhs

import (
	"log"
	"math/rand"
	"nhsvalidator/nhs/validator"
	"strconv"
)


type NHS struct{}


func NewNHSManager() (*NHS, error){
	return &NHS{}, nil
}

func (n *NHS) Generate() (string, error) {
	validationChain, err := validator.NewValidationChain()
	if err != nil {
		log.Fatalf("Validation chain could not been established %v", err)
	}
	var NHSNumber string
	for {
		var checkDigit int

		NHSNumber = ""
		for i := 0; i < 9; i++ {
			digit := rand.Intn(10) // Random digit between 0 and 9
			NHSNumber += strconv.Itoa(digit)
		}

		weightedSum, err := validationChain.ProcessMap(NHSNumber)
		if err != nil {
			continue
			// return "", err
		}

		remainder, err := validationChain.GetRemainder(weightedSum)
		if err != nil {
			continue
		}

		checkDigit = 11 - remainder

		if checkDigit == 11 {
			checkDigit = 0
		} else if checkDigit == 10 { // Indicate an invalid number
			continue 
		}

		NHSNumber += strconv.Itoa(checkDigit)

		// Validation of generated number
		isValid, _ := validationChain.Validate(NHSNumber)
		if isValid {
			break
		}
	}
	
	return NHSNumber, err
}

func (n *NHS) Validate(nhs string) (bool, error) {
	// Init validation chain
	validationChain, err := validator.NewValidationChain()
	if err != nil {
		log.Fatalf("Validation chain could not been established %v", err)
	}

	// Entry-point for the validation process
	isNHSValid, err := validationChain.Validate(nhs)

	return isNHSValid, err
}