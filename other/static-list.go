package other

// GENDER LIST
func CheckGender(value string) bool {
	if value == "male" || value == "female" {
		return true
	}
	return false
}