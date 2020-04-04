package util

import "fmt"

func Audit(message string)  {
	fmt.Println(message)
}

func Error(errorMsg error, message string)  {
	fmt.Println(message, errorMsg)
}