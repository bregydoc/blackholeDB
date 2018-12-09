package main

import (
	"bytes"
	"io/ioutil"
)

func (db *BlackHoleDB) writeKeyValuePair(key, value string) error {
	txn := db.localDB.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *BlackHoleDB) readKeyValuePair(key string) (string, error) {
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

func (db *BlackHoleDB) Set(key string, value []byte) error {

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

func (db *BlackHoleDB) Get(key string) ([]byte, error) {
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

func (db *BlackHoleDB) GetQmFromKey(key string) (string, error) {
	hash, err := db.readKeyValuePair(key)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (db *BlackHoleDB) Update(key string, value []byte) error {

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
