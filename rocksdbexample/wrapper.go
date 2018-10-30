package main

import (
	"fmt"
	"os"
	"path"

	rocksdb "github.com/tecbot/gorocksdb"
)

func OpenDbWithAllCFs(name string) (*rocksdb.DB, map[string]*rocksdb.ColumnFamilyHandle, error) {
	mo := NumberMergeOperator{}

	opts := rocksdb.NewDefaultOptions()
	opts.SetMergeOperator(mo)

	defer opts.Destroy()
	opts.SetCreateIfMissing(true)

	// how many LOG.old file to keep
	opts.SetKeepLogFileNum(10)
	opts.SetCompression(rocksdb.ZLibCompression)

	// ugly hack when db is not available
	if _, err := os.Stat(path.Join(name, "CURRENT")); err != nil {
		db, err := rocksdb.OpenDb(opts, name)

		return db, nil, err
	} else {
		// XXX: this won't work if the database does not exists
		cfNames, err := rocksdb.ListColumnFamilies(opts, name)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("column families: %v\n", cfNames)

		cfOpts := make([]*rocksdb.Options, len(cfNames))
		for i := range cfNames {
			cfOpts[i] = opts
		}

		db, cfs, err := rocksdb.OpenDbColumnFamilies(opts, name, cfNames, cfOpts)

		if err != nil {
			return nil, nil, err
		}

		cfMap := make(map[string]*rocksdb.ColumnFamilyHandle)

		for i := range cfNames {
			cfMap[cfNames[i]] = cfs[i]
		}

		return db, cfMap, nil
	}
}

func CreateCFExample(db *rocksdb.DB) error {
	cfOpts := rocksdb.NewDefaultOptions()
	defer cfOpts.Destroy()
	cf, err := db.CreateColumnFamily(cfOpts, "cf1")
	if err != nil {
		return err
	}
	defer cf.Destroy()

	return nil
}
