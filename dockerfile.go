// templatetest
package main

import "fmt"

import "html/template"

//import "text/template"
import "os"

type dockerfile struct {
	Base       string
	CopyFrom   string
	CopyTo     string
	Entrypoint string
}

func main() {

	mysql := dockerfile{Base: "mysql:5.6.34",
		CopyFrom:   "entrypoint.sh r.sql my.cnf start.sh",
		CopyTo:     "/tmp/",
		Entrypoint: "/usr/bin/hostname"}

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
