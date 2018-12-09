
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
	PrivateKey         string // Your encoding key
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
	PrivateKey:         "black_hole",
}
```