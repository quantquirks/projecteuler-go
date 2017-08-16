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
		fmt.Println(products)
	}
	return products
}
