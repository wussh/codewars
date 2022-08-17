package kata

func Contamination(text, char string) string {
	panjang := len(text)
	newtext := ""
	if panjang > 1 {
		for i := 0; i < panjang; i++ {
			newtext += char
		}
	} else {
		return ""
	}
	return newtext
}
