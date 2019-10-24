package main

import (
	"encoding/json"
	"fmt"
	"jsonmod/github"
	"jsonmod/movie"
	templatePackage "jsonmod/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// t1()
	// testGithub()
	// templatePackage.TestAutoEscape()
	// network()
	testadd1()
}

func t1() {
	var movies = []movie.Movie{
		{
			Title: "casa", Year: 1992, Color: false, Actors: []string{"jack", "james"},
		},
		{
			Title: "casa", Year: 1992, Color: true, Actors: []string{"jack", "james"},
		},
	}
	// data, err := json.Marshal(movies)
	// 带上缩进
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshling failed: %s", err)
	}
	fmt.Printf("%s \n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s \n", err)
	}
	fmt.Println(titles)
}

func testGithub() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// s, err := json.MarshalIndent(*result, "", "  ")
	// fmt.Printf("%s\n", s)
	// fmt.Printf("%d issues: \n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	// }
	// templatePackage.TTemp1(*result)
	templatePackage.ParseHtml(*result)
}

func network() {
	http.HandleFunc("/github-search", githubSearchHandler)
	fmt.Println("run on 1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("listen and serve:", err)
	}

}
func githubSearchHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := r.Form["q"]
	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	templatePackage.ParseHtmlToSereve(*result, w)
}

func add1(r rune) rune {
	return r + 1
}

func testadd1() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Printf("%*s </%s>\n", 2, "a", "div")
	var f = func() {
		fmt.Println("ff")
	}

	f()
}
