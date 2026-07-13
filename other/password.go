package other
// HASH PASSWORD 
func HiddenPassword(password string)string{
	var hidden string;
	for i:=0;i < len(password) ; i++{
		hidden+="*"
	}
	return hidden
}