package blackhole

import (
	"errors"

	"github.com/dgraph-io/badger"
	"github.com/ipfs/go-ipfs-api"
)

const (
	RequiredKeyLength = 32
)

type DB struct {
	encryptKey    []byte
	principalNode string

	localDB     *badger.DB
	remoteShell *shell.Shell
}

type Options struct {
	PrivateKey         []byte
	EndPointConnection string
	PrincipalNode      string

	LocalDBDir      string
	LocalDBValueDir string
}

var DefaultOptions *Options = &Options{
	LocalDBDir:         "/tmp/badger",
	LocalDBValueDir:    "/tmp/badger",
	EndPointConnection: "localhost:5001",
}

// ValidateKey takes a byte slice and checks that minimum requirements are met for
// the key. It returns an error if the requirements are not met.
func ValidateKey(k []byte) error {
	if len(k) == 0 {
		return errors.New("No PrivateKey set.")
	}
	if len(k) != RequiredKeyLength {
		return errors.New("Invalid PrivateKey length. Key must be 32 bytes.")
	}
	return nil
}

func Open(options *Options) (*DB, error) {
	if err := ValidateKey(options.PrivateKey); err != nil {
		return nil, err
	}

	db := new(DB)
	db.encryptKey = options.PrivateKey
	db.principalNode = options.PrincipalNode

	opts := badger.DefaultOptions
	opts.Dir = options.LocalDBDir
	opts.ValueDir = options.LocalDBValueDir

	ldb, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	db.localDB = ldb
	sh := shell.NewShell(options.EndPointConnection)
	db.remoteShell = sh

	return db, nil

}

func (db *DB) Close() {
	db.localDB.Close()
}
