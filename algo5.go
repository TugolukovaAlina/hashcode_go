package main

import (
	"math"
	"sort"
)

func Algo5() {

	carsCopy := make([]Car, len(cars))
	copy(carsCopy, cars)

	//fmt.Println(routes)
	for len(carsCopy) > 0 {
		//find which car finishes first and assing it the route

		sort.Slice(carsCopy, func(i, j int) bool { return carsCopy[i].finishTime < carsCopy[j].finishTime })

		idCar := carsCopy[0].id

		//try choose n with min potential start and then choose the biggest money
		minPotentialStart := math.MaxInt64
		maxPotentialMoney := 0
		//maxCost := 0.0
		idRoute := -1
		for id, r := range routes {

			car := carsCopy[0]
			potentialStart, potentialMoney, _ := car.addPotentialRoute(r)
			//_, _, cost := car.addPotentialRoute(r)

			/*
				if maxCost < cost && cost > 0 {
					maxCost = cost
					idRoute = id
				}*/

			if potentialStart < minPotentialStart && potentialMoney > 0 {
				minPotentialStart = potentialStart
				maxPotentialMoney = potentialMoney
				idRoute = id
			} else if potentialStart == minPotentialStart && potentialMoney > maxPotentialMoney {
				maxPotentialMoney = potentialMoney
				idRoute = id
			}

			//check if there is other car that can start earlier
		}
		if idRoute != -1 {
			cars[idCar].addRoute(routes[idRoute])
			carsCopy[0].addRoute(routes[idRoute])
			routes = append(routes[:idRoute], routes[idRoute+1:]...)
		} else {
			//delete car if no route can be added
			carsCopy = carsCopy[1:]
		}
	}
	evaluateAlgo()
}
