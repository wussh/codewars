package kata_test

import (
	. "codewarrior/kata"
	"fmt"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PositiveSum", func() {
	Describe("Fixed Tests", func() {
		for i, p := range basicTests {
			seq, out := p.in, p.out
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := PositiveSum(seq)
				if ans != out {
					fmt.Printf("<LOG::Input>%#v\n", seq)
				}
				Expect(ans).To(Equal(out))
			})
		}
	})

	Describe("Random Tests", func() {
		for i := 0; i < 50; i += 1 {
			seq := randArray(randInt(10, 20))
			ref := refSolution(seq)
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := PositiveSum(seq)
				if ans != ref {
					fmt.Printf("<LOG::Input>%#v\n", seq)
				}
				Expect(ans).To(Equal(ref))
			})
		}
	})
})

type testCase struct {
	in  []int
	out int
}

var basicTests = []testCase{
	{[]int{1, 2, 3, 4, 5}, 15},
	{[]int{1, -2, 3, 4, 5}, 13},
	{[]int{}, 0},
	{[]int{-1, -2, -3, -4, -5}, 0},
	{[]int{-1, 2, 3, 4, -5}, 9},
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randInt(a, b int) int { // assumes a <= b
	return r.Intn(b-a) + a
}

func randArray(size int) []int {
	xs := make([]int, size)
	for i := 0; i < size; i += 1 {
		xs[i] = randInt(-100, 100)
	}
	return xs
}

func refSolution(numbers []int) int {
	var result int
	for _, value := range numbers {
		if value > 0 {
			result += value
		}
	}
	return result
}
