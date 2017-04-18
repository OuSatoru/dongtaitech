package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {
	// 	sqlstr := `CREATE TABLE JSRUN.jiguan_pingfen
	// (
	//     pf_year VARCHAR(4),
	//     operator VARCHAR(8),
	//     department VARCHAR(30),
	//     jobtitle VARCHAR(15),
	//     job_vary VARCHAR(1),
	//     {{range .}}{{.}} DECIMAL(10,2),
	//     {{end}}
	// );`
	sqlstr := `SELECT {{.}}
FROM JSRUN.JIGUAN_PINGFEN
WHERE PF_YEAR = '2017' AND JOB_VARY = '0'`
	b, err := ioutil.ReadFile("jgpfsql.txt")
	if err != nil {
		panic(err)
	}
	namet := strings.Split(string(b), "\n")
	names := []string{}
	for _, v := range namet {
		names = append(names, fmt.Sprintf("avg(%s) %s", v, v))
	}
	fmt.Println(names)
	nameavg := strings.Join(names, ",\n")
	tmpl, err := template.New("test").Parse(sqlstr)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nameavg)
	if err != nil {
		panic(err)
	}
}
