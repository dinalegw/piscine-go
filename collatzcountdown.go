package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0
	n := start

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps
}
