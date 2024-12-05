package main

import (
	"strconv"
	"strings"
)

func partTwo(input []byte) int {

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

	// For each page number in an update, check if it's placed before
	// any page number associated with it in the order graph
	for _, update := range updates {
		pages := strings.Split(update, ",")
		ordered := true

		for i, page := range pages {
			for j := range pages[:i] {
				for _, successor := range orderGraph[page] {
					if pages[j] == successor {
						ordered = false
						// Swap pages to restore the order
						pages[j], pages[i] = pages[i], pages[j]
						break
					}
				}
			}
		}

		// Add middle page number to the sum only if the update was incorrectly-ordered
		if !ordered {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middle
		}
	}

	return sum
}
