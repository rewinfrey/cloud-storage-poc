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
createBucket duration: 1.161841512s
writeFile duration: 84.233117ms
writeFile duration: 220.977067ms
writeFile duration: 67.304721ms
writeFile duration: 131.919906ms
writeFile duration: 68.446034ms
writeFile duration: 130.787404ms
writeFile duration: 133.894671ms
writeFile duration: 114.2448ms
writeFile duration: 79.603226ms
writeFile duration: 86.643531ms
writeFile duration: 62.11291ms
writeFile duration: 87.90533ms
writeFile duration: 91.29069ms
writeFile duration: 212.076542ms
writeFile duration: 97.595439ms
writeFile duration: 69.844116ms
writeFile duration: 68.262244ms
writeFile duration: 81.477255ms
writeFile duration: 76.607526ms
writeFile duration: 72.132031ms
writeFile duration: 65.540061ms
writeFile duration: 108.550604ms
writeFile duration: 66.069561ms
writeFile duration: 66.791373ms
writeFile duration: 191.131882ms
writeFile duration: 120.192773ms
writeFile duration: 123.990689ms
writeFile duration: 68.177575ms
writeFile duration: 62.975746ms
writeFile duration: 66.940321ms
writeFile duration: 88.308985ms
writeFile duration: 78.190148ms
writeFile duration: 64.621205ms
writeFile duration: 96.059621ms
writeFile duration: 119.97724ms
writeFile duration: 70.669957ms
writeFile duration: 65.779812ms
writeFile duration: 118.441643ms
writeFile duration: 84.209545ms
writeFile duration: 61.931225ms
readFile duration: 60.191134ms
readFile duration: 53.938248ms
readFile duration: 54.413438ms
readFile duration: 54.498463ms
readFile duration: 54.570228ms
readFile duration: 67.306062ms
readFile duration: 58.358407ms
readFile duration: 60.10724ms
readFile duration: 51.783302ms
readFile duration: 60.207034ms
readFile duration: 54.170624ms
readFile duration: 147.128846ms
readFile duration: 69.798542ms
readFile duration: 53.503248ms
readFile duration: 55.645859ms
readFile duration: 78.515355ms
readFile duration: 52.811149ms
readFile duration: 58.392547ms
readFile duration: 55.591633ms
readFile duration: 66.485848ms
readFile duration: 64.323466ms
readFile duration: 49.823737ms
readFile duration: 52.647052ms
readFile duration: 53.485356ms
readFile duration: 58.174582ms
readFile duration: 55.965365ms
readFile duration: 56.975399ms
readFile duration: 51.999561ms
readFile duration: 53.632249ms
readFile duration: 55.602754ms
readFile duration: 52.064963ms
readFile duration: 72.86589ms
readFile duration: 55.476741ms
readFile duration: 51.675462ms
readFile duration: 54.281554ms
readFile duration: 55.898656ms
readFile duration: 62.412338ms
readFile duration: 57.599902ms
readFile duration: 54.243503ms
readFile duration: 120.729649ms
listAllBuckets duration: 391.217875ms
listBucket duration: 60.206904ms
deleteBucketItems duration: 657.801923ms
deleteBucket duration: 580.390264ms
```

### Azure Results

```
createContainer duration: 805.651097ms
writeFile duration: 120.904257ms
writeFile duration: 56.525181ms
writeFile duration: 61.141167ms
writeFile duration: 61.798165ms
writeFile duration: 56.417374ms
writeFile duration: 57.357617ms
writeFile duration: 85.13006ms
writeFile duration: 106.893505ms
writeFile duration: 60.080406ms
writeFile duration: 59.223696ms
writeFile duration: 59.557656ms
writeFile duration: 57.705416ms
writeFile duration: 70.006174ms
writeFile duration: 60.593492ms
writeFile duration: 59.893762ms
writeFile duration: 66.957528ms
writeFile duration: 64.603641ms
writeFile duration: 68.2491ms
writeFile duration: 115.684587ms
writeFile duration: 58.258856ms
writeFile duration: 53.514134ms
writeFile duration: 65.548473ms
writeFile duration: 73.725519ms
writeFile duration: 163.485985ms
writeFile duration: 76.106064ms
writeFile duration: 106.140691ms
writeFile duration: 63.372473ms
writeFile duration: 53.847126ms
writeFile duration: 55.602093ms
writeFile duration: 53.444554ms
writeFile duration: 61.719476ms
writeFile duration: 60.274745ms
writeFile duration: 88.868353ms
writeFile duration: 63.798292ms
writeFile duration: 60.882102ms
writeFile duration: 74.781366ms
writeFile duration: 52.450384ms
writeFile duration: 67.115764ms
writeFile duration: 56.269905ms
writeFile duration: 83.672887ms
readFile duration: 62.387672ms
readFile duration: 50.599848ms
readFile duration: 54.878362ms
readFile duration: 48.063101ms
readFile duration: 59.057487ms
readFile duration: 121.036759ms
readFile duration: 48.453459ms
readFile duration: 55.817957ms
readFile duration: 71.289002ms
readFile duration: 60.712534ms
readFile duration: 52.24889ms
readFile duration: 58.178247ms
readFile duration: 56.096223ms
readFile duration: 49.70057ms
readFile duration: 52.364893ms
readFile duration: 55.300163ms
readFile duration: 64.198184ms
readFile duration: 124.483575ms
readFile duration: 50.227683ms
readFile duration: 56.10527ms
readFile duration: 48.488952ms
readFile duration: 57.479839ms
readFile duration: 61.390577ms
readFile duration: 49.955612ms
readFile duration: 49.840566ms
readFile duration: 51.253752ms
readFile duration: 52.392804ms
readFile duration: 51.432015ms
readFile duration: 53.863554ms
readFile duration: 58.150968ms
readFile duration: 50.457554ms
readFile duration: 51.618929ms
readFile duration: 59.34131ms
readFile duration: 53.370225ms
readFile duration: 52.307458ms
readFile duration: 49.83252ms
readFile duration: 50.888567ms
readFile duration: 104.855961ms
readFile duration: 50.27452ms
readFile duration: 78.387852ms
listContainer duration: 111.329234ms
deleteContainer duration: 51.624311ms
```

### Key findings

Azure blob storage performance overall is consistently under 100 ms for reads and writes. Whereas for the same blob files, AWS overall performs similarly but with a greater variance (from 50ms to 500ms).

Azure blob storage's SDK has builtin retry logic for reads and is a little more fault tolerant than AWS SDK.

### Dependencies

For the AWS test, you will need AWS account credentials stored in `~/.aws/credentials`. See the [getting started with AWS S3 guide](https://aws.amazon.com/s3/getting-started/) for more info.

For the Azure test, you will need to export the following env vars:

```
AZURE_STORAGE_ACCOUNT
AZURE_STORAGE_ACCESS_KEY
```

Please see the [Azure blob storage quickstart guide](https://docs.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-go?tabs=linux) for more info.

### Data

File used for the write / read tests can be found in `data`. In total 41 files encoded in binary format were used and have a range of sizes from 1K to 50K.

### Build

```sh
make build
```

### Run

```sh
bin/aws
bin/azure
```

### License

MIT
