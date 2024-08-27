package lessson_04_arrays_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
