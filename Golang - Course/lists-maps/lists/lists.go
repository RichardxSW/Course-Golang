package lists

import "fmt"

// type Product struct {
// 	title string
// 	id string
// 	price float64
// }

func main() {
	prices := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(prices)
	featuredPrices := prices[1:]
	featuredPrices[0] = 119.9
	fmt.Println(featuredPrices)
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	discountPrices := []float64{5.5, 6.6}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
