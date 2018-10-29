# Bolt

An example of trying [Bolt](https://github.com/etcd-io/bbolt)

I prefer this because this is purely in Go, no need to compile external library like LevelDB or RocksDB.

Remarks:

 * ACID support. 
   * But do we really need ACID at this level?
   * Since this is a single process application. ACID can be achieved easily using a lock.
 * Use Cursor.Last() & Cursor.Prev() to navigate from the end of the key space.
 * Read() guarantees to work
 * Batch() blocks until all transactions finished. The transaction function may be called multiple time so the function must be idempotent

References:
 * https://mycodesmells.com/post/first-steps-with-boltdb
 * [Badger vs Boltdb](https://blog.dgraph.io/post/badger-lmdb-boltdb/)
 * [The new InfluxDB storage engine: from LSM Tree to B+Tree and back again to create the Time Structured Merge Tree](https://docs.influxdata.com/influxdb/v1.6/concepts/storage_engine/)
