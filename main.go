package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
Usage: 

Examples:

Options:
`
	args, _ := docopt.ParseDoc(usage)
	fmt.Println(args)
}
