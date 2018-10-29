# RocksDB example

First I need to clone and make rockdbs:

  * Clone rocksdb
  * `make shared_lib` to compile RocksDB

Make sure to set the following environments:

```
export ROCKSDB_PATH=~/Projects/sandbox/rocksdb
export CGO_CFLAGS="-I$ROCKSDB_PATH/include"
export CGO_LDFLAGS="-L$ROCKSDB_PATH -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4"
export LD_LIBRARY_PATH="$ROCKSDB_PATH:$LD_LIBRARY_PATH"
export DYLD_LIBRARY_PATH="$ROCKSDB_PATH:$DYLD_LIBRARY_PATH"
```

Remarks:

  * Interface is quite verbose, not as clean as BoltDB
  * You need to create Read/Write options. Make sure to free each of them.
  * There is no Go document, so you need to read Go code or [the document for C++](https://github.com/facebook/rocksdb/wiki)
  * Iterator is similar to Cursor in BoltDB
    * Tailing Iterator allows you to see new data as it come in.
  * Merge: Provide a way to perform read-modify-write operation, such as Counter
    * You need to define a merge operator
    * Merge happpens lazily, until a Put/Delete/Get
    * Full merge is to merge existing value and an operand
    * Partial merge to merge two operands. It can be used to speed up performance (reduce number of full merge)
  * Column Familiy are like Buckets in BoltDB
    * You need to make sure to specify column families when you open the DB. Or you need to open ALL column families.
    * It is quite ugly when you want to combine with `SetCreateIfMissing`
  * Compaction Filter: a background garbage collector to remove unused keys
