package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	reqURL   = "http://66.185.19.67/hrm/resource/HrmResourceSysOperation.jsp"
	loginURL = "http://66.185.19.67/login/VerifyLogin.jsp"
)

func main() {
	client := &http.Client{}
	user := url.Values{"loginfile": {"/wui/theme/ecology8/page/login.jsp?templateId=3&logintype=1&gopage="},
		"logintype": {"1"}, "fontName": {"微软雅黑"}, "message": {""}, "gopage": {""}, "formmethod": {"get"},
		"rnd": {""}, "serial": {""}, "username": {""}, "isie": {"false"}, "islanguid": {"7"},
		"loginid": {"sysadmin"}, "userpassword": {"1"}, "submit": {"登录"}}
	loginReq, _ := http.NewRequest("POST", loginURL, bytes.NewBufferString(user.Encode()))
	resp, err := client.Do(loginReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
