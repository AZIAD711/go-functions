package other

import "time"
// CACULATE AGE FROM DATE OF BRITH 
func CalculateAge(year int) int {
	return time.Now().Year() - year
}
// GENERATE CURRENT DATE 
func generateCurrenYear()int{
	return time.Now().Year()
}