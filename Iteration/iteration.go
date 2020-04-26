package iteration

func Repeat(character string, ni int) string {
	var repeated string
	for i := 0; i < ni; i++ {
		repeated += character
	}

	return repeated
}
