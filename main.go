package main

import (
	"github.com/docopt/docopt-go"
	"github.com/sergsr/goleet/server"
)

func main() {
	usage := `
Usage: 
  goleet solve <id>
  goleet server

Examples:

Options:
  -h --help     Show this screen.
`
	opts, _ := docopt.ParseDoc(usage)
	if serve, _ := opts.Bool("server"); serve {
		server.Run()
	}
}
