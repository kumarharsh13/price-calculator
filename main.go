package main

import (
	"fmt"

	"github.com/kumarharsh13/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(taxRate)
		pricesJob.Process()
	}
	fmt.Println(results)
}
