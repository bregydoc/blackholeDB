package main

import (
	"encoding/hex"
	"fmt"
	"log"

	blackhole "github.com/bregydoc/blackholeDB"
)

func main() {

	// initialize db options
	opts := blackhole.DefaultOptions

	// Set PrivateKey. This should come from an ENV or a secret store in the real world
	opts.PrivateKey, _ = hex.DecodeString("44667768254d593b7ea48c3327c18a651f6031554ca4f5e3e641f6ff1ea72e98")

	db, err := blackhole.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Set("answer", []byte("42"))
	if err != nil {
		log.Fatal(err)
	}

	answer, err := db.Get("answer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The answer of life is: ", string(answer))
	// The answer of life is:  42
}
