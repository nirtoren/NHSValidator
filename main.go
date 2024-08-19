package main

import (
	"fmt"
	"log"
	nhs "nhsvalidator/nhs"
)



func main(){


	NHS, err := nhs.NewNHSManager()
	if err != nil {
		log.Fatal("Could not instansiate NHS")
	}

	// nhs := "5990128088" // Should be valid
	// nhs := "1275988113" // Should be valid
	// nhs := "4536026665" // Should be valid
	nhs := "5990128087" // Should be INVALID
	// nhs := "4536016660" // Should be INVALID
	
	// generatedNHS := "9168746121"

	isValid, err := NHS.Validate(nhs)
	fmt.Println(isValid)

	// newNHS, err := NHS.Generate()
	// fmt.Println(newNHS)

	// should be readable, testable, organized correctly.
}