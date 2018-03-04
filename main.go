package main

import (
	"fmt"
	"strconv"
)

var cars []Car
var routes []Route
var bonus int
var endTime int

type Point struct {
	x, y int
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(p1, p2 Point) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
}

type Route struct {
	startPoint, finishPoint  Point
	earliestStart, latestEnd int
	id                       int
}

func (r Route) getDistance() int {
	return distance(r.startPoint, r.finishPoint)
}

func (r *Route) String() string {
	return fmt.Sprint(r.earliestStart)
}

type Car struct {
	assignRoutes []Route
	finishTime   int   //time when car will be free
	finishPoint  Point //point in which taxi will finish

	totalMoney int
	id         int
}

func (c Car) String() string {
	s := fmt.Sprint(c.id, " ")
	for _, r := range c.assignRoutes {
		s += fmt.Sprint(r.id, " ")
	}
	s += fmt.Sprint(c.finishTime, " ")
	return s
}

func (c *Car) addRoute(r Route) {
	c.assignRoutes = append(c.assignRoutes, r)
	//time to go to beginning of route plus distance of route

	//arrival time at the beginning of the route
	taxiReadyTime := c.finishTime + distance(c.finishPoint, r.startPoint)
	taxiDepartureTime := taxiReadyTime
	if taxiReadyTime <= r.earliestStart {
		c.totalMoney += bonus
		taxiDepartureTime = r.earliestStart
	}

	//taxiDepartureTime := math.Max(taxiReadyTime, r.earliestStart)

	c.finishTime = taxiDepartureTime + r.getDistance()

	if c.finishTime < r.latestEnd {
		c.totalMoney += r.getDistance()
	}

	c.finishPoint = r.finishPoint
}

func (c *Car) addRoute2(r Route) {
	c.assignRoutes = append(c.assignRoutes, r)
	//time to go to beginning of route plus distance of route

	//arrival time at the beginning of the route
	taxiReadyTime := c.finishTime + distance(c.finishPoint, r.startPoint)
	taxiDepartureTime := taxiReadyTime
	if taxiReadyTime <= r.earliestStart {
		c.totalMoney += bonus
		taxiDepartureTime = r.earliestStart
	}

	//taxiDepartureTime := math.Max(taxiReadyTime, r.earliestStart)

	c.finishTime = taxiDepartureTime + r.getDistance()

	if c.finishTime < r.latestEnd {
		c.totalMoney += r.getDistance()
	} else {
		fmt.Println("no bonus ", r.id, " earliest time ", r.earliestStart, " distance to it ", distance(Point{0, 0}, r.startPoint))
	}

	c.finishPoint = r.finishPoint
}

var totalScore int

func evaluateAlgo() {
	score := 0
	for _, c := range cars {
		score += c.totalMoney
	}
	fmt.Println(score)
	totalScore += score
}

//func FinishTimeComparator(i, j int) bool
func main() {

	totalScore = 0
	/*
		files := []string{"a_example.in", "b_should_be_easy.in", "c_no_hurry.in", "d_metropolis.in", "e_high_bonus.in"}
		algos := []func(){Algo1, Algo2, Algo3, Algo5}
		description := []string{"mytimest+", "timest", "mytimest", "shortest"}
	*/
	files := []string{"a_example.in", "b_should_be_easy.in", "c_no_hurry.in", "d_metropolis.in", "e_high_bonus.in"}
	algos := []func(){Algo5}
	description := []string{"best_result"}

	for i, al := range algos {
		totalScore = 0
		for _, f := range files {
			ReadFile(f)
			al()
			WriteFile("results/" + description[i] + strconv.Itoa(i) + f)
		}
		fmt.Println("total score ", description[i], " ", totalScore)
	}

}
