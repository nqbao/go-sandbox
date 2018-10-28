package main

import (
	"fmt"

	rocksdb "github.com/tecbot/gorocksdb"
)

func main() {
	opts := rocksdb.NewDefaultOptions()
	defer opts.Destroy()
	opts.SetCreateIfMissing(true)

	// how many LOG.old file to keep
	opts.SetKeepLogFileNum(10)

	db, err := rocksdb.OpenDb(opts, "./db")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	// write a key
	writeOpts := rocksdb.NewDefaultWriteOptions()
	defer writeOpts.Destroy()
	err = db.Put(writeOpts, []byte("hello"), []byte("world"))

	if err != nil {
		panic(err)
	}

	// read a key
	readOpts := rocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	result, err := db.Get(readOpts, []byte("hello"))
	defer result.Free()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result.Data())
}
