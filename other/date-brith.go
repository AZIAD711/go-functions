package other

import "time"
// CACULATE AGE FROM DATE OF BRITH 
func CalculateAge(year int) int {
	return time.Now().Year() - year
}
// GENERATE CURRENT YEAR 
func generateCurrenYear()int{
	return time.Now().Year()
}
// GENERATE CURRENT DAY 
func generateCurrenDay()int{
	return time.Now().Day()
}