package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func listAllBuckets(svc *s3.S3) {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func createBucket(svc *s3.S3, name string) {
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", name, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", name)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", name)
	}

	fmt.Printf("Bucket %q successfully created\n", name)
}

func writeBucket(uploader *s3manager.Uploader, bucket string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func readBucket(downloader *s3manager.Downloader, bucket string, filename string) {
	file, err := os.Create("tmp/" + filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", filename, err)
	}
	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", filename, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func listBucket(svc *s3.S3, bucket string) {
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func deleteBucket(svc *s3.S3, bucket string) {
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
		exitErrorf("Unable to delete objects from bucket %q, %v", bucket, err)
	}

	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait until bucket is deleted before finishing
	fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
}

func main() {
	// Configure new Amazon session (reads from ~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		exitErrorf("Unable to configure new Amazon session, %v", err)
		return
	}

	// Create S3 service client.
	svc := s3.New(sess)

	// Create S3 file uploader and downloader.
	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)

	// Use a pseudo-randomly generated name.
	name := "example." + strconv.Itoa(rand.Int())

	// Clean up previous runs (if needed).
	deleteBucket(svc, name)

	// Create the bucket.
	createBucket(svc, name)

	// Write files to the bucket.
	writeBucket(uploader, name, "tree")
	writeBucket(uploader, name, "blob1")
	writeBucket(uploader, name, "blob2")
	writeBucket(uploader, name, "index")

	// Read files from the bucket.
	readBucket(downloader, name, "tree")
	readBucket(downloader, name, "blob1")
	readBucket(downloader, name, "blob2")
	readBucket(downloader, name, "index")

	// List all buckets.
	listAllBuckets(svc)

	// List all within a single bucket.
	listBucket(svc, name)
}
