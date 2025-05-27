package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const artisanFile = "/tmp/git-artisan.sh"

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

	fmt.Println("Starting the artisan...")

	f, err := os.Create(artisanFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString("echo 'Hello World'")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Successfully created /tmp/git-artisan.sh")

	out, err := exec.Command("bash", artisanFile).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got following output from command execution: %s\n", out)

	defer func() {
		err := os.Remove(artisanFile)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("The artisan is done.")
}
