### About

PoC demonstrating the use of AWS S3 SDK:

* Create bucket
* Delete bucket items
* Delete bucket
* Write bucket items
* Read bucket items

And records timing metrics for each operation.

### AWS Results

```shell
deleteBucketItems:
deleteBucketItems duration: 620.660946ms

deleteBucket:
Waiting for bucket "example.5577006791947779410" to be deleted...
deleteBucket duration: 484.245965ms

createBucket:
Waiting for bucket "example.5577006791947779410" to be created...
Bucket "example.5577006791947779410" successfully created
createBucket duration: 838.556069ms

writeFile:
Successfully uploaded "tree" to "example.5577006791947779410"
writeFile duration: 62.002226ms

writeFile:
Successfully uploaded "blob1" to "example.5577006791947779410"
writeFile duration: 159.3472ms

writeFile:
Successfully uploaded "blob2" to "example.5577006791947779410"
writeFile duration: 187.575771ms

writeFile:
Successfully uploaded "index" to "example.5577006791947779410"
writeFile duration: 98.432172ms

readFile:
Downloaded tmp/tree 128 bytes
readFile duration: 48.616313ms

readFile:
Downloaded tmp/blob1 7670 bytes
readFile duration: 48.45348ms

readFile:
Downloaded tmp/blob2 9526 bytes
readFile duration: 54.116281ms

readFile:
Downloaded tmp/index 263 bytes
readFile duration: 58.240697ms

listAllBuckets:
Buckets:
* example.5577006791947779410 created on 2020-07-14 00:15:10 +0000 UTC
listAllBuckets duration: 58.908136ms

listBucket:
Name:          blob1
Last modified: 2020-07-14 00:15:11 +0000 UTC
Size:          7670
Storage class: STANDARD

Name:          blob2
Last modified: 2020-07-14 00:15:11 +0000 UTC
Size:          9526
Storage class: STANDARD

Name:          index
Last modified: 2020-07-14 00:15:11 +0000 UTC
Size:          263
Storage class: STANDARD

Name:          tree
Last modified: 2020-07-14 00:15:11 +0000 UTC
Size:          128
Storage class: STANDARD

listBucket duration: 56.044107ms
```

### Dependencies

Creates a session using AWS conventions and expects AWS credentials to exist in `~/.aws/credentials`.

### Build

```sh
make build
```

### Run

```sh
bin/example
```

### License

MIT
