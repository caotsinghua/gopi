package templatePackage

import (
	"html/template"
	"io"
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
	// report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tpl)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := report.Execute(os.Stdout, data); err != nil {
	// 	log.Fatal(err)
	// }
	var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tpl))
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func ParseHtml(data github.IssuesSearchResult) {
	var issueList = template.Must(
		template.New("issueList").Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
		`))
	if err := issueList.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
func ParseHtmlToSereve(data github.IssuesSearchResult, w io.Writer) {
	var issueList = template.Must(
		template.New("issueList").Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
		`))
	if err := issueList.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
func TestAutoEscape() {
	const tpl = `<p>{{.A}}</p><p>{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(tpl))
	var data struct {
		A string
		B template.HTML
	}

	data.A = "<h1>这里isA</h1>"
	data.B = "<h1>this is B</h1>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
