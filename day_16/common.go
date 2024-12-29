package main

type Point struct {
	x, y int
}

var turnsCount = map[[2]rune]int{
	{'^', '<'}: 1, {'^', 'v'}: 2, {'^', '>'}: 1,
	{'<', 'v'}: 1, {'<', '>'}: 2, {'<', '^'}: 1,
	{'v', '>'}: 1, {'v', '^'}: 2, {'v', '<'}: 1,
	{'>', '^'}: 1, {'>', '<'}: 2, {'>', 'v'}: 1,
}

var directions = []struct {
	dir rune
	dx  int
	dy  int
}{
	{'^', 0, -1},
	{'>', 1, 0},
	{'v', 0, 1},
	{'<', -1, 0},
}
