package main

import "fmt"

func main() {
	var distance uint16
	const minDistance uint16 = 50
	const maxDistance uint16 = 10000
	fmt.Printf("Введите расстояние от %d до %d км.\n", minDistance, maxDistance)
	fmt.Scan(&distance)

	if distance < minDistance || distance > maxDistance {
		fmt.Printf("Расстояние должно быть от %d до %d.\n", minDistance, maxDistance)
		return
	}

	var consumption uint8
	const minConsumption uint8 = 5
	const maxConsumption uint8 = 25
	fmt.Printf("Введите расход в литрах на 100 км от %d до %d км.\n", minConsumption, maxConsumption)
	fmt.Scan(&consumption)

	if consumption < minConsumption || consumption > maxConsumption {
		fmt.Printf("Расход в литрах должен быть от %d до %d.\n", minConsumption, maxConsumption)
		return
	}

	const gasolineCost float32 = 48.0
	tripCost := float32(distance) / 100 * float32(consumption) * gasolineCost
	fmt.Printf("Стоимость поездки: %.2f", tripCost)
}
