package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(ErrJson(1, "nima"))
}
