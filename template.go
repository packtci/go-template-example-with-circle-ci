package main

import (
	"bytes"
	"fmt"
	"log"

	"text/template"
)

type Soldier struct {
	Name          string
	Rank          string
	TimeInService int
}

const templateText = `
Name is {{.Name}}
Rank is {{.Rank}}
Time in service is {{.TimeInService}}
`

func parseTemplate(soldier Soldier, tmpl string) *bytes.Buffer {
	var buff = new(bytes.Buffer)
	t := template.New("A template file")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return buff
	}

	err = t.Execute(buff, soldier)
	if err != nil {
		log.Fatal("Execute: ", err)
		return buff
	}
	return buff
}

func main() {
	newSoldier := Soldier{
		Name:          "Johnny Boy",
		Rank:          "SSG",
		TimeInService: 6,
	}
	txt := parseTemplate(newSoldier, templateText)
	fmt.Printf(txt.String())
}
