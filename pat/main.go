package main

import (
	"flag"
	"fmt"
	"github.com/julz/pat"
)

func main() {
	server := flag.Bool("server", false, "true to run the HTTP server interface")
	pushes := flag.Int("pushes", 1, "number of pushes to attempt")
	concurrency := flag.Int("concurrency", 1, "max number of pushes to attempt in parallel")
	output := flag.Bool("output", false, "true to run the commands and print output the terminal")
	flag.Parse()

	if *server == true {
		fmt.Println("Starting in server mode")
		pat.Serve()
		pat.Bind()
	} else {
		resp := pat.RunCommandLine(*pushes, *concurrency, *output)
		fmt.Printf("\n\nTest Complete. Total Time: %d", resp.TotalTime)
	}
}
