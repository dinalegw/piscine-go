package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Extract current values
	va := ***a     // value inside a
	vb := *b       // value inside b
	vc := *******c // value inside c
	vd := ****d    // value inside d

	// Perform the required assignments:
	*******c = va // a -> c
	****d = vc    // c -> d
	*b = vd       // d -> b
	***a = vb     // b -> a
}
