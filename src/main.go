package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, inactive},
		PowerPlant{hydro, 275, active},
		PowerPlant{solar, 85, inactive},
		PowerPlant{solar, 50, active},
		PowerPlant{wind, 40, active},
		PowerPlant{wind, 25, unavailable},
	}

	grid := PowerGrid{300, plants}

	generateReport(grid)

}

func generateReport(grid PowerGrid) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		grid.generatePlantReport()
	case "2":
		//TODO
	default:
		fmt.Println("Invalid choice. No action taken.")
	}
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for index, capa := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", index, capa)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for index, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", index)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type: ", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", p.capacity)
		fmt.Printf("%-20s%s\n", "Status: ", p.status)
		fmt.Println("")
	}
}
