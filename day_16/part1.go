package main

import (
	"bytes"
	"container/heap"
	"math"
)

func partOne(input []byte) int {
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
	heap.Push(pq, &Item{point: start, score: 0, dir: '>'})

	visited := make(map[Point]bool)
	minScore := math.MaxInt32

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		if item.point == end {
			minScore = min(minScore, item.score)
			continue
		}

		if visited[item.point] {
			continue
		}
		visited[item.point] = true

		for _, dir := range directions {
			next := Point{item.point.x + dir.dx, item.point.y + dir.dy}

			if matrix[next.y][next.x] == '#' || visited[next] {
				continue
			}

			turns := turnsCount[[2]rune{item.dir, dir.dir}]
			newScore := item.score + 1 + turns*1000
			heap.Push(pq, &Item{point: next, score: newScore, dir: dir.dir})
		}
	}

	return minScore
}
