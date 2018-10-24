package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("./db", 0666, nil)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("test"))

		if err != nil {
			return err
		}

		return b.Put([]byte("hello"), []byte("world"))
	})

	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("test"))

		next, _ := b.NextSequence()

		return b.Put(
			[]byte(time.Now().Format(time.RFC3339)),
			[]byte(strconv.Itoa(int(next))),
		)
	})

	if err != nil {
		panic(err)
	}

	// try reading
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("test")).Cursor()

		// range read
		for k, v := c.Seek([]byte("2018-10")); k != nil && bytes.HasPrefix(k, []byte("2018-10")); k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
