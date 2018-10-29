package main

import (
	"fmt"
	"time"

	rocksdb "github.com/tecbot/gorocksdb"
)

func IteratorWithPrefixExample(db *rocksdb.DB) error {
	readOpts := rocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	it2 := db.NewIterator(readOpts)
	defer it2.Close()
	it2.Seek([]byte("2018-"))

	fmt.Printf("prefix scan\n")
	for ; it2.ValidForPrefix([]byte("2018-")); it2.Next() {
		fmt.Printf("Key: %s\n", it2.Key().Data())
	}

	return nil
}

func IteratorExample(db *rocksdb.DB) error {
	readOpts := rocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	// loop through all keys
	it := db.NewIterator(readOpts)
	defer it.Close()
	it.SeekToFirst()

	for ; it.Valid(); it.Next() {
		fmt.Printf("Key: %s\n", it.Key().Data())
	}
	return nil
}

// simple example of using tailing iterator to read newly inserted data
func WatchChangeExample(db *rocksdb.DB, stop chan bool) {
	readOpts := rocksdb.NewDefaultReadOptions()
	readOpts.SetTailing(true)
	defer readOpts.Destroy()

	it := db.NewIterator(readOpts)
	defer it.Close()

	prefix := []byte("2018-")
	it.Seek(prefix)

	shouldStop := false
	for {
		select {
		case <-stop:
			shouldStop = true

		default:
			// copy the key to new byte buffer
			curkey := make([]byte, len(it.Key().Data()))
			copy(curkey, it.Key().Data())

			fmt.Printf("got key: %s %v\n", curkey, it.Valid())
			it.Next()

			if !it.ValidForPrefix(prefix) {
				// means there is no more available key right now, seek to previous key
				// fmt.Printf("seek %s\n", curkey)
				it.Seek(curkey)
				time.Sleep(10 * time.Microsecond)
			}
		}

		if shouldStop {
			break
		}
	}

	fmt.Printf("stopping the loop")
}
