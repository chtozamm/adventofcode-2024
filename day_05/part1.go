package main

import (
	"strconv"
	"strings"
)

func partOne(input []byte) int {

	var (
		sum        = 0
		parts      = strings.Split(string(input), "\n\n")
		rules      = strings.Split(parts[0], "\n")
		updates    = strings.Split(parts[1], "\n")
		orderGraph = make(map[string][]string)
	)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		x, y := parts[0], parts[1]
		orderGraph[x] = append(orderGraph[x], y)
	}

	for _, update := range updates {
		pages := strings.Split(update, ",")
		ordered := true

		// For each page number in an update, check if it's placed before
		// any page number associated with it in the order graph
	checkOrder:
		for i, page := range pages {
			for j := range pages[:i] {
				for _, successor := range orderGraph[page] {
					if pages[j] == successor {
						ordered = false
						break checkOrder
					}
				}
			}
		}

		if ordered {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middle
		}
	}

	return sum
}
