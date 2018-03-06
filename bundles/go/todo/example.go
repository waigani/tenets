package main

import "log"

// actual TODO: wait until comment parser is fixed: https://github.com/codelingo/lexicon/issues/205

// main is the entry point to this example
func main() {
	err := serve()
	if err != nil {
		log.Fatal(err.Error()) // TODO: check error type before exiting
	}
}

// What about this comment?

// serve runs the server, returning an error if it crashes
func serve() error {
	// TODO: implement server
	return nil
}
