package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	"github.com/abo/graw/patn"
)

const (
	currentVersion = "0.0.1"
	usageinfo      = `graw - A pattern generator and searcher.

usage: graw [<options>] FIELD1 FIELD2

	-v, --verbose      be more verbose
	-V, --version      print version and quit
	-f, --format       output format
`
)

var (
	version bool
	verbose bool
	format  string
	patner  = patn.NewPatner()
	tpl     = template.New("graw")
)

func init() {
	flag.BoolVar(&version, "V", false, "Show version number and quit")
	flag.BoolVar(&version, "version", false, "Show version number and quit")
	flag.BoolVar(&verbose, "v", false, "Print the generated pattern")
	flag.BoolVar(&verbose, "verbose", false, "Print the generated patterns")
	flag.StringVar(&format, "f", "{{range .}}{{.}} {{end}}", "format the output")
	flag.StringVar(&format, "format", "{{range .}}{{.}} {{end}}", "format the output")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(2)
}

func contains(str string, substrs []string) bool {
	for _, sub := range substrs {
		if !strings.Contains(str, sub) {
			return false
		}
	}
	return true
}

func newExtractor(raw string, parts []string) (extractor patn.Extractor, ok bool) {
	if !contains(raw, parts) {
		return patn.Extractor{}, false
	}

	if verbose {
		fmt.Printf("raw line selected: %s\ngenerating pattern...\n", raw)
	}
	exprs, err := patner.Generate(raw, parts)
	if err != nil {
		exit(fmt.Sprint("cannot generate pattern.", err))
	}

	if verbose {
		fmt.Println("patterns generated: ", exprs)
	}
	extractor, err = patn.NewExtractor(exprs)
	if err != nil {
		exit(fmt.Sprint("invalid pattern generated.", err))
	}
	return extractor, true
}

func output(w io.Writer, parts []string) {
	tpl.Execute(w, parts)
	fmt.Fprintln(w)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	flag.Usage = func() {
		exit(usageinfo)
	}
	flag.Parse()

	if version {
		exit(fmt.Sprintf("graw version %s", currentVersion))
	}

	parts := flag.Args()
	if len(parts) == 0 {
		exit(usageinfo)
	}

	if _, err := tpl.Parse(format); err != nil {
		exit(fmt.Sprintf("invalid format: %s", format))
	}

	var (
		extractor patn.Extractor
		inited    bool
		pending   []string
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()

		if inited {
			output(os.Stdout, extractor.Extract(l))
			continue
		}

		pending = append(pending, l)

		if extractor, inited = newExtractor(l, parts); inited {
			for _, v := range pending {
				output(os.Stdout, extractor.Extract(v))
			}
			pending = nil
		}
	}

	if !inited {
		exit(fmt.Sprintf("the target(%v) not found in any line", parts))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
