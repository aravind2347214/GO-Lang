// You're developing a weather monitoring system in GoLang for a research
// institute. The system needs to store data about temperature, humidity, and wind
// speed. Define variables to hold these measurements for a specific location at a
// given point in time. Ensure you choose suitable data types for representing
// numerical measurements accurately.

package main

import (
	"fmt"
)

// WeatherData struct represents weather measurements for a specific location.
type WeatherData struct {
	Location    string
	Temperature float64 // Temperature in Celsius
	Humidity    float64 // Humidity as a percentage
	WindSpeed   float64 // Wind speed in meters per second
}

func main() {
	// Take user input for weather data
	weather := getUserInput()

	// Print the weather data
	fmt.Printf("\nWeather Data for %s:\n", weather.Location)
	fmt.Printf("Temperature: %.2f°C --- [Type : %T]\nHumidity: %.2f%% --- [Type : %T]\nWind Speed: %.2fm/s --- [Type : %T]\n", weather.Temperature, weather.Temperature, weather.Humidity, weather.Humidity, weather.WindSpeed, weather.WindSpeed)
}

// getUserInput takes user input for weather data and returns a WeatherData struct.
func getUserInput() WeatherData {
	var weather WeatherData

	// Take user input for location
	weather.Location = getStringInput("Enter Location: ")

	// Take user input for temperature
	weather.Temperature = getFloatInput("Enter Temperature (in Celsius): ")

	// Take user input for humidity
	weather.Humidity = getFloatInput("Enter Humidity (as a percentage): ")

	// Take user input for wind speed
	weather.WindSpeed = getFloatInput("Enter Wind Speed (in meters per second): ")

	return weather
}

// getStringInput takes user input for string values.
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// getFloatInput takes user input for float values.
func getFloatInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}
