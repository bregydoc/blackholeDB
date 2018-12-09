package main

import (
	"fmt"
	"github.com/bregydoc/blackholeDB"
)

func main() {
	db, err := blackhole.Open(blackhole.DefaultOptions)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Set("answer", []byte("42"))
	if err != nil {
		panic(err)
	}

	answer, err := db.Get("answer")
	if err != nil {
		panic(err)
	}

	fmt.Println("The answer of the life is: ", string(answer))
	// The answer of the life is:  42
}
