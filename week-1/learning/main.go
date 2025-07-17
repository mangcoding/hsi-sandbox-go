package main

import "fmt"
import "errors"

func devide(a,b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// Basic syntax and data types

	var name string = "Nugraha"
	var birthYear int = 1989
	var isMarried bool = true
	
	// short declaration operator
	age := 2025 - birthYear

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Married:", isMarried)

	// constants
	const pi float64 = 3.14
	fmt.Println("Pi:", pi)

	// Basic control flow
	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Basic loop
	for i := 0; i < 10; i++ {
		fmt.Println("Loop:", i)
	}

	// sample call function
	result, err := devide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
