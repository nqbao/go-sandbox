package main

import (
	"fmt"

	rocksdb "github.com/tecbot/gorocksdb"
)

func OpenAllCFs(name string) (*rocksdb.DB, []*rocksdb.ColumnFamilyHandle, error) {
	opts := rocksdb.NewDefaultOptions()
	defer opts.Destroy()
	opts.SetCreateIfMissing(true)

	// how many LOG.old file to keep
	opts.SetKeepLogFileNum(10)
	opts.SetCompression(rocksdb.ZLibCompression)

	cfs, err := rocksdb.ListColumnFamilies(opts, "./db")
	if err != nil {
		return nil, nil, err
	}

	cfOpts := make([]*rocksdb.Options, len(cfs))
	for i := range cfs {
		cfOpt := rocksdb.NewDefaultOptions()
		cfOpt.SetKeepLogFileNum(10)
		cfOpt.SetCompression(rocksdb.ZLibCompression)

		cfOpts[i] = cfOpt
	}

	// free all the options for cfs
	defer func() {
		for i := range cfOpts {
			cfOpts[i].Destroy()
		}
	}()

	fmt.Printf("column families: %v\n", cfs)

	return rocksdb.OpenDbColumnFamilies(opts, "./db", cfs, cfOpts)
}

func main() {
	db, cfs, err := OpenAllCFs("./db")

	if err != nil {
		panic(err)
	}

	defer func() {
		for _, cf := range cfs {
			cf.Destroy()
		}
		db.Close()
	}()

	// write a key
	writeOpts := rocksdb.NewDefaultWriteOptions()
	defer writeOpts.Destroy()
	err = db.Put(writeOpts, []byte("hello"), []byte("world"))

	if err != nil {
		panic(err)
	}

	// write multiple keys at a time
	batch := rocksdb.NewWriteBatch()
	batch.Put([]byte("test"), []byte("value"))
	batch.Put([]byte("test2"), []byte("value2"))
	err = db.Write(writeOpts, batch)

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

	// loop through all keys
	it := db.NewIterator(readOpts)
	defer it.Close()
	it.SeekToFirst()

	for ; it.Valid(); it.Next() {
		fmt.Printf("Key: %s\n", it.Key().Data())
	}

	// column families
	// cfOpts := rocksdb.NewDefaultOptions()
	// defer cfOpts.Destroy()
	// cf, err := db.CreateColumnFamily(cfOpts, "cf1")
	// if err != nil {
	// 	panic(err)
	// }
	// defer cf.Destroy()

	// ListColumnFamilies
}
