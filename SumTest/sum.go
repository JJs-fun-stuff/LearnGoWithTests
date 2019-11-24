package main

func Sum(numbers [5]int) int {

	sum:=0
	for _, number := range numbers{
		sum += number
	}

	// for i:=0; i<5;i++{
	// 	sum +=numbers[i]
	// }
	return sum
}