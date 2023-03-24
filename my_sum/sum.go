package main

// range permite percorrer uma array e retorna dois
// valores sendo um o indice e outro o valor
func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

/*
func SumAll(allnums ...[]int) (sums []int) {
	size := len(allnums)
	// O make te permite criar um slice com uma capacidade inicial de len de numerosParaSomar que precisamos percorrer.
	sums = make ([]int, size)
	for i, nums := range allnums {
		sums[i] = Sum(nums)
	}
	return sums
}
*/

func SumAll(allnums ...[]int) []int {
	var sums []int
	for _, nums := range allnums {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllRest(numbersForSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersForSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			end := nums[1:]
			sums = append(sums, Sum(end))
		}
	}
	return sums
}