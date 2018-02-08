package main

import (
	"flag"
	"fmt"
	"os"
)

var name = flag.String("name", "", "Name of project to be created.")

func main() {
	flag.Parse()
	err := os.Mkdir(*name, 0755)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
