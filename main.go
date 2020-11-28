package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kou-pg-0131/docker-tags/src/interfaces/controllers"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: docker-tags <IMAGE>")
	}
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	c := controllers.NewDockerTagsControllerFactory().Create()
	if err := c.ShowAll(args[0]); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
