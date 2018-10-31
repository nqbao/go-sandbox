package main

import (
	"fmt"
	"strconv"
	"time"

	rocksdb "github.com/tecbot/gorocksdb"
)

func main() {
	stop := make(chan bool)

	db, cfs, err := OpenDbWithAllCFs("./db")

	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Printf("shutting down\n")

		time.Sleep(1 * time.Second)
		close(stop)

		for _, cf := range cfs {
			cf.Destroy()
		}
		db.Close()
	}()

	// go WatchChangeExample(db, stop)

	// write a key
	writeOpts := rocksdb.NewDefaultWriteOptions()
	writeOpts.SetSync(true)
	defer writeOpts.Destroy()
	err = db.Put(writeOpts, []byte("hello"), []byte("world"))

	if err != nil {
		panic(err)
	}

	// write multiple keys at a time
	batch := rocksdb.NewWriteBatch()
	// batch.Delete([]byte("hello"))
	batch.Put([]byte("test"), []byte("value"))
	batch.Put([]byte("test2"), []byte("value2"))

	newkey := time.Now().Format(time.RFC3339)

	batch.Put(
		[]byte(newkey),
		[]byte(strconv.Itoa(int(time.Now().Unix()))),
	)

	err = db.Write(writeOpts, batch)

	if err != nil {
		panic(err)
	}

	fmt.Printf("write done: %v\n", newkey)

	// read a key
	readOpts := rocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	result, err := db.Get(readOpts, []byte("hello"))
	defer result.Free()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result.Data())

	IteratorWithPrefixExample(db)

	// err = TransactionExample()
	// if err != nil {
	// 	panic(err)
	// }

	/*err = SnapshotExample(db)
	if err != nil {
		panic(err)
	}*/

	/*err = MergeExample(db)
	if err != nil {
		panic(err)
	}*/
}
