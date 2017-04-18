package main

import "regexp"
import "fmt"

const Xml = `<xml>
 <ToUserName>32424</ToUserName>
 <FromUserName>sfgsdfgsdg</FromUserName>
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[text]]></MsgType>
 <Content><![CDATA[this is a test]]></Content>
 <MsgId>1234567890123456</MsgId>
 </xml>`

func main() {
	reg := regexp.MustCompile(`<MsgType>(.*?)</MsgType>`)
	typ := reg.FindStringSubmatch(Xml)
	fmt.Println(typ[1])
}
