package math
// SUM NUMBERS FUNCTIONS
func SumNumbers(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// MINUS NUMBERS FUNCTION
func MinusNumbers(numbers ...float64) float64 {
	var min  float64
	for _, number := range numbers {
		min -= number
	}
	return min
}

// MULTIPLIY NUMBERS FUNCTION
func MultiplyNumbers(numbers ...float64) float64{
	var product float64;
	for _,number :=range numbers{
		product*=number
	}
	return product
}
// POWER NUMBERS FUNCTION 
func PowerNumber(number  float64) float64{
	return  number * number
}
