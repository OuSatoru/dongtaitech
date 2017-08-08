package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var dateFlag string
	flag.StringVar(&dateFlag, "date", "yesterday", "dates")
	flag.Parse()
	if dateFlag == "yesterday" {
		now := time.Now()
		yester := now.AddDate(0, 0, -1)
		fmt.Printf("%s", yester.Format("20060102"))
	}
}
