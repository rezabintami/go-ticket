package seats

import "math/rand"

func RandomString(n int) string {
	var letterRunes = []rune("ABCDEFGHIJKL")
	var numberRunes = []rune("123456789")

	b := make([]rune, n)
	for i := range b {
		if i == 0 {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		} else {
			b[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	}
	return string(b)
}
