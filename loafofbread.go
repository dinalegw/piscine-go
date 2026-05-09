package piscine

func LoafOfBread(str string) string {
	// Convert to runes for safe indexing
	r := []rune(str)
	n := len(r)

	// Count non-space runes
	nonSpaceCount := 0
	for _, ch := range r {
		if ch != ' ' {
			nonSpaceCount++
		}
	}

	// Empty or only-spaces => just newline
	if nonSpaceCount == 0 {
		return "\n"
	}

	// Fewer than 5 non-space chars => invalid
	if nonSpaceCount < 5 {
		return "Invalid Output\n"
	}

	groups := []string{}
	i := 0

	for i < n {
		// build one group of up to 5 non-space characters (ignoring spaces while collecting)
		buf := make([]rune, 0, 5)
		j := i
		for j < n && len(buf) < 5 {
			if r[j] != ' ' {
				buf = append(buf, r[j])
			}
			j++
		}

		// if we didn't collect any non-space chars, break
		if len(buf) == 0 {
			break
		}

		groups = append(groups, string(buf))

		// move i to j (we consumed up to j-1), then skip exactly one original character if present
		i = j
		if i < n {
			i++
		}
	}

	// join groups with single spaces and add newline
	out := ""
	for idx, g := range groups {
		if idx != 0 {
			out += " "
		}
		out += g
	}
	out += "\n"
	return out
}
