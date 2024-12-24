package main

import (
	"slices"
	"strconv"
	"strings"
)

func partOne(input []byte) int {
	parts := strings.Split(string(input), "\n\n")
	lines := strings.Split(parts[0], "\n")
	queue := strings.Split(parts[1], "\n")
	store := make(map[string]int)

	for _, line := range lines {
		tmp := strings.Split(line, ":")
		num, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))
		store[tmp[0]] = num
	}

	for len(queue) > 0 {
		line := queue[0]
		queue = queue[1:]

		fields := strings.Fields(line)
		a, b, c, op := fields[0], fields[2], fields[4], fields[1]

		val1, ok1 := store[a]
		val2, ok2 := store[b]

		if !ok1 || !ok2 {
			queue = append(queue, line)
			continue
		}

		switch op {
		case "AND":
			if val1 == 1 && val2 == 1 {
				store[c] = 1
			} else {
				store[c] = 0
			}
		case "OR":
			if val1 == 1 || val2 == 1 {
				store[c] = 1
			} else {
				store[c] = 0
			}
		case "XOR":
			if val1 != val2 {
				store[c] = 1
			} else {
				store[c] = 0
			}
		}
	}

	binary := make([]string, 99)
	zwires := 0

	for k, v := range store {
		if k[0] == 'z' {
			idx, _ := strconv.Atoi(k[1:])
			zwires = max(zwires, idx)
			binary[idx] = strconv.Itoa(v)
		}
	}

	binary = binary[:zwires+1]
	slices.Reverse(binary)
	str := strings.Join(binary, "")
	res, _ := strconv.ParseInt(str, 2, 0)

	return int(res)
}
