package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/abo/graw/patner"
	patn "github.com/abo/patnsvc"
)

const (
	currentVersion = "0.0.1"
	usageinfo      = `graw - A pattern generator and searcher.

usage: graw [<options>] [<file>]	

	-v, --verbose		be more verbose
	-V, --version		print version and quit
	-e, --expected <fields> expected fields of first line, FIELD1 FIELD2...
	-m, --maxsamples <n>	Only sample up to <n> lines
	-f, --format		output format
`
)

var (
	version    bool
	verbose    bool
	expected   = flag.String("e", "", "expected fields of first line, FIELD1 FIELD2...")
	maxsamples uint
	format     string
)

func init() {
	flag.BoolVar(&version, "V", false, "Show version number and quit")
	flag.BoolVar(&version, "version", false, "Show version number and quit")
	flag.BoolVar(&verbose, "v", false, "Print the generated pattern")
	flag.BoolVar(&verbose, "verbose", false, "Print the generated pattern")
	flag.UintVar(&maxsamples, "m", 1, "sample up to n lines")
	flag.UintVar(&maxsamples, "maxsamples", 1, "sample up to n lines")
	flag.StringVar(&format, "f", "", "format the output")
	flag.StringVar(&format, "format", "", "format the output")
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

func sample(raw []string, exp string) []patn.Line {
	fs := strings.Split(exp, " ")
	var ret []patn.Line
	for _, l := range raw {
		if contains(l, fs) {
			for _, f := range fs {
				ret = append(ret, patn.Line{
					Raw:      l,
					Expected: f,
				})
			}
		}
	}
	return ret
}

func compile(patns []patn.Pattern) []*regexp.Regexp {
	var res []*regexp.Regexp
	for _, p := range patns {
		if re, err := regexp.Compile(p.Expr); err == nil {
			res = append(res, re)
		}
	}
	return res
}

func extract(regexps []*regexp.Regexp, raw string) []string {
	ret := make([]string, len(regexps))
	for i, re := range regexps {
		matches := re.FindStringSubmatch(raw)
		if len(matches) >= 2 {
			ret[i] = matches[1]
		}
	}
	return ret
}

//TODO subcommand - regexp, extract
func main() {
	// 1. sample data for pattern generation (need top n?)
	// 2. generate pattern & compile?
	// 3. extract fields
	// 4. format & output

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	flag.Usage = func() {
		exit(usageinfo)
	}
	flag.Parse()

	if version {
		exit(fmt.Sprintf("graw version %s", currentVersion))
	}

	if *expected == "" {
		exit(usageinfo)
	}

	var top []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := uint(0); i < maxsamples && scanner.Scan(); i++ {
		l := scanner.Text()
		top = append(top, l)
	}

	spls := sample(top, *expected)

	if len(spls) == 0 {
		exit(fmt.Sprintf("the expected(%s) not found in top %d line(s)", *expected, maxsamples))
	}

	patner := patner.NewPatner()
	patns, err := patner.Generate(spls)
	if verbose {
		fmt.Println(patns)
	}
	if err != nil {
		exit(err.Error())
	}

	regexps := compile(patns)

	for _, l := range top {
		matches := extract(regexps, l)
		fmt.Println(strings.Join(matches, "\t"))
	}

	for scanner.Scan() {
		matches := extract(regexps, scanner.Text())
		fmt.Println(strings.Join(matches, "\t"))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
