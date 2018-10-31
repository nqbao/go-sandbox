package main

import (
	"errors"
	"reflect"

	rocksdb "github.com/tecbot/gorocksdb"
)

func SnapshotExample(db *rocksdb.DB) error {
	testKey := []byte("this is new")

	writeOpts := rocksdb.NewDefaultWriteOptions()
	defer writeOpts.Destroy()

	// make sure key is not there before the snapshot
	db.Delete(writeOpts, testKey)

	snapshot := db.NewSnapshot()
	defer db.ReleaseSnapshot(snapshot)

	// add a new key to the db
	err := db.Put(writeOpts, testKey, []byte("123"))
	if err != nil {
		return err
	}

	// verify if the key is not in the snapshot
	readOpt := rocksdb.NewDefaultReadOptions()
	defer readOpt.Destroy()
	readOpt.SetSnapshot(snapshot)
	it := db.NewIterator(readOpt)
	it.SeekToFirst()

	foundKey := false
	for ; it.Valid(); it.Next() {
		if reflect.DeepEqual(it.Key().Data(), testKey) {
			foundKey = true
		}
	}

	if foundKey {
		return errors.New("key should not be found as it is in snapshot")
	}

	return nil
}
