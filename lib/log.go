package lib

import "fmt"

func Audit(message string, errorMsg error)  {
	fmt.Printf("%s: %s", message, errorMsg)
}
