package main

import (
	"flag"
	"fmt"
	"os"
)

var year int
var githubProject string

func init() {
	flag.IntVar(&year, "year", 1995, "Year of contribution to print the art.")
	flag.StringVar(&githubProject, "project", "", "Github project used for dummy commits.")

	flag.Parse()

	if githubProject == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	fmt.Printf("The pixel art will be printed in following year: %d\n", year)
	fmt.Printf("For the dummy commits following project will be used: %s\n", githubProject)

}
