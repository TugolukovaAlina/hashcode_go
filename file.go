package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getParams(params []string) (paramsInt []int) {

	paramsInt = make([]int, len(params))
	for i := range params {
		paramsInt[i], _ = strconv.Atoi(params[i])
	}
	return
}
func ReadFile(fileName string) {

	fileHandle, _ := os.Open(fileName)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	// first line : rows, columns, vehicles, rides, bonus, steps
	fileScanner.Scan()
	params := getParams(strings.Split(fileScanner.Text(), " "))

	cars = make([]Car, params[2])
	for j := range cars {
		cars[j].id = j
	}

	routes = []Route{}
	//routes = make([]Route, params[3])
	bonus = params[4]
	endTime = params[5]

	i := 0
	for fileScanner.Scan() {
		// x_start y_start x_end y_end earlist_start latest_finish
		paramsLine := getParams(strings.Split(fileScanner.Text(), " "))
		route := Route{
			startPoint:    Point{x: paramsLine[0], y: paramsLine[1]},
			finishPoint:   Point{x: paramsLine[2], y: paramsLine[3]},
			earliestStart: paramsLine[4],
			latestEnd:     paramsLine[5],
			id:            i,
		}
		//fmt.Println(route)
		routes = append(routes, route)
		i++
	}

}

func WriteFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range cars {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
