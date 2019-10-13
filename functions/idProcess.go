package functions

import (
	// "fmt"
	"strconv"
)

func IdProcess(n int, mode string) []string  {
	// const puerto string
	var id string
	var ids []string

	if mode == "local" {
		for i := 0; i < n; i++ {
			// id = "127.0.0.1:140" + strconv.Itoa(i)
			id = strconv.Itoa(i)+ ":" + "127.0.0.1:140" + strconv.Itoa(i)
			ids = append(ids, id)
			// ids = ids + "," + id
			// fmt.Println(ids)
		}

	} else if mode == "remote" {
		for i := 191; (i < n+210) && (n+191 > i); i++{
			id = strconv.Itoa(i)+ ":" + "144.210.154." + strconv.Itoa(i) + ":1400"
			ids = append(ids, id)
			// ids = ids + "," + id
			// fmt.Println(ids)
		}

	}

	return ids
}
