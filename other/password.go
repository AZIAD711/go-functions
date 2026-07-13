package other
// HASH PASSWORD 
func HiddenPassword(password string)string{
	var hidden string;
	for i:=0;i < len(password) ; i++{
		hidden+="*"
	}
	return hidden
}
// GENERATE RANDOM OTP 
var seed uint64 = 12345

func GenerateOTP() uint64 {
	seed = (1664525*seed + 1013904223) % 4294967296
	return 1000 + (seed % 9000)
}