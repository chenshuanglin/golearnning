package main

import (
	"chap2/github/lib"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := lib.GetResult(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total number : %d\n", result.TotalCount)
	/*
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	*/

	const temp = `{{.TotalCount}} issue: 
    {{range .Items}}----------------------------
    Number: {{.Number}}
    User:   {{.User.Login}}
    Title:  {{.Title | printf "%.55s"}}
    Age:    {{.CreatedAt | daysAgo }} days
    {{end}}`

	report := template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp))
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
