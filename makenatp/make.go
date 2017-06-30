package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"os"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Result struct {
	XMLName xml.Name `xml:"Natp"`
	Version string
	Head    Head
	Data    Data
}

type Head struct {
	XMLName      xml.Name
	TemplateCode string
	TransCode    string
	ReservedCode string
}

type Data struct {
	XMLName xml.Name
	Item    []Item
}

type Item struct {
	Name  string
	Value string
	Desp  string
}

func utf8togbk(utf []byte) []byte {
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(utf), simplifiedchinese.GBK.NewDecoder()))
	return data
}

func main() {
	file, err := os.Open("NatpNames.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	natp, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	params := strings.Split(string(natp), "\n")
	var result Result
	var items []Item
	result.Version = "3"
	result.Head.TemplateCode = params[0]
	result.Head.TransCode = params[1]
	for _, i := range params[2:] {
		var item Item
		item.Name = i
		items = append(items, item)
	}
	result.Data.Item = items

	xm, err := xml.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	xms := `<?xml version="1.0" encoding="GBK"?>` + "\n\n" + string(xm)
	xmgb := utf8togbk([]byte(xms))
	err = ioutil.WriteFile("gen.pck", xmgb, 0644)
	if err != nil {
		panic(err)
	}
}
