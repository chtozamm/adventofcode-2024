package main

import (
	"strings"
)

func partOne(input []byte) int {
	schematics := strings.Split(string(input), "\n\n")
	locks := make([]string, 0)
	keys := make([]string, 0)

	for _, schematic := range schematics {
		if strings.HasPrefix(schematic, "#####") {
			locks = append(locks, schematic)
		} else {
			keys = append(keys, schematic)
		}
	}

	const lockHeight = 6
	const columns = 5

	locksPinsHeights := make([][]int, 0)
	keysTeethHeights := make([][]int, 0)

	for _, lock := range locks {
		lines := strings.Split(lock, "\n")[1:]
		heights := make([]int, columns)
		for i := 0; i < columns; i++ {
			pinHeight := 0
			for j := 0; j < lockHeight; j++ {
				if lines[j][i] == '#' {
					pinHeight++
				}
			}
			heights[i] = pinHeight
		}
		locksPinsHeights = append(locksPinsHeights, heights)
	}

	for _, key := range keys {
		lines := strings.Split(key, "\n")[:lockHeight]
		heights := make([]int, columns)
		for i := 0; i < columns; i++ {
			toothHeight := 0
			for j := lockHeight - 1; j >= 0; j-- {
				if lines[j][i] == '#' {
					toothHeight++
				}
			}
			heights[i] = toothHeight
		}
		keysTeethHeights = append(keysTeethHeights, heights)
	}

	pairs := 0

	for _, lock := range locksPinsHeights {
	keysLoop:
		for _, key := range keysTeethHeights {
			for i := 0; i < columns; i++ {
				if lock[i]+key[i] > 5 {
					continue keysLoop
				}
			}
			pairs++
		}
	}

	return pairs
}
