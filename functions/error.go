package functions

import (
	"os"
)

func Error(err error, message string) {
	if err != nil {
		panic(message +" "+ err.Error())
		os.Exit(1)
	}
}