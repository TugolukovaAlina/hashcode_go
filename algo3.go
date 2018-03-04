package main

import (
	"sort"
)

func Algo3() {
	//sort routes by earliest date
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].earliestStart < routes[j].earliestStart
	})
	//fmt.Println(routes)
	for _, r := range routes {
		//find which car finishes first and assing it the route

		carsCopy := make([]Car, len(cars))
		copy(carsCopy, cars)

		for i := range carsCopy {
			carsCopy[i].addRoute(r)
		}

		//sort.Sort(ByFinishTime(carsCopy))

		sort.Slice(carsCopy, func(i, j int) bool { return carsCopy[i].finishTime < carsCopy[j].finishTime })

		//which car will suffer less if I add route to it
		idCar := carsCopy[0].id

		cars[idCar].addRoute(r)
	}
	evaluateAlgo()
}
