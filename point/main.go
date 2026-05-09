package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func writeStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func writeNum(n int) {
	if n < 0 {
		writeStr("-")
		n = -n
	}
	if n >= 10 {
		writeNum(n / 10)
	}
	d := n % 10

	// start from '0' and increment to get the correct digit rune
	var ch rune = '0'
	i := 0
	for i < d {
		ch++
		i++
	}
	z01.PrintRune(ch)
}

func main() {
	points := &point{}
	setPoint(points)

	writeStr("x = ")
	writeNum(points.x)
	writeStr(", y = ")
	writeNum(points.y)
	writeStr("\n")
}
