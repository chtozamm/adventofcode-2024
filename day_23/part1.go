package main

import (
	"fmt"
	"sort"
	"strings"
)

func partOne(input []byte) int {
	lines := strings.Split(string(input), "\n")
	pairs := make([][]string, len(lines))

	for i, line := range lines {
		pair := strings.Split(line, "-")
		pairs[i] = []string{pair[0], pair[1]}
	}

	connections := make(map[string]map[string]bool)

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if connections[a] == nil {
			connections[a] = make(map[string]bool)
		}
		if connections[b] == nil {
			connections[b] = make(map[string]bool)
		}
		connections[a][b] = true
		connections[b][a] = true
	}

	triangles := make(map[string]struct{})

	for a := range connections {
		for b := range connections[a] {
			for c := range connections[b] {
				if c != a && connections[a][c] {
					if a[0] != 't' && b[0] != 't' && c[0] != 't' {
						continue
					}
					triangle := []string{a, b, c}
					sort.Strings(triangle)
					triangleKey := fmt.Sprintf("%s,%s,%s", triangle[0], triangle[1], triangle[2])
					triangles[triangleKey] = struct{}{}
				}
			}
		}
	}

	return len(triangles)
}
