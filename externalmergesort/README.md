# External Merge Sort

0. Build the binary from the [cmd](./cmd) folder.

1. Generate a test file

```
./cmd -command generate -input 1000M.txt -count 1000000000
```

2. Sort the file using external merge sort with chunk size of 100M. The result will be written to `chunk/final`

```
./cmd -command sort -input 1000M.txt -chunk 100
```

3. Validate the file

```
./cmd -command validate -input chunk/final
```
