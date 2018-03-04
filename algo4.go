package main

import (
	"math"
	"sort"
)

// Algo1 jjj
func Algo4() {
	//sort routes by earliest date
	sort.Slice(routes, func(i, j int) bool {
		//return routes[i].getDistance() > routes[j].getDistance()
		return routes[i].earliestStart < routes[j].earliestStart
	})
	for _, r := range routes {
		//find which car finishes first and assing it the route

		carsCopy := make([]Car, len(cars))
		copy(carsCopy, cars)

		bringsMoney := false

		minDiffTime := math.MaxInt16
		minDiffTimeId := 0

		for i := range carsCopy {
			//if does not bring money - drop
			before := carsCopy[i].totalMoney
			beforeTime := carsCopy[i].finishTime
			carsCopy[i].addRoute(r)

			after := carsCopy[i].totalMoney
			afterTime := carsCopy[i].finishTime

			if (afterTime - beforeTime) < minDiffTime {
				minDiffTimeId = carsCopy[i].id
				minDiffTime = afterTime - beforeTime
			}
			if after > before {
				bringsMoney = true
			}
		}

		if bringsMoney {
			cars[minDiffTimeId].addRoute(r)
		}
	}
	evaluateAlgo()
}
