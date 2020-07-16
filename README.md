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

Creates a session using AWS conventions and expects AWS credentials to exist in `~/.aws/credentials`.

For the Azure test, you need to export the following env vars:

```
AZURE_STORAGE_ACCOUNT
AZURE_STORAGE_ACCESS_KEY
```

Both the AWS and Azure credentials require configuration and setup of AWS and Azure portal accounts.

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
