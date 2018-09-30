package aws

import (
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteS3Object(objectKey string) {
	sess := GetClientConfigProvider()
	svc := s3.New(sess)
	bucket := os.Getenv("AWS_S3_UPLOAD_BUCKET_NAME")

	deleteObjectInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(objectKey),
	}
	_, err := svc.DeleteObject(deleteObjectInput)
	if err != nil {
		log.Println("Unable to delete object %q from bucket %q, %v", objectKey, bucket, err)
	}
}

