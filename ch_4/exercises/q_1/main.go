package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(notify(1))
}

func notify(sec int) string {
	var result string
	select {
	case <-time.After(time.Duration(sec) * time.Second):
		result = "Passed " + strconv.Itoa(sec) + " sec"
	}

	return result
}
