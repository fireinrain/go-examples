package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "File contarins values for sorting")
var outfile *string = flag.String("o", "outfile", "File to recieve sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	fmt.Println(*infile)
}
