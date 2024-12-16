package main

import (
	"bytes"
	"container/heap"
	"math"
)

type Point struct {
	x, y int
}

type Item struct {
	point Point
	score int
	dir   rune
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

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

	turnsCount := map[[2]rune]int{
		{'^', '<'}: 1, {'^', 'v'}: 2, {'^', '>'}: 1,
		{'<', 'v'}: 1, {'<', '>'}: 2, {'<', '^'}: 1,
		{'v', '>'}: 1, {'v', '^'}: 2, {'v', '<'}: 1,
		{'>', '^'}: 1, {'>', '<'}: 2, {'>', 'v'}: 1,
	}

	directions := []struct {
		dir rune
		dx  int
		dy  int
	}{
		{'^', 0, -1}, {'>', 1, 0}, {'v', 0, 1}, {'<', -1, 0},
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

			if next.x < 0 || next.x >= len(matrix[0]) || next.y < 0 || next.y >= len(matrix) {
				continue
			}
			if matrix[next.y][next.x] == '#' || visited[next] {
				continue
			}

			// Calculate the new score and push the next item to the queue
			turns := turnsCount[[2]rune{item.dir, dir.dir}]
			newScore := item.score + 1 + turns*1000
			heap.Push(pq, &Item{point: next, score: newScore, dir: dir.dir})
		}
	}

	return minScore
}
