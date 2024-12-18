package main

import (
	"container/heap"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Item struct {
	point   Point
	pathLen int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].pathLen < pq[j].pathLen
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func partOne(input []byte, size int, amount int) int {
	corruptedBytes := strings.Split(string(input), "\n")[:amount]
	corruptedBytesMap := make(map[Point]bool, amount)
	exit := Point{size - 1, size - 1}

	for _, c := range corruptedBytes {
		sides := strings.Split(c, ",")
		x, _ := strconv.Atoi(sides[0])
		y, _ := strconv.Atoi(sides[1])
		corruptedBytesMap[Point{x, y}] = true
	}

	directions := []Point{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: Point{0, 0}, pathLen: 0})

	visited := make(map[Point]bool)
	shortestPath := math.MaxInt32

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		if item.point == exit {
			shortestPath = min(shortestPath, item.pathLen)
			continue
		}

		if visited[item.point] {
			continue
		}
		visited[item.point] = true

		for _, dir := range directions {
			next := Point{item.point.x + dir.x, item.point.y + dir.y}

			if next.x < 0 || next.x >= size || next.y < 0 || next.y >= size {
				continue
			}

			if corruptedBytesMap[next] || visited[next] {
				continue
			}

			heap.Push(pq, &Item{point: next, pathLen: item.pathLen + 1})
		}
	}

	return shortestPath
}
