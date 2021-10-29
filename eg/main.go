package main

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
