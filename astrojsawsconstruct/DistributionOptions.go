// The CDK Construct Library of Astro
package astrojsawsconstruct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// CloudFront distribution which is based on CDK DistributionProps, remove defaultBehavior, defaultRootObject.
// Experimental.
type DistributionOptions struct {
	// Additional behaviors for the distribution, mapped by the pathPattern that specifies which requests to apply the behavior to.
	// Experimental.
	AdditionalBehaviors *map[string]*awscloudfront.BehaviorOptions `field:"optional" json:"additionalBehaviors" yaml:"additionalBehaviors"`
	// A certificate to associate with the distribution.
	//
	// The certificate must be located in N. Virginia (us-east-1).
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Any comments you want to include about the distribution.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Alternative domain names for this distribution.
	//
	// If you want to use your own domain name, such as www.example.com, instead of the cloudfront.net domain name,
	// you can add an alternate domain name to your distribution. If you attach a certificate to the distribution,
	// you must add (at least one of) the domain names of the certificate to this list.
	// Experimental.
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Enable or disable the distribution.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether CloudFront will respond to IPv6 DNS requests with an IPv6 address.
	//
	// If you specify false, CloudFront responds to IPv6 DNS requests with the DNS response code NOERROR and with no IP addresses.
	// This allows viewers to submit a second request, for an IPv4 address for your distribution.
	// Experimental.
	EnableIpv6 *bool `field:"optional" json:"enableIpv6" yaml:"enableIpv6"`
	// Enable access logging for the distribution.
	// Experimental.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// How CloudFront should handle requests that are not successful (e.g., PageNotFound).
	// Experimental.
	ErrorResponses *[]*awscloudfront.ErrorResponse `field:"optional" json:"errorResponses" yaml:"errorResponses"`
	// Controls the countries in which your content is distributed.
	// Experimental.
	GeoRestriction awscloudfront.GeoRestriction `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront.
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	// Experimental.
	HttpVersion awscloudfront.HttpVersion `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// The Amazon S3 bucket to store the access logs in.
	// Experimental.
	LogBucket awss3.IBucket `field:"optional" json:"logBucket" yaml:"logBucket"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this distribution.
	// Experimental.
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Specifies whether you want CloudFront to include cookies in access logs.
	// Experimental.
	LogIncludesCookies *bool `field:"optional" json:"logIncludesCookies" yaml:"logIncludesCookies"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	// Experimental.
	MinimumProtocolVersion awscloudfront.SecurityPolicyProtocol `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify PriceClass_All, CloudFront responds to requests for your objects from all CloudFront edge locations.
	// If you specify a price class other than PriceClass_All, CloudFront serves your objects from the CloudFront edge location
	// that has the lowest latency among the edge locations in your price class.
	// Experimental.
	PriceClass awscloudfront.PriceClass `field:"optional" json:"priceClass" yaml:"priceClass"`
	// The SSL method CloudFront will use for your distribution.
	//
	// Server Name Indication (SNI) - is an extension to the TLS computer networking protocol by which a client indicates
	// which hostname it is attempting to connect to at the start of the handshaking process. This allows a server to present
	// multiple certificates on the same IP address and TCP port number and hence allows multiple secure (HTTPS) websites
	// (or any other service over TLS) to be served by the same IP address without requiring all those sites to use the same certificate.
	//
	// CloudFront can use SNI to host multiple distributions on the same IP - which a large majority of clients will support.
	//
	// If your clients cannot support SNI however - CloudFront can use dedicated IPs for your distribution - but there is a prorated monthly charge for
	// using this feature. By default, we use SNI - but you can optionally enable dedicated IPs (VIP).
	//
	// See the CloudFront SSL for more details about pricing : https://aws.amazon.com/cloudfront/custom-ssl-domains/
	// Experimental.
	SslSupportMethod awscloudfront.SSLMethod `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	// Experimental.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

