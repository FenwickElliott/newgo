package main

import (
	"errors"
	"log"
	"os"
	"runtime"
)

var body = `package main

import (
	"log"
	"runtime"
)

func main() {

}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func fatal(err error) {
	if err != nil {
		_, filename, linenumber, _ := runtime.Caller(1)
		log.Printf("fatal error thrown by '%s:%d', error: '%s'", filename, linenumber, err)
	}
}
`

func main() {
	_, err := os.Stat("main.go")
	switch {
	case err == nil:
		log.Fatal("main.go already exists")
	case errors.Is(err, os.ErrNotExist):
		break // desired case
	case err != nil:
		log.Fatalf("unexpected error: '%s'", err)
	default:
		log.Fatal("unhandled case! This should never happen")
	}

	f, err := os.OpenFile("main.go", os.O_CREATE|os.O_WRONLY, 0644)
	fatal(err)
	defer f.Close()

	_, err = f.WriteString(body)
	fatal(err)

	log.Println("sucsesfully wrote stub main file")
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func fatal(err error) {
	if err != nil {
		_, filename, linenumber, _ := runtime.Caller(1)
		log.Printf("fatal error thrown by '%s:%d', error: '%s'", filename, linenumber, err)
	}
}
