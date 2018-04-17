package main

import (
	"fmt"
)

func reduce(elements []float64, callback func(float64) float64) float64 {
	memo := 0.0
	for _, value := range elements {
		memo += callback(value)
	}
	return memo
}

func average(params ...float64) (float64, string) {
	var total float64
	for _, value := range params {
		total += value
	}
	return total / float64(len(params)), "ok"
}

func filterFactory(callback func(float64) bool) func([]float64) []float64 {
	return func(numbers []float64) []float64 {
		var filtered []float64
		for _, value := range numbers {
			if callback(value) {
				filtered = append(filtered, value)
			}
		}
		return filtered
	}
}

func genericFunc(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println(x, "float64")
	default:
		fmt.Println("unknown")
	}
}

func funcs() {
	data := []float64{43, 2, 34, 35, 65}
	val, _ := average(data...)

	numFilter := filterFactory(func(num float64) bool {
		if num > 4 {
			return true
		}
		return false
	})

	reduced := reduce(numFilter(data), func(num float64) float64 {
		return val + num
	})

	genericFunc(reduced)
}
