package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	mainTemplate = `package main

import (
	"log"
)

func main() {

}

func init() {

}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}`
)

func main() {
	_, err := os.Stat("main.go")
	if err != nil && os.IsNotExist(err) {
		err = ioutil.WriteFile("main.go", []byte(mainTemplate), 0644)
		fatal(err)
	}

	_, err = os.Stat("config")
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir("config", 0744)
		fatal(err)
	}

	_, err = os.Stat(path.Join("config", "dev.json"))
	if err != nil {
		if os.IsNotExist(err) {
			err = ioutil.WriteFile(path.Join("config", "dev.json"), []byte("{\n\n}"), 0644)
			fatal(err)
		} else {
			log.Println(err)
		}
	}

	_, err = os.Stat(path.Join("config", "prod.json"))
	if err != nil && os.IsNotExist(err) {
		err = ioutil.WriteFile(path.Join("config", "prod.json"), []byte("{\n\n}"), 0644)
		fatal(err)
	}

}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
