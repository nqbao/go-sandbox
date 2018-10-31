package main

import rocksdb "github.com/tecbot/gorocksdb"

func TransactionExample() error {
	opts := rocksdb.NewDefaultOptions()
	opts.SetCompression(rocksdb.ZLibCompression)
	defer opts.Destroy()

	txdbOpts := rocksdb.NewDefaultTransactionDBOptions()
	defer txdbOpts.Destroy()

	_, err := rocksdb.OpenTransactionDb(opts, txdbOpts, "./db")

	if err != nil {
		return err
	}

	// txdb.Name

	// defer txdb.Close()

	return nil
}
