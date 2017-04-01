package main

import "fmt"

func main() {
	plantCapacities := []float64{30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	generateReport(plantCapacities, activePlants, gridLoad)

}

func generateReport(plantCapacities []float64, activePlants []int, gridLoad float64) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities...)
	case "2":
		generatePowerGridReport(activePlants, plantCapacities, gridLoad)
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
