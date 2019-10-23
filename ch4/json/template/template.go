package templatePackage

import (
	"html/template"
	"jsonmod/github"
	"log"
	"os"
	"time"
)

func TTemp1(data github.IssuesSearchResult) {
	const tpl = `
	{{.TotalCount}} issues:
	{{range .Items}}-------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title|printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}
	`
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
