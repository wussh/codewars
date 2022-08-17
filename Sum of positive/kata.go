package kata

func PositiveSum(numbers []int) int {
	sum := 0
	neg := 0
	for _, nums := range numbers {
		if nums > 0 {
			sum += nums
		} else if nums < 0 {
			neg += nums
		} else if numbers == nil {
			return 0
		}
	}
	return sum
}
