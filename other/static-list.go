package other

// GENDER LIST
func CheckGender(value string) bool {
	if value == "male" || value == "female" {
		return true
	}
	return false
}
// CHECK PHONE NUMBER LENGTH FUNCTION 
func CheckPhoneNumberLength(phoneNumber string)bool{
	if len(phoneNumber)>11{
		return  false
	}
	return true
}
// CHECK NATIONAL ID 
func CheckNationalID(nationalId string)bool{
	if len(nationalId)==14{
		return true 
	}
	return false
}