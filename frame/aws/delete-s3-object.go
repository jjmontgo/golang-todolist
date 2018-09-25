package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"log"
)

func DeleteS3Object(objectKey string) {
	sess := getClientConfigProvider()
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

/**
 * If local, replace access keys in environment with S3 access keys and specify S3 region
 * If prod, simply establish a new session and return the ConfigProvider
 *
 * It's unclear to me why this works.  I suspect because S3 is accessed in prod through
 * the VPC endpoint, and because *anyone* from the VPC has access to S3, no credentials
 * are needed.  Meanwhile, remote access works through the IAM user I set up specifically
 * for S3.
 *
 * @return {client.ConfigProvider} A config provider instantiated differently depending on MODE
 */
func getClientConfigProvider() client.ConfigProvider {
	var sess client.ConfigProvider
	if (os.Getenv("MODE") == "prod") {
		sess = session.Must(session.NewSession())
	} else {
		os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("AWS_S3_ACCESS_KEY_ID"))
		os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_S3_SECRET_ACCESS_KEY"))
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_S3_REGION"))},
		))
	}
	return sess
}
