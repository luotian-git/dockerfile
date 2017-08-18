// templatetest
package main

import "fmt"

import "html/template"

//import "text/template"
import "os"

type dockerfile struct {
	Base     string
	CopyFrom string
	CopyTo   string
	Command  string
}

func main() {

	mysql := dockerfile{Base: "luotian1/learn-ping",
		CopyFrom: "*",
		CopyTo:   "/tmp/",
		Command:  "ping 10.110.18.70"}

	//fmt.Println("dockerfile = ", mysql)
	var tmpl *template.Template
	//tmpl := template.New("dockerfile")
	//tmpl.Parse("FROM {{.Base}}\nCOPY {{.CopyFrom}} {{.CopyTo}}\nENTRYPOINT {{.Entrypoint}} \n")
	tmpl, err := template.ParseFiles("dockerfile-template.txt")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, mysql)
	fmt.Println("")
}
