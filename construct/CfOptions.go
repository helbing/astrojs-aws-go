// The CDK Construct Library of Astro
package construct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// CloudFront options.
// Experimental.
type CfOptions struct {
	// Use a custom certificate for the distribution from AWS Certificate Manager (ACM).
	// See: https://aws.amazon.com/premiumsupport/knowledge-center/custom-ssl-certificate-cloudfront/
	//
	// Experimental.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// Domains of the website.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The CloudFront functions to invoke before serving the contents.
	// Experimental.
	CfFunctions *[]*awscloudfront.FunctionAssociation `field:"optional" json:"cfFunctions" yaml:"cfFunctions"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	// Experimental.
	EdgeFunctions *[]*awscloudfront.EdgeLambda `field:"optional" json:"edgeFunctions" yaml:"edgeFunctions"`
	// How CloudFront should handle requests that are not successful (e.g., PageNotFound).
	// Experimental.
	ErrorResponses *[]*awscloudfront.ErrorResponse `field:"optional" json:"errorResponses" yaml:"errorResponses"`
	// Controls the countries in which your content is distributed.
	// Experimental.
	GeoRestriction awscloudfront.GeoRestriction `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// The Amazon S3 bucket to store the access logs in.
	// Experimental.
	LogBucket awss3.IBucket `field:"optional" json:"logBucket" yaml:"logBucket"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this distribution.
	// Experimental.
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Specifies whether you want CloudFront to include cookies in access logs.
	// Experimental.
	LogIncludesCookies *bool `field:"optional" json:"logIncludesCookies" yaml:"logIncludesCookies"`
	// The price class for the distribution (this impacts how many locations CloudFront uses for your distribution, and billing).
	// Experimental.
	PriceClass awscloudfront.PriceClass `field:"optional" json:"priceClass" yaml:"priceClass"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	//
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	// Experimental.
	WebACLId *string `field:"optional" json:"webACLId" yaml:"webACLId"`
}

