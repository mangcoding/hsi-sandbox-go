package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConvertTemperature converts a temperature from Celsius to both Fahrenheit and Reamur scales.
//
// The function accepts a temperature value in Celsius as a float64 and returns two integer values:
//   - The first return value is the equivalent temperature in Fahrenheit, rounded to the nearest integer.
//   - The second return value is the equivalent temperature in Reamur, rounded to the nearest integer.
//
// Conversion formulas used:
//   - Fahrenheit = (Celsius × 9/5) + 32
//   - Reamur    = Celsius × 4/5
//
// Example:
//   fahrenheit, reamur := ConvertTemperature(25.0)
//   fahrenheit == 77, reamur == 20

func ConvertTemperature(celsius float64) (int, int) {
	fahrenheit := (celsius * 9/5) + 32
	reamur := celsius * 4/5
	return int(fahrenheit), int(reamur)
}


func main() {
	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan suhu dalam Celsius: ")

	/*
	We can use this code to read input from user
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	*/
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// Check if input is numeric
	celsius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input Tidak Valid, hanya menerima angka")
		return
	}

	fahrenheit, reamur := ConvertTemperature(celsius)
	fmt.Printf("Suhu dalam Fahrenheit = %d\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur = %d\n", reamur)
}
