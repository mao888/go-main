package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "/infile", "File contains values for sorting")
var outfile *string = flag.String("o", "/outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "/qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
}
