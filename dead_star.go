package blackhole

import (
	"errors"

	"github.com/dgraph-io/badger"
	shell "github.com/ipfs/go-ipfs-api"
)

const (
	// RequiredKeyLength determinate the exact length of the encrypt key
	RequiredKeyLength = 32
)

// DB represents a blackhole db instance,
// this struct is the real db when it's sync with the ipfs files
type DB struct {
	encryptKey    []byte
	principalNode string

	localDB     *badger.DB
	remoteShell *shell.Shell
	options     *Options
}

// Options is the options configuration of blackholole DB
// TODO: Define better options and a new paradigm to set it
type Options struct {
	PrivateKey         []byte
	EndPointConnection string
	PrincipalNode      string
	LocalDBDir         string
}

// DefaultOptions is used with any options passed,
// this config saves your db file into your temporal computer files (UNIX)
// TODO: Improve for another SO
var DefaultOptions = &Options{
	LocalDBDir:         "/tmp/badger",
	EndPointConnection: "localhost:5001",
}

// ValidateKey takes a byte slice and checks that minimum requirements are met for
// the key. It returns an error if the requirements are not met.
func ValidateKey(k []byte) error {
	if len(k) == 0 {
		return errors.New("no PrivateKey set")
	}
	if len(k) != RequiredKeyLength {
		return errors.New("invalid PrivateKey length. Key must be 32 bytes")
	}
	return nil
}

// Open opens a new instance of blackholedb
func Open(options *Options) (*DB, error) {
	if err := ValidateKey(options.PrivateKey); err != nil {
		return nil, err
	}

	db := new(DB)
	db.encryptKey = options.PrivateKey
	db.principalNode = options.PrincipalNode
	db.options = options

	opts := badger.DefaultOptions(options.LocalDBDir)

	ldb, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	db.localDB = ldb
	sh := shell.NewShell(options.EndPointConnection)
	db.remoteShell = sh

	return db, nil
}

// Close ...
// TODO
func (db *DB) Close() {
	db.localDB.Close()
}
