package main

import (
	"calculator/filemanager"
	"calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewJob(fm, taxRate)
		go priceJob.Process(doneChans[index])
	}

	for _, err := range doneChans {
		if <-err != nil {
			fmt.Println("Could not process job.")
			fmt.Println(<-err)
			break
		}
	}
}
