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
