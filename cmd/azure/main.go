package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// This is based on https://github.com/Azure-Samples/storage-blobs-go-quickstart.

// Azure Storage Quickstart Sample - Demonstrate how to upload, list, download, and delete blobs.
//
// Documentation References:
// - What is a Storage Account - https://docs.microsoft.com/azure/storage/common/storage-create-storage-account
// - Blob Service Concepts - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-Concepts
// - Blob Service Go SDK API - https://godoc.org/github.com/Azure/azure-storage-blob-go
// - Blob Service REST API - https://docs.microsoft.com/rest/api/storageservices/Blob-Service-REST-API
// - Scalability and performance targets - https://docs.microsoft.com/azure/storage/common/storage-scalability-targets
// - Azure Storage Performance and Scalability checklist https://docs.microsoft.com/azure/storage/common/storage-performance-checklist
// - Storage Emulator - https://docs.microsoft.com/azure/storage/common/storage-use-emulator

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			switch serr.ServiceCode() { // Compare serviceCode to ServiceCodeXxx constants
			case azblob.ServiceCodeContainerAlreadyExists:
				fmt.Println("Received 409. Container already exists")
				return
			}
		}
		log.Fatal(err)
	}
}

func withTiming(op func(), opName string, logfile *os.File) {
	start := time.Now()
	defer func() {
		msg := fmt.Sprintf("%s duration: %v\n", opName, time.Since(start))
		logfile.WriteString(msg)
		logfile.Sync()
		fmt.Println(msg)
	}()
	fmt.Println("\n" + opName + ":")
	op()
}

func createContainer(ctx context.Context, name string, containerURL azblob.ContainerURL) func() {
	return func() {
		_, err := containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
		handleErrors(err)
	}
}

func writeFile(ctx context.Context, containerURL azblob.ContainerURL, filepath string, filename string) func() {
	return func() {
		// Here's how to upload a blob.
		blobURL := containerURL.NewBlockBlobURL(filename)
		file, err := os.Open(filepath)
		handleErrors(err)

		// You can use the low-level PutBlob API to upload files. Low-level APIs are simple wrappers for the Azure Storage REST APIs.
		// Note that PutBlob can upload up to 256MB data in one shot. Details: https://docs.microsoft.com/en-us/rest/api/storageservices/put-blob
		// Following is commented out intentionally because we will instead use UploadFileToBlockBlob API to upload the blob
		// _, err = blobURL.PutBlob(ctx, file, azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
		// handleErrors(err)

		// The high-level API UploadFileToBlockBlob function uploads blocks in parallel for optimal performance, and can handle large files as well.
		// This function calls PutBlock/PutBlockList for files larger 256 MBs, and calls PutBlob for any file smaller
		_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16})
		handleErrors(err)
	}
}

func readFile(ctx context.Context, containerURL azblob.ContainerURL, filename string) func() {
	return func() {
		file, err := os.Create("tmp/azure" + filename)
		if err != nil {
			handleErrors(err)
		}
		defer file.Close()

		// First construct the blobURL.
		blobURL := containerURL.NewBlockBlobURL(filename)

		// Here's how to download the blob
		downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)

		// NOTE: automatically retries are performed if the connection fails
		bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

		// read the body into a buffer
		downloadedData := bytes.Buffer{}
		_, err = downloadedData.ReadFrom(bodyStream)
		handleErrors(err)

		downloadedData.WriteTo(file)
	}
}

func listContainer(ctx context.Context, containerURL azblob.ContainerURL) func() {
	return func() {
		// List the container that we have created above
		fmt.Println("Listing the blobs in the container:")
		for marker := (azblob.Marker{}); marker.NotDone(); {
			// Get a result segment starting with the blob indicated by the current Marker.
			listBlob, err := containerURL.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
			handleErrors(err)

			// ListBlobs returns the start of the next segment; you MUST use this to get
			// the next segment (after processing the current result segment).
			marker = listBlob.NextMarker

			// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
			for _, blobInfo := range listBlob.Segment.BlobItems {
				fmt.Print("	Blob name: " + blobInfo.Name + "\n")
			}
		}
	}
}

func deleteContainer(ctx context.Context, containerURL azblob.ContainerURL) func() {
	return func() {
		containerURL.Delete(ctx, azblob.ContainerAccessConditions{})
	}
}

func main() {
	logfile, err := os.Create("azure_log.txt")
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	ctx := context.Background()

	// From the Azure portal, get your storage account name and key and set environment variables.
	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Define the container name.
	containerName := "blob-storage"

	// From the Azure portal, get your storage account blob service URL endpoint.
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)

	// Create the container.
	withTiming(createContainer(ctx, containerName, containerURL), "createContainer", logfile)

	// Write files to the container.
	files := make(map[string]string)
	root := "data/"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files[info.Name()] = path
		return nil
	})
	if err != nil {
		handleErrors(err)
	}
	for filename, filepath := range files {
		withTiming(writeFile(ctx, containerURL, filepath, filename), "writeFile", logfile)
	}

	// Read files from the bucket.
	for filename, _ := range files {
		withTiming(readFile(ctx, containerURL, filename), "readFile", logfile)
	}

	// List items in the container.
	withTiming(listContainer(ctx, containerURL), "listContainer", logfile)

	// Delete the container.
	withTiming(deleteContainer(ctx, containerURL), "deleteContainer", logfile)

}
