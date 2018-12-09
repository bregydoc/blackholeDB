package main

import (
	"github.com/dgraph-io/badger"
	"github.com/ipfs/go-ipfs-api"
)

type BlackHoleDB struct {
	encryptKey    string
	principalNode string

	localDB     *badger.DB
	remoteShell *shell.Shell
}

type BHOptions struct {
	PrivateKey         string
	EndPointConnection string
	PrincipalNode      string

	LocalDBDir      string
	LocalDBValueDir string
}

var DefaultOptions *BHOptions = &BHOptions{
	LocalDBDir:         "/tmp/badger",
	LocalDBValueDir:    "/tmp/badger",
	EndPointConnection: "localhost:5001",
	PrivateKey:         "black_hole",
}

func Open(options *BHOptions) (*BlackHoleDB, error) {
	db := new(BlackHoleDB)
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

func (db *BlackHoleDB) Close() {
	db.localDB.Close()
}
