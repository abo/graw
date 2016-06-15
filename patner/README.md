Tutorial
--------
```
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abo/graw/patner"
	"github.com/abo/patnsvc"
)

func main() {
	mailsv := `Thu May 15 2015 00:15:05 mailsv1 sshd[2716]: Failed password for invalid user postgres from 86.212.199.60 port 4093 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[2596]: Failed password for invalid user whois from 86.212.199.60 port 3311 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[24947]: pam_unix(sshd:session): session opened for user djohnson by (uid=0)
Thu May 15 2015 00:15:05 mailsv1 sshd[3006]: Failed password for invalid user info from 86.212.199.60 port 4078 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[5298]: Failed password for invalid user postgres from 86.212.199.60 port 1265 ssh2`
	lines := strings.Split(mailsv, "\n")

	patner := patner.NewPatner()
	ps, err := patner.Generate([]patnsvc.Line{
		patnsvc.Line{
			Raw:      "Thu May 15 2015 00:15:05 mailsv1 sshd[2716]: Failed password for invalid user postgres from 86.212.199.60 port 4093 ssh2",
			Expected: "2716",
		},
	})

	if err != nil {
		panic(err)
	}

	for _, p := range ps {
		re := regexp.MustCompile(p.Expr)
		pids := make([]string, len(lines))
		for i, l := range lines {
			matches := re.FindStringSubmatch(l)
			pids[i] = matches[1]
		}
		fmt.Println(pids)
		// Output: [2716 2596 24947 3006 5298]
	}
}
```