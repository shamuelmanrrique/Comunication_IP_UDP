package functions

import (
	"fmt"
	"os"
)

func Error(err error, message string) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		os.Exit(1)
	}
}