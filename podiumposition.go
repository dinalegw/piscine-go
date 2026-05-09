package piscine

func PodiumPosition(podium [][]string) [][]string {
	i := 0
	j := len(podium) - 1

	for i < j {
		podium[i], podium[j] = podium[j], podium[i]
		i++
		j--
	}

	return podium
}
