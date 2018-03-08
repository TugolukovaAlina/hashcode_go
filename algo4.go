package main

import (
	"math"
)

func Algo4() {

	carsCopy := make([]Car, len(cars))
	copy(carsCopy, cars)

	//fmt.Println(routes)
	for len(routes) > 0 {

		minPotentialStart := math.MaxInt64
		maxPotentialMoney := 0
		idRoute := -1
		idCar := -1

		for i := range cars {
			for j, r := range routes {
				car := cars[i]
				potentialStart, potentialMoney, _ := car.addPotentialRoute(r)

				if potentialStart < minPotentialStart && potentialMoney > 0 {
					minPotentialStart = potentialStart
					maxPotentialMoney = potentialMoney
					idRoute = j
					idCar = i
				} else if potentialStart == minPotentialStart && potentialMoney > maxPotentialMoney {
					maxPotentialMoney = potentialMoney
					idRoute = j
					idCar = i
				}

			}
		}

		if idRoute != -1 {
			cars[idCar].addRoute(routes[idRoute])
			routes = append(routes[:idRoute], routes[idRoute+1:]...)
		}
	}
	evaluateAlgo()
}
