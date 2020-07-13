### About

PoC demonstrating the use of AWS S3 SDK:

* Create bucket
* Delete bucket items
* Delete bucket
* Write bucket items
* Read bucket items

And records timing metrics for each operation.

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
