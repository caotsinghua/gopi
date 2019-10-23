package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func t1() {
	ages := make(map[string]int)
	fmt.Println(ages)
	ages = map[string]int{
		"alice":  12,
		"carlie": 22,
	}
	fmt.Println(ages)
	delete(ages, "alice")
	fmt.Println(ages)
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)
	var names = make([]string, 0, len(ages))
	fmt.Println("len", len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages2 map[string]int
	// ages2["a"] = 1
	fmt.Println(ages2, ages2 == nil)
	tmp, ok := ages2["bob"]
	if !ok {
		fmt.Println("!ok", tmp) // 不存在的0值
	}
	// _ = &ages["bob"]
	// fmt.Println(_)
}
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	fmt.Println("len", len(x))
	for k, vx := range x {
		if yv, ok := y[k]; !ok || yv != vx {
			return false
		}
	}
	return true
}
func t2() {
	// x := map[string]int{}
	var x, y map[string]int
	fmt.Println(equal(x, y))
}

func depup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner((os.Stdin))
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "depup %v\n", err)
		os.Exit(1)
	}
}

func charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int //
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 字符值，长度，错误
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "chartcount %v \n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
func main() {
	// t1()
	// t2()
	// depup()
	// charcount()
	// x := map[string]int{"A": 0}
	// y := map[string]int{"B": 0}
	// fmt.Println(equal(x, y))
	// wordfreq()
	charcount2()
}

func wordfreq() {
	report := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		report[input.Text()]++
	}
	for k, v := range report {
		fmt.Printf("word:%s counts:%d \n", k, v)
	}
}

// charcount 改造
func charcount2() {
	letters := make(map[rune]int)
	nums := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int //
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 字符值，长度，错误
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "chartcount %v \n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsNumber(r) {
			nums[r]++
		}
		if unicode.IsLetter(r) {
			letters[r]++
		}
		utflen[n]++
	}
	fmt.Printf("letters\tcount\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("nums\tcount\n")
	for c, n := range nums {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
