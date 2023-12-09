package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludePriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludePriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func (job *TaxIncludePriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("could not read file")
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("reading the file failed")
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("string conversion to float failed")
			file.Close()
			return
		}

		prices = append(prices, floatPrice)
	}

	job.InputPrices = prices
}

func NewJob(taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
