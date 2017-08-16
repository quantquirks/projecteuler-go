//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
package main

import (
	"fmt"
)

const limit = 1000

var products []int

func main() {
	startingNumbers := []int{3, 5}
	sumOfMultiples := FindSumOfMultiples(startingNumbers...)
	sumOfProducts := FindSumOfMultiples(FindProducts(startingNumbers...)...)
	fmt.Println(sumOfMultiples - sumOfProducts)
}

func FindSumOfMultiples(numbers ...int) int {
	sum := 0
	//	var multiples []int
	for _, number := range numbers {
		for i := 0; i < limit; i++ {
			multiple := number * i
			if multiple >= 1000 {
				break
			}
			sum += multiple
		}
	}
	return sum
}

func FindProducts(numbers ...int) []int {
	starter := numbers[0]
	numbers = append(numbers[:0], numbers[1:]...)
	for i := 0; i < len(numbers); i++ {
		products = append(products, starter*numbers[i])
	}
	return products
}
