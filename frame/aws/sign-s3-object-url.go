package aws

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

/**
 * Step by step instructions at:
 * https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-creating-signed-url-canned-policy.html
 *
 * @param objectKey The path and filename key of the uploaded object with no leading slash
 * @return The signed URL for the object with a 1 hour expiry
 *
 * Final URL format:
 * <cloudfrontUrl>/<objectKey>? \
 * 	Policy=<base64encodedpolicy>&
 * 	Signature=<hashedAndSignedVersionOfPolicy>&
 * 	Key-Pair-Id=<CloudfrontKeyPairId>
*/
func SignS3ObjectUrl(objectKey string) string {
	baseURL := os.Getenv("CLOUDFRONT_PRIVATE_URL")
	cloudfrontKeyPairId := os.Getenv("CLOUDFRONT_KEY_PAIR_ID")

	cloudfrontPrivateKey := os.Getenv("CLOUDFRONT_PRIVATE_KEY")
	// key is stored on a single line with newlines escaped; unescape them
	cloudfrontPrivateKey = strings.Replace(cloudfrontPrivateKey, "\\n", "\n", -1) // -1 means replace all
	pem, _ := pem.Decode([]byte(cloudfrontPrivateKey))
	rsaPrivateKey, _ := x509.ParsePKCS1PrivateKey(pem.Bytes)

	// time.Time 1 hour from now
	expiryTime := time.Now().Local().Add(time.Hour * time.Duration(1))

	resource := baseURL + "/" + objectKey
	newCannedPolicy := sign.NewCannedPolicy(resource, expiryTime)

	// get b64 encoded signature
	b64Signature, _, err := newCannedPolicy.Sign(rsaPrivateKey)
	if (err != nil) {
		log.Fatalln(err)
	}

	expiryString := strconv.FormatInt(expiryTime.Unix(), 10)

	return resource +
		"?Expires=" + expiryString +
		"&Signature=" + string(b64Signature) +
		"&Key-Pair-Id=" + cloudfrontKeyPairId
}
