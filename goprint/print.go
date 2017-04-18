package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println("goroutine " + strconv.Itoa(i))
		}(i)
	}
	time.Sleep(2 * time.Second)
}
