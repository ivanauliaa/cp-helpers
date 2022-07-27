package golang

func strReverse(str *string) string {
	strRune := []rune(*str)
	n := len(strRune)

	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		strRune[i], strRune[j] = strRune[j], strRune[i]
	}

	return string(strRune)
}
