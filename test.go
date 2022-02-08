package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("pilot have", fuel)

}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 50000
	case "Mars":
		fuel = 70000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	//fuelGauge()
	//fmt.Println(calculateFuel("Venus"))
	fuel := 1000000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
	fmt.Println(flyToPlanet("Mercury", 600000))

}
