package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	fatal(ioutil.WriteFile("main.go", []byte(`package main

import (
	"log"

	"github.com/brideclick/initialize"
)

func main() {

}

func init() {
	err := initialize.Viper()
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handle(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func print(err error) {
	if err != nil {
		log.Println(err)
	}
}`), 0644))

	fatal(os.Mkdir("config", 0755))

	fatal(ioutil.WriteFile(path.Join("config", "dev.json"), []byte("{\n\n}"), 0644))
	fatal(ioutil.WriteFile(path.Join("config", "prod.json"), []byte("{\n\n}"), 0644))
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
