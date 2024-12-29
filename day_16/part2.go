package main

import (
	"bytes"
	"container/heap"
)

func partTwo(input []byte) int {
	matrix := bytes.Split(input, []byte("\n"))

	var start, end Point

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 'S' {
				start = Point{x, y}
			} else if matrix[y][x] == 'E' {
				end = Point{x, y}
			}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: start, score: 0, dir: '>', path: []Point{start}})

	type state struct {
		point Point
		dir   rune
	}

	prevPositions := make(map[state]int)
	bestTiles := make(map[Point]bool)
	minScore := -1

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		if minScore != -1 && minScore < item.score {
			continue
		}

		if item.point == end {
			minScore = item.score
			for _, tile := range item.path {
				bestTiles[tile] = true
			}
			continue
		}

		if prevPositions[state{item.point, item.dir}] != 0 &&
			prevPositions[state{item.point, item.dir}] < item.score {
			continue
		}
		prevPositions[state{item.point, item.dir}] = item.score

		for _, dir := range directions {
			next := Point{item.point.x + dir.dx, item.point.y + dir.dy}

			if matrix[next.y][next.x] == '#' {
				continue
			}

			turns := turnsCount[[2]rune{item.dir, dir.dir}]
			newScore := item.score + 1 + turns*1000
			newPath := append([]Point{}, item.path...)
			newPath = append(newPath, next)
			heap.Push(pq, &Item{point: next, score: newScore, dir: dir.dir, path: newPath})
		}
	}

	return len(bestTiles)
}
