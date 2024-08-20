package main

import (
	"bufio"
	"fmt"
	"log"
	nhs "nhsvalidator/nhs"
	"os"
	"strconv"
	"strings"
)



func main(){

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an action: ")
		fmt.Println("1: Generate a new number")
		fmt.Println("2: Validate an existing number")
		fmt.Println("0: Exit")
		fmt.Print("Enter your choice: ")
		fmt.Println()

		userChoiceStr, _ := reader.ReadString('\n')
		userChoiceStr = strings.TrimSpace(userChoiceStr)
		choice, err := strconv.Atoi(userChoiceStr)
		if err != nil || (choice != 1 && choice != 2 && choice != 0) {
			fmt.Println("Invalid choice. Please enter 1, 2, or 0.")
			continue
		}

		NHS, err := nhs.NewNHSManager()
			if err != nil {
				log.Fatal("Could not instansiate NHS")
			}

		switch choice {
		case 1:
			generatedNumber := NHS.Generate() // NHS generation process entry-point
			fmt.Println(generatedNumber)
			
		case 2:
			fmt.Print("Enter the 10-digit number to validate: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			isValid, err := NHS.Validate(number) // NHS validation process entry-point
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(isValid)
			}
			
		case 0:
			fmt.Println("Exiting...")
			return
		}
	}

}