package arraysandslices
/* 
	Diffrence between slice and array
	mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}
*/ 

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sum_slices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}


// Variadic function
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum_slices(numbers)
	}

	return sums
}

// Using append
func SumAll2(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum_slices(numbers))
	}

	return sums
}


func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum_slices(tail))
		}
	}
	return sums
}