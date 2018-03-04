package main

import "sort"

// Algo1 jjj
func Algo1() {
	//sort routes by earliest date
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].earliestStart < routes[j].earliestStart
	})
	for _, r := range routes {
		//find which car finishes first and assing it the route

		carsCopy := make([]Car, len(cars))
		copy(carsCopy, cars)

		bringsMoney := false
		for i := range carsCopy {
			//if does not bring money - drop
			before := carsCopy[i].totalMoney
			carsCopy[i].addRoute(r)
			after := carsCopy[i].totalMoney

			if after > before {
				bringsMoney = true
			}
		}

		if bringsMoney {

			sort.Slice(carsCopy, func(i, j int) bool {
				return carsCopy[i].finishTime < carsCopy[j].finishTime
			})

			//which car will suffer less if I add route to it
			idCar := carsCopy[0].id
			cars[idCar].addRoute(r)
		}
	}
	evaluateAlgo()
}
