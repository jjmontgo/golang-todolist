package aws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"log"
	"time"
)
// Authenticating Requests in Browser-Based Uploads Using POST (AWS Signature Version 4)
// https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html

// Returns map with the following keys, to be added as hidden inputs to the file upload form:
// 		"aws_upload_url"
// 		"policy"
// 		"success_action_status"
// 		"success_action_redirect"
// 		"x_amz_algorithm"
// 		"x_amz_credential"
// 		"x_amz_date"
// 		"x_amz_signature"
func S3BrowserBasedUploadFormVariables(keyPath string, successActionStatus string, successActionRedirect string) map[string]string {

	vars := make(map[string]string)

	// Creating a POST Policy
	// https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-HTTPPOSTConstructPolicy.html

	awsUploadBucketName := os.Getenv("AWS_S3_UPLOAD_BUCKET_NAME")
	if awsUploadBucketName == "" {
		log.Println("AWS_S3_UPLOAD_BUCKET_NAME environment variable not set")
	}

	vars["aws_upload_url"] = "https://" + awsUploadBucketName + ".s3.amazonaws.com/"

	vars["x_amz_algorithm"] = "AWS4-HMAC-SHA256"

	currentTime := time.Now()
	date := currentTime.Format("2006") + currentTime.Format("01") + currentTime.Format("02")
	vars["x_amz_date"] = date + "T000000Z"

	awsAccessKeyId := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	if awsAccessKeyId == "" {
		log.Println("AWS_S3_ACCESS_KEY_ID environment variable not set")
	}
	awsRegion := os.Getenv("AWS_S3_REGION")
	if awsRegion == "" {
		log.Println("AWS_S3_REGION environment variable not set")
	}
	vars["x_amz_credential"] = awsAccessKeyId + "/" + date + "/" + awsRegion + "/s3/aws4_request"

	// expires in 1 hour
	expiration := currentTime.UTC().Add(time.Hour * time.Duration(1)).Format(time.RFC3339)

	policy :=
		"{\"expiration\": \"" + expiration + "\"," +
			"\"conditions\": [" +
				"{\"bucket\": \"" + awsUploadBucketName + "\" }," +
				"[\"starts-with\", \"$key\", \"" + keyPath + "\"]," +
				"{\"success_action_status\": \"" + successActionStatus + "\"}," +
				"{\"success_action_redirect\": \"" + successActionRedirect + "\"}," +
				"{\"x-amz-algorithm\": \"" + vars["x_amz_algorithm"] + "\"}," +
				"{\"x-amz-credential\": \"" + vars["x_amz_credential"] + "\"}," +
				"{\"x-amz-date\": \"" + vars["x_amz_date"] + "\"}," +
			"]" +
		"}"
	vars["policy"] = base64.StdEncoding.EncodeToString([]byte(policy))

	awsSecretAccessKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	if awsSecretAccessKey == "" {
		log.Println("AWS_S3_SECRET_ACCESS_KEY environment variable not set")
	}

	// Authenticating Requests: Browser-Based Uploads Using POST (AWS Signature Version 4)
	// https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-authentication-HTTPPOST.html

	dateKeyHmac := hmac.New(sha256.New, []byte("AWS4" + awsSecretAccessKey))
	dateKeyHmac.Write([]byte(date))
	dateKey := dateKeyHmac.Sum(nil)

	dateRegionKeyHmac := hmac.New(sha256.New, []byte(dateKey))
	dateRegionKeyHmac.Write([]byte(awsRegion))
	dateRegionKey := dateRegionKeyHmac.Sum(nil)

	dateRegionServiceKeyHmac := hmac.New(sha256.New, []byte(dateRegionKey))
	dateRegionServiceKeyHmac.Write([]byte("s3"))
	dateRegionServiceKey := dateRegionServiceKeyHmac.Sum(nil)

	signingKeyHmac := hmac.New(sha256.New, []byte(dateRegionServiceKey))
	signingKeyHmac.Write([]byte("aws4_request"))
	signingKey := signingKeyHmac.Sum(nil)

	signatureHmac := hmac.New(sha256.New, []byte(signingKey))
	signatureHmac.Write([]byte(vars["policy"]))

	vars["x_amz_signature"] = hex.EncodeToString(signatureHmac.Sum(nil))

	return vars
}
