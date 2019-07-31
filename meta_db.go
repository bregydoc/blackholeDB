package blackhole

import (
	"bytes"
	"io/ioutil"

	"github.com/dgraph-io/badger"
	shell "github.com/ipfs/go-ipfs-api"
)

// MetaDB is a db by reference, a way you can to create a very distrubuted DB
type MetaDB struct {
	encryptKey    []byte
	principalNode string

	publicKey  string
	privateKey []byte

	localDB     *badger.DB
	remoteShell *shell.Shell
}

func (db *DB) snapshotMetaDB(publicKey string, privateKey []byte) (*MetaDB, error) {
	data, err := encodeDBFile(db.options.LocalDBDir, privateKey)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	cid, err := db.remoteShell.Add(buf)
	if err != nil {
		return nil, err
	}

	return &MetaDB{
		encryptKey:    db.encryptKey,
		principalNode: db.principalNode,
		publicKey:     cid,
		privateKey:    privateKey,
		localDB:       db.localDB,
		remoteShell:   db.remoteShell,
	}, nil
}

func (db *DB) readMetaDB(publicKey string, privateKey []byte) (*MetaDB, error) {
	reader, err := db.remoteShell.Cat(publicKey)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = decodeDBFile(data, privateKey, db.options.LocalDBDir)
	if err != nil {
		return nil, err
	}

	ldb, err := badger.Open(badger.DefaultOptions(db.options.LocalDBDir))
	if err != nil {
		return nil, err
	}

	return &MetaDB{
		encryptKey:    db.encryptKey,
		principalNode: db.principalNode,
		publicKey:     publicKey,
		privateKey:    privateKey,
		localDB:       ldb,
		remoteShell:   db.remoteShell,
	}, nil

}
