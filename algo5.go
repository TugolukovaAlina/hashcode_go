package main

import (
	"sort"
)

func Algo5() {

	carsCopy := make([]Car, len(cars))
	copy(carsCopy, cars)

	//fmt.Println(routes)
	for len(carsCopy) > 0 {
		//find which car finishes first and assing it the route

		sort.Slice(carsCopy, func(i, j int) bool {
			return carsCopy[i].finishTime < carsCopy[j].finishTime
		})

		idCar := carsCopy[0].id

		maxAlpha := 0.0
		idRoute := -1
		for id, r := range routes {

			car := carsCopy[0]
			timeBefore := car.finishTime
			moneyBefore := car.totalMoney
			car.addRoute(r)
			timeAfter := car.finishTime
			moneyAfter := car.totalMoney

			alpha := float64(r.getDistance()) / float64((timeAfter - timeBefore))
			//alpha := float64(timeAfter - timeBefore - r.getDistance())

			//fmt.Println(alpha)

			if alpha > maxAlpha && moneyAfter > moneyBefore {
				maxAlpha = alpha
				idRoute = id
			}
		}
		if idRoute != -1 {
			cars[idCar].addRoute(routes[idRoute])
			carsCopy[0].addRoute(routes[idRoute])
			routes = append(routes[:idRoute], routes[idRoute+1:]...)
		} else {
			//delete car if no route can be added
			carsCopy = carsCopy[1:]
			//carsCopy = append(carsCopy[:idCar], carsCopy[idCar+1:]...)
			//fmt.Println(len(carsCopy))
		}
	}
	evaluateAlgo()
}
