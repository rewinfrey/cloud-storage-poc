package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(exit bool, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	if exit {
		os.Exit(1)
	}
}

func listAllBuckets(svc *s3.S3) func() {
	return func() {
		result, err := svc.ListBuckets(nil)
		if err != nil {
			exitErrorf(true, "Unable to list buckets, %v", err)
		}

		fmt.Println("Buckets:")

		for _, b := range result.Buckets {
			fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
		}
	}
}

func createBucket(svc *s3.S3, name string) func() {
	return func() {
		_, err := svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(name),
		})
		if err != nil {
			exitErrorf(true, "Unable to create bucket %q, %v", name, err)
		}

		// Wait until bucket is created before finishing
		fmt.Printf("Waiting for bucket %q to be created...\n", name)

		err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(name),
		})
		if err != nil {
			exitErrorf(true, "Error occurred while waiting for bucket to be created, %v", name)
		}

		fmt.Printf("Bucket %q successfully created\n", name)
	}
}

func writeBucket(uploader *s3manager.Uploader, bucket string, filepath string, filename string) func() {
	return func() {
		file, err := os.Open(filepath)
		if err != nil {
			exitErrorf(true, "Unable to open file %q, %v", err)
		}

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
			Body:   file,
		})
		if err != nil {
			// Print the error and exit.
			exitErrorf(true, "Unable to upload %q to %q, %v", filepath, bucket, err)
		}

		fmt.Printf("Successfully uploaded %q to %q\n", filepath, bucket)
	}
}

func readBucket(downloader *s3manager.Downloader, bucket string, filepath string) func() {
	return func() {
		file, err := os.Create("tmp/data" + filepath)
		if err != nil {
			exitErrorf(true, "Unable to open file %q, %v", filepath, err)
		}
		defer file.Close()

		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(filepath),
			})
		if err != nil {
			exitErrorf(true, "Unable to download item %q, %v", filepath, err)
		}

		fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	}
}

func listBucket(svc *s3.S3, bucket string) func() {
	return func() {
		resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
		if err != nil {
			exitErrorf(true, "Unable to list items in bucket %q, %v", bucket, err)
		}

		for _, item := range resp.Contents {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
}

func deleteBucketItems(svc *s3.S3, bucket string) func() {
	return func() {
		iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
			Bucket: aws.String(bucket),
		})

		if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
			exitErrorf(false, "Unable to delete objects from bucket %q, %v", bucket, err)
		}
	}
}

func deleteBucket(svc *s3.S3, bucket string) func() {
	return func() {
		_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			exitErrorf(false, "Unable to delete bucket %q, %v", bucket, err)
		}

		// Wait until bucket is deleted before finishing
		fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

		err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucket),
		})
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

func main() {
	// Configure new Amazon session (reads from ~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		exitErrorf(true, "Unable to configure new Amazon session, %v", err)
		return
	}

	// Create S3 service client.
	svc := s3.New(sess)

	// Create S3 file uploader and downloader.
	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)
	logfile, err := os.Create("tmp/log.txt")
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	// Use a pseudo-randomly generated name.
	name := "blob-storage"

	// Clean up previous runs (if needed).
	withTiming(deleteBucketItems(svc, name), "deleteBucketItems", logfile)
	withTiming(deleteBucket(svc, name), "deleteBucket", logfile)

	// Create the bucket.
	withTiming(createBucket(svc, name), "createBucket", logfile)

	// Write files to the bucket.
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
		panic(err)
	}
	for filename, filepath := range files {
		withTiming(writeBucket(uploader, name, filepath, filename), "writeFile", logfile)
	}

	// Read files from the bucket.
	for filename, _ := range files {
		withTiming(readBucket(downloader, name, filename), "readFile", logfile)
	}

	// List all buckets.
	withTiming(listAllBuckets(svc), "listAllBuckets", logfile)

	// List all within a single bucket.
	withTiming(listBucket(svc, name), "listBucket", logfile)
}
