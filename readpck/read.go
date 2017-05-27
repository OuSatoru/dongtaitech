package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	xml, _ := ioutil.ReadFile("0220100.pck")
	//fmt.Println(string(xml))
	reg := regexp.MustCompile(`<Name>(.*?)</Name>\s*?<Value>(.*?)</Value>`)
	for _, item := range reg.FindAllStringSubmatch(string(xml), -1) {
		fmt.Println(item[1], item[2])
	}
}
