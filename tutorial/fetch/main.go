package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// for _, url := range os.Args[1:] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch %v \n", err)
	// 		os.Exit(1)
	// 	}
	// 	b, err := ioutil.ReadAll(resp.Body)
	// 	resp.Body.Close()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch reading %s : %v", url, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("%s \n", b)
	// }
	// test1_7()
	fetchAll()
}

func test1_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get %v \n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err:%v\n", err)
			os.Exit(1)
		}
	}
}

func test1_8() {
	for _, url := range os.Args[1:] {
		fmt.Println(strings.HasPrefix(url, "http://"), url)
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get %v \n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err:%v\n", err)
			os.Exit(1)
		}
	}
}
