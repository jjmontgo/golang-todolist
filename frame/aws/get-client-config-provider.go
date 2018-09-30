package aws

import (
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
)

/**
 * Configure and return an AWS session
 *
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

var sessionInitialized bool
var sess client.ConfigProvider

func GetClientConfigProvider() client.ConfigProvider {
	if sessionInitialized {
		return sess
	}

	if (os.Getenv("MODE") == "prod") {
		sess = session.Must(session.NewSession())
	} else {
		os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("AWS_S3_ACCESS_KEY_ID"))
		os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_S3_SECRET_ACCESS_KEY"))
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_S3_REGION"))},
		))
	}

	sessionInitialized = true
	return sess
}
