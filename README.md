
<p align="center">
  <img src="logo.png"/>
</p>

# Black Hole DB (WIP)
BlackHoleDB (or only HoleDB) is a concept of Key-Value distributed Database.
HoleDB uses [IPFS](https://ipfs.io) as decentralized filesystem, 
and [BadgerDB](https://github.com/dgraph-io/badger) for store the local key value pairs.

**Warning:** BlackHole is work in progress, please don't use in production.
 
## How it Works
BlackHoleDB create an encrypted file into IPFS filesystem and this return an Qm name (the decentralized path), 
this Qm path is saved into BadgerDB instance as value where the key is the initial key choose. When you want get your 
value from the distributed web BlackHoleDB get the Qm linked your key (from BadgerDB) and with this Qm path HoleDB gets
the encrypted file from IPFS and finally it decrypted it.

Example code:
```go
options := blackhole.DefaultOptions
db, err := blackhole.Open(options)
if err != nil {
	panic(err)
}
	
key := "answer"

err = db.Set(key, []byte("Hello World, from BlackHoleDB"))
if err != nil {
	panic(err)
}

data, err := db.Get(key)
if err != nil {
	panic(err)
}

fmt.Println("Answer: ", string(data))
// Answer: Hello World, from BlackHoleDB

```

## About Options Configuration

You can configure the params of your blackhole instance, 
you can see the struct related above.

```go
type Options struct {
	PrivateKey         []byte // Your encoding key
	EndPointConnection string // Your IPFS Node endpoint
	PrincipalNode      string // Useless now (WIP)

	LocalDBDir      string // Your Local Badger DB
	LocalDBValueDir string // Your Local Badger DB
}
```

The default configuration is:

```go
var DefaultOptions *Options = &Options{
	LocalDBDir:         "/tmp/badger",
	LocalDBValueDir:    "/tmp/badger",
	EndPointConnection: "localhost:5001",
}
// Note: You need to define your privateKey like this:
// opts.PrivateKey, _ = hex.DecodeString("44667768254d593b7ea48c3327c18a651f6031554ca4f5e3e641f6ff1ea72e98")
```

## Basic Usage Example
```go
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
```
## TODO

- [ ] Create an ORM layer for save complex structs
- [ ] Improve the security
- [ ] Make Blackhole a metaDB. Save all param configurations on Distributed web (IPFS)
- [ ] Use [MsgPack](https://msgpack.org/index.html) to serialize the data


## Contributing to BlackHoleDB

BlackHoleDB is an open source project and contributors are welcome!

