package main

import (
	"fmt"

	"github.com/kumarharsh13/price-calculator/fileManager"
	"github.com/kumarharsh13/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	// results := make(map[float64][]float64)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		// pricesJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go pricesJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process Job")
		// 	fmt.Println(err)
		// }

	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done...!!!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }
	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
