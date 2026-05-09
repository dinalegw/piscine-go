package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for {
		if IsPrime(nb) { // Use the IsPrime from isprime.go
			return nb
		}
		nb++
	}
}
