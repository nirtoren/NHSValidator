package validator

import (
	"strconv"
	"sync"
)

func multiplyByFactors(nhs string, mapping map[int]int, resultChan chan<- int, wg *sync.WaitGroup, errChan chan error){
	defer wg.Done()
	defer close(errChan)

	// Iterate over the number's digits and apply the corresponding factors
	for i, digitChar := range nhs {
		// Convert the current character to an integer
		digit, err := strconv.Atoi(string(digitChar))
		if err != nil {
			errChan <- err
			return
		}

		// Calculate the factor for the current position
		factor := mapping[i+1]

		// Multiply the digit by the factor and add to the result
		result := digit * factor
		resultChan <- result
	}
	errChan <- nil // In case no error
}

func sumResults(resultChan <-chan int, finalResultChan chan<- int) {
	sum := 0
	for result := range resultChan {
		sum += result
	}
	finalResultChan <- sum
}

func (vc *ValidationChain) ProcessMap(nhs string) (int, error) {

	var err error
	resultChan := make(chan int)
	finalResultChan := make(chan int)
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	
	go multiplyByFactors(nhs, vc.factorsPerDigits, resultChan, &wg, errChan)
	go sumResults(resultChan, finalResultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	err = <-errChan

	return <-finalResultChan, err
}

func (vc *ValidationChain) GetRemainder(number int) (int, error) {
	if number < 0 || number == 0 {
		return 0, ErrDivisionNotAllowed
	}
	remainder := number % 11
	return remainder, nil
}
