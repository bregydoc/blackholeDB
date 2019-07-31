package blackhole

import (
	"io/ioutil"
)

func encodeDBFile(filename string, encryptKey []byte) ([]byte, error) {
	dataFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	encodedData := encrypt(dataFile, encryptKey)
	return encodedData, nil
}

func decodeDBFile(encodedDBData []byte, encryptKey []byte, whereFilename string) error {
	decodedData := decrypt(encodedDBData, encryptKey)
	err := ioutil.WriteFile(whereFilename, decodedData, 644)
	if err != nil {
		return err
	}
	return nil
}
