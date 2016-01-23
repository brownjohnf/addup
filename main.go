// Addup adds integers in columns from stdin
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var k = flag.Int("k", 1, "field to sum")
var v = flag.Bool("v", false, "verbose errors")

func main() {
	flag.Parse()

	var sum int
	var line int

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line++

		f := strings.Fields(input.Text())

		if len(f) < *k {
			if *v {
				fmt.Fprintf(os.Stderr, "%d: Insufficient fields: %v; skipping\n", line, f)
			}
			continue
		}

		i, err := strconv.Atoi(f[*k-1])
		if err != nil {
			if *v {
				fmt.Fprintf(os.Stderr, "%d: Cannot parse field: %v; skipping\n", line, f[*k-1])
			}
			// Next line if we can't parse this one
			continue
		}

		sum += i
	}

	// Output the sum
	fmt.Println(sum)
}
