package main

import (
	"fmt"
)

func main() {
	options := DefaultOptions
	db, err := Open(options)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	key := "answer"

	err = db.Set(key, []byte("Hello World, from BlackHoleDB"))
	if err != nil {
		panic(err)
	}

	data, err := db.Get(key)
	if err != nil {
		panic(err)
	}

	qm, err := db.GetQmFromKey(key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Qm: ", qm)
	fmt.Println("Answer: ", string(data))
}
