package main

import (
	"encoding/binary"
	"fmt"

	rocksdb "github.com/tecbot/gorocksdb"
)

type NumberMergeOperator struct {
	rocksdb.MergeOperator
}

func (mo NumberMergeOperator) Name() string {
	return "NumberMergeOperator"
}

func (mo NumberMergeOperator) FullMerge(key, existingValue []byte, operands [][]byte) ([]byte, bool) {
	var current uint16 = 0
	var result = make([]byte, 4)

	if len(existingValue) > 0 {
		current = binary.BigEndian.Uint16(existingValue)
	}

	for _, v := range operands {
		op := binary.BigEndian.Uint16(v)
		current = current + op
	}

	binary.BigEndian.PutUint16(result, current)
	return result, true
}

func (mo NumberMergeOperator) PartialMerge(key, leftOperand, rightOperand []byte) ([]byte, bool) {
	return nil, false
}

func MergeExample(db *rocksdb.DB) error {
	opts := rocksdb.NewDefaultWriteOptions()
	defer opts.Destroy()

	bs := make([]byte, 4)
	binary.BigEndian.PutUint16(bs, 1)

	err := db.Merge(opts, []byte("add_me"), bs)

	if err != nil {
		return err
	}

	readOpts := rocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	// this will materialize the merge
	slice, err := db.Get(readOpts, []byte("add_me"))

	if err != nil {
		return err
	}

	defer slice.Free()
	fmt.Printf("After merge: %v\n", binary.BigEndian.Uint16(slice.Data()))

	return nil
}
