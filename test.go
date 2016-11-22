package main

import (
	"fmt"

	"github.com/kardianos/osext"
)

func main() {
	filename, _ := osext.ExecutableFolder()
	fmt.Println(filename)
}
