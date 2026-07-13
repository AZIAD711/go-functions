package other
var seed uint64 = 12345
// HASH PASSWORD 
func HiddenPassword(password string)string{
	var hidden string;
	for i:=0;i < len(password) ; i++{
		hidden+="*"
	}
	return hidden
}
// GENERATE RANDOM OTP 
func GenerateOTP() uint64 {
	seed = (1664525*seed + 1013904223) % 4294967296
	return 1000 + (seed % 9000)
}
// GENERATE RANDOM PASSWORD 
func GeneratePasswordAsNumbers(length uint64) uint64 {
	seed = (1664525*seed + 1013904223) % 4294967296
	return (length*1000) + (seed % (length*9000))
}