package kata_test

import (
	. "codewarrior/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strings"
	"time"
)

func dotest(text, char, expected string) {
	Expect(Contamination(text, char)).To(Equal(expected), "With text = \"%s\", char = \"%s\"", text, char)
}

func referenceSolution(a, b string) string {
	return strings.Repeat(b, len(a))
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest("abc", "z", "zzz")
		dotest("", "z", "")
		dotest("abc", "", "")
		dotest("_3ebzgh4", "&", "&&&&&&&&")
		dotest("//case", " ", "      ")
	})
	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		var text, char string
		for i := 0; i < 100; i++ {
			arr := make([]rune, rand.Intn(31))
			for j := range arr {
				arr[j] = rune(32 + rand.Intn(95))
			}
			text = string(arr)
			if rand.Intn(20) > 0 {
				char = string(rune(32 + rand.Intn(95)))
			} else {
				char = ""
			}
			dotest(text, char, referenceSolution(text, char))
		}
	})
})
