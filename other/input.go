package other
import "fmt"
// INPUT FUNCTION WITH SINGLE LINE
func Input[T any](variable *T){
	fmt.Scan(variable)
}
// INPUT FUNCTION WITH NEW LINE 
func InputWithNewLine [T any] (variable * T){
fmt.Scanln(variable)
}
