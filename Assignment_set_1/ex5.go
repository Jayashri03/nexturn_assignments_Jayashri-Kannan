package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name         string
	Temperature  float64
	Rainfall     float64
}

func main() {
	cities := []City{
		{"New York", 15.0, 1200.0},
		{"Los Angeles", 20.0, 800.0},
		{"Chicago", 10.0, 1000.0},
		{"Houston", 25.0, 1400.0},
		{"Phoenix", 30.0, 200.0},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nClimate Data Analysis Menu:")
		fmt.Println("1. View city with highest temperature")
		fmt.Println("2. View city with lowest temperature")
		fmt.Println("3. Calculate average rainfall")
		fmt.Println("4. Filter cities by rainfall")
		fmt.Println("5. Search city by name")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 6.")
			continue
		}

		switch option {
		case 1:
			highestTempCity := findCityWithHighestTemperature(cities)
			fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highestTempCity.Name, highestTempCity.Temperature)
		case 2:
			lowestTempCity := findCityWithLowestTemperature(cities)
			fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowestTempCity.Name, lowestTempCity.Temperature)
		case 3:
			avgRainfall := calculateAverageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", avgRainfall)
		case 4:
			fmt.Print("Enter rainfall threshold (mm): ")
			thresholdInput, _ := reader.ReadString('\n')
			thresholdInput = strings.TrimSpace(thresholdInput)
			threshold, err := strconv.ParseFloat(thresholdInput, 64)
			if err != nil {
				fmt.Println("Invalid input, please enter a numeric value.")
				continue
			}
			filterCitiesByRainfall(cities, threshold)
		case 5:
			fmt.Print("Enter city name: ")
			cityName, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityName)
			city, err := searchCityByName(cities, cityName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("City: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}
		case 6:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please choose between 1 and 6.")
		}
	}
}

func findCityWithHighestTemperature(cities []City) City {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

func findCityWithLowestTemperature(cities []City) City {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

func calculateAverageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	found := false
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
			found = true
		}
	}
	if !found {
		fmt.Println("No cities found with rainfall above the threshold.")
	}
}

func searchCityByName(cities []City, name string) (City, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return City{}, fmt.Errorf("City '%s' not found.", name)
}
