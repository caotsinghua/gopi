package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int32

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count = 0
		mu.Unlock()
		fmt.Fprint(w, "clear")
	})
	http.HandleFunc("/form", handlerForm)
	http.HandleFunc("/gif", handlerGif)
	fmt.Fprint(os.Stdout, "run on http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}

func handler(rw http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(rw, "url.path=%q\n", req.URL.Path)
	fmt.Println("use handler")
}
func counter(resw http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(resw, "count : %d\n", count)
	mu.Unlock()
	fmt.Println("use counter")
}

func handlerForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}

func handlerGif(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles := r.Form.Get("cycles")
	cyclesInt, err := strconv.ParseFloat(cycles, 64)
	if err != nil {
		log.Print(err)
	}
	lissajous(w, cyclesInt)
}
