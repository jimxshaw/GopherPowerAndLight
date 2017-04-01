package main

import (
	"fmt"
)

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 60, 100}

	var capacity = plantCapacities[0] + plantCapacities[1] + plantCapacities[2]

	var gridLoad = 75.

	utilization := gridLoad / capacity

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", utilization*100)
}
