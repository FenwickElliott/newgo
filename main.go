package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var name = flag.String("name", "", "Name of project to be created.")

func main() {
	flag.Parse()
	err := os.Mkdir(*name, 0755)
	check(err)
	err = ioutil.WriteFile(path.Join(".", *name, "main.go"), []byte("package main\n\nfunc main() {\n\n}\n"), 0644)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
