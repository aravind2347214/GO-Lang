package main

import "fmt"

// This program is Made By Aravind Nandakumar (2347214)

// Problem Statement:
// You are tasked with creating a temperature converter program in Go. The program should allow users to convert temperatures between Celsius, Fahrenheit, and Kelvin.

// 1. Celsius to Fahrenheit
// 2. Fahrenheit to Celsius
// 3. Celsius to Kelvin
// 4. Kelvin to Celsius
// 5. Fahrenheit to Kelvin
// 6. Kelvin to Fahrenheit

// function to convert  celsius to fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// function to convert    fahrenheit to celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// function to convert    celsius to kelvin
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// function to convert   kelvin to celsius
func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

//function to convert   fahrenheit to kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

// function to convert    kelvin to fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

func main() {

	// choice variable
	var ch int

	// user input temperature without unit
	var temperature float64

	// menu
	fmt.Println("Temperature Converter ")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Celsius to Kelvin")
	fmt.Println("4. Kelvin to Celsius")
	fmt.Println("5. Fahrenheit to Kelvin")
	fmt.Println("6. Kelvin to Fahrenheit")

	// user prompts
	fmt.Print("Enter your choice (1-6): ")
	fmt.Scanln(&ch)

	fmt.Print("Enter the temperature value: ")
	fmt.Scanln(&temperature)

	// A switch case to choose which action we want to perform
	switch ch {

	//Case to convert celcius to fahrenhiet
	case 1:
		result := celsiusToFahrenheit(temperature)
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temperature, result)

	//Case to convert  fahrenhiet to celcius
	case 2:
		result := fahrenheitToCelsius(temperature)
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temperature, result)

	//Case to convert celcius to Kelvin
	case 3:
		result := celsiusToKelvin(temperature)
		fmt.Printf("%.2f Celsius is %.2f Kelvin\n", temperature, result)

	//Case to convert  Kelvin to celcius
	case 4:
		result := kelvinToCelsius(temperature)
		fmt.Printf("%.2f Kelvin is %.2f Celsius\n", temperature, result)

	//Case to convert  Fahrenheit to Kelvin
	case 5:
		result := fahrenheitToKelvin(temperature)
		fmt.Printf("%.2f Fahrenheit is %.2f Kelvin\n", temperature, result)

	//Case to convert  Kelvin to fahrenheit
	case 6:
		result := kelvinToFahrenheit(temperature)
		fmt.Printf("%.2f Kelvin is %.2f Fahrenheit\n", temperature, result)
	default:
		fmt.Println("Invalid ch. Please select a valid option.")
	}

}
