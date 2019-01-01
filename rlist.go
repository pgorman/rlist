// Rlist outputs a ranged numeric list as HTML.
// Paul Gorman, 2019
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Format defines the output format (e.g., "HTML" or "LaTeX").
var Format = "HTML"

func init() {
	c := flag.Bool("c", false, "Don't format list items, just fence them off as '```' code blocks.")
	x := flag.Bool("x", false, "Output LaTeX instead of HTML.")
	flag.Parse()
	if *c {
		Format = "Fenced Code"
	}
	if *x {
		Format = "LaTeX"
	}
}

func printList(items []string, ranged bool) {
	if len(items) < 1 {
		return
	}
	if ranged == true {
		switch Format {
		case "Fenced Code":
			fmt.Println("```")
			for _, it := range items {
				fmt.Println(it)
			}
			fmt.Println("```")
		case "LaTeX":
			fmt.Println(`\begin{description}`)
			for _, it := range items {
				si := strings.SplitAfterN(it, ".", 2)
				dt := si[0]
				dd := si[1]
				fmt.Printf("\\item[%s] %s\n", dt, dd)
			}
			fmt.Println(`\end{description}`)
		default:
			fmt.Println(`<dl class="rlist">`)
			for _, it := range items {
				si := strings.SplitAfterN(it, ".", 2)
				dt := si[0]
				dd := si[1]
				fmt.Printf("<dt>%s</dt>\n<dd>%s</dd>\n", dt, dd)
			}
			fmt.Println(`</dl>`)
		}
	} else {
		for _, it := range items {
			fmt.Println(it)
		}
	}
}

func main() {
	rn := regexp.MustCompile(`^\s*[0-9]+\.\s+`)
	rr := regexp.MustCompile(`^\s*[0-9]+(?:\s*[\-–]{1,2}\s*[0-9]+)\.\s+`)
	list := false
	ranged := false
	items := make([]string, 0, 100)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()

		if !ranged && rr.MatchString(t) {
			ranged = true
		}

		if rn.MatchString(t) || rr.MatchString(t) {
			list = true
			items = append(items, t)
		} else {
			list = false
		}

		if !list {
			printList(items, ranged)
			items = items[0:0]
			fmt.Println(t)
		}
	}
	// In case the last line of input is a list item, flush:
	printList(items, ranged)

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
}
