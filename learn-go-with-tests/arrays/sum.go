package arrays

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	var sums []int
	for _, v := range arrs {
		sums = append(sums, Sum(v))
	}
	return sums
}

func SumAllTails(arrs ...[]int) []int {

	var sums []int
	for _, arr := range arrs {
		if len(arr) == 0 {
			sums = append(sums, 0)

		} else {
			sums = append(sums, Sum(arr[1:]))
		}
	}

	return sums
}
