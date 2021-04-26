package iteration

func Repeat(char string) string {
	var res string
	for i := 0; i < 5; i++ {
		res += char
	}
	return res
}

