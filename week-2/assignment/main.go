package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type location struct {
	city           string
	temperature    int
	description    string
	tempFahrenheit int
	tempReamur     int
}

func getCategories(temperature int) string {
	if temperature > 25 {
		return "panas"
	} else if temperature > 18 {
		return "hangat"
	} else {
		return "dingin"
	}
}

func ConvertTemperature(celsius int) (int, int) {
	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	return int(fahrenheit), int(reamur)
}

func getWeatherInfo(loc *location) {
	loc.tempFahrenheit, loc.tempReamur = ConvertTemperature(loc.temperature)
	loc.description = getCategories(loc.temperature)

	fmt.Printf("Suhu di %s adalah %s\n", loc.city, loc.description)
	fmt.Printf("Suhu di %s dalam Reamur = %d\n", loc.city, loc.tempReamur)
	fmt.Printf("Suhu di %s dalam Fahrenheit = %d\n", loc.city, loc.tempFahrenheit)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var city location

	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan lokasi pengukuran suhu: ")
	input, _ := reader.ReadString('\n')
	city.city = strings.TrimSpace(input)

	fmt.Print("Masukkan suhu dalam Celsius: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Check if input is numeric
	celsius, err := strconv.Atoi(input)
	city.temperature = celsius

	if err != nil {
		fmt.Println("Input Tidak Valid, hanya menerima angka")
		return
	}

	getWeatherInfo(&city)
}
