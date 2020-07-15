### About

PoC demonstrating the use of AWS S3 SDK and Azure Blob Storage:

* Create bucket
* Delete bucket items
* Delete bucket
* Write bucket items
* Read bucket items

And records timing metrics for each operation.

### AWS Results

```shell
deleteBucketItems duration: 1.349719927s
deleteBucket duration: 547.164294ms
createBucket duration: 1.270393972s
writeFile duration: 89.186308ms
writeFile duration: 180.911338ms
writeFile duration: 240.82832ms
writeFile duration: 202.042548ms
writeFile duration: 58.803466ms
writeFile duration: 171.244482ms
writeFile duration: 88.472728ms
writeFile duration: 737.18784ms
writeFile duration: 62.67784ms
writeFile duration: 70.309715ms
writeFile duration: 65.327159ms
writeFile duration: 150.434911ms
writeFile duration: 63.947327ms
writeFile duration: 62.73731ms
writeFile duration: 118.315474ms
writeFile duration: 152.002883ms
writeFile duration: 72.416319ms
writeFile duration: 217.690247ms
writeFile duration: 230.460652ms
writeFile duration: 123.886487ms
writeFile duration: 63.74425ms
writeFile duration: 97.072228ms
writeFile duration: 66.61683ms
writeFile duration: 81.142789ms
writeFile duration: 64.020091ms
writeFile duration: 136.394217ms
writeFile duration: 69.461851ms
writeFile duration: 108.172827ms
writeFile duration: 332.925741ms
writeFile duration: 86.820791ms
writeFile duration: 87.370231ms
writeFile duration: 110.25239ms
writeFile duration: 100.975776ms
writeFile duration: 76.672588ms
writeFile duration: 68.963028ms
writeFile duration: 245.135691ms
writeFile duration: 160.685921ms
writeFile duration: 78.041691ms
writeFile duration: 805.599891ms
writeFile duration: 76.396857ms
readFile duration: 70.452756ms
readFile duration: 56.099654ms
readFile duration: 87.838749ms
readFile duration: 55.935787ms
readFile duration: 83.426588ms
readFile duration: 57.763171ms
readFile duration: 54.582571ms
readFile duration: 53.855343ms
readFile duration: 65.837233ms
readFile duration: 52.17258ms
readFile duration: 100.311028ms
readFile duration: 50.744022ms
readFile duration: 55.529471ms
readFile duration: 58.768819ms
readFile duration: 107.473226ms
readFile duration: 91.264562ms
readFile duration: 59.558916ms
readFile duration: 53.434542ms
readFile duration: 113.721101ms
readFile duration: 66.497887ms
readFile duration: 54.471983ms
readFile duration: 55.943433ms
readFile duration: 52.801834ms
readFile duration: 72.007989ms
readFile duration: 52.35898ms
readFile duration: 68.15785ms
readFile duration: 51.958272ms
readFile duration: 54.776943ms
readFile duration: 57.262742ms
readFile duration: 59.974658ms
readFile duration: 53.035263ms
readFile duration: 57.000806ms
readFile duration: 53.139609ms
readFile duration: 65.096253ms
readFile duration: 58.129643ms
readFile duration: 59.713496ms
readFile duration: 57.974573ms
readFile duration: 54.658018ms
readFile duration: 61.669057ms
readFile duration: 63.094626ms
listAllBuckets duration: 365.719981ms
listBucket duration: 64.095717ms
```

### Azure Results

```
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
