package blackhole

import (
	"bytes"
	"io/ioutil"
)

func (db *DB) writeKeyValuePair(key, value string) error {
	txn := db.localDB.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (db *DB) readKeyValuePair(key string) (string, error) {
	txn := db.localDB.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return "", err
	}

	var returnedValue []byte
	err = item.Value(func(val []byte) error {
		returnedValue = val
		return nil
	})
	if err != nil {
		return "", err
	}

	return string(returnedValue), nil

}

// Set acts in two stages,
// first encode your <value> and put into distributed IPFS.
// second, save (locally) a relation between your <key> and
// the cid generated from your save performed by ipfs
func (db *DB) Set(key string, value []byte) error {
	data := encrypt(value, db.encryptKey)

	buf := bytes.NewBuffer(data)

	cid, err := db.remoteShell.Add(buf)
	if err != nil {
		return err
	}

	err = db.writeKeyValuePair(key, cid)
	if err != nil {
		return err
	}

	return nil
}

// Get performs a get action of your Blachole db,
// it works in two stages (like set function), firs read
// the relation between your key and a one cid, second, collect
// this cid from ipfs and finally tries to decrypt and return
// the value of your key
func (db *DB) Get(key string) ([]byte, error) {
	hash, err := db.readKeyValuePair(key)
	if err != nil {
		return nil, err
	}

	reader, err := db.remoteShell.Cat(hash)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	dData := decrypt(data, db.encryptKey)
	return dData, nil
}

// GetQmFromKey  returns your cid related to your key,
// if exist, of course
func (db *DB) GetQmFromKey(key string) (string, error) {
	hash, err := db.readKeyValuePair(key)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Update takes your value, encode and save the relation,
// it is exactly equal to set action
func (db *DB) Update(key string, value []byte) error {
	data := encrypt(value, db.encryptKey)

	buf := bytes.NewBuffer(data)

	cid, err := db.remoteShell.Add(buf)
	if err != nil {
		return err
	}

	err = db.writeKeyValuePair(key, cid)
	if err != nil {
		return err
	}

	return nil
}
