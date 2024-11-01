package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	doubledAgain := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	tripled := transformNumbers(&numbers, triple)
	sum := sumup(10, 5 , 5)

	// pakai ... pada slice agar bisa dikirim ke fungsi
	anotherSum := sumup(numbers...)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(doubledAgain)
	fmt.Println(tripled)
	fmt.Println(sum)
	fmt.Println(anotherSum)

}

// numbers bisa diisi oleh banyak angka seperti 1, 2, 3
func sumup(numbers ...int) int{
	sum:= 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

// func transformNumbers (numbers *[]int, transform func(int) int) []int {
func transformNumbers (numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
// func doubleNumbers (numbers *[]int) []int {
// 	dNumbers := []int{}

// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, val * 2)
// 	}

// 	return dNumbers
// }

// func tripleNumbers (numbers *[]int) []int {
// 	tNumbers := []int{}

// 	for _, val := range *numbers {
// 		tNumbers = append(tNumbers, val * 3)
// 	}

// 	return tNumbers
// }

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3	
}