package aws

import (
	"bytes"
	"image/jpeg"
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
)

func ResizeS3ObjectImage(objectKey string, maxWidth uint, maxHeight uint) {
	session := GetClientConfigProvider()
	bucketName := os.Getenv("AWS_S3_UPLOAD_BUCKET_NAME")

	downloadBuffer := &aws.WriteAtBuffer{}
	s3Download := s3manager.NewDownloader(session)
	_, err := s3Download.Download(downloadBuffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Could not download from S3: %v", err)
		return
	}

	imageBytes := downloadBuffer.Bytes()
	reader := bytes.NewReader(imageBytes)
	image, err := jpeg.Decode(reader)
	if err != nil {
		log.Printf("%s: %s", objectKey, err)
		return
	}

	resizedImage := resize.Thumbnail(maxWidth, maxHeight, image, resize.Lanczos3)

	uploadBuffer := new(bytes.Buffer)
	err = jpeg.Encode(uploadBuffer, resizedImage, nil)
	if err != nil {
		log.Printf("JPEG encoding error: %v", err)
		return
	}

	s3Upload := s3manager.NewUploader(session)
	_, err = s3Upload.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(uploadBuffer.Bytes()),
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Failed to upload: %v", err)
		return
	}
}
