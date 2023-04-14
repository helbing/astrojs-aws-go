// The CDK Construct Library of Astro
package construct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The options for the StaticAstroSite.
// Experimental.
type StaticAstroSiteProps struct {
	// The directory of built files, e.g. path.join(__dirname, "../dist").
	// Experimental.
	StaticDir *string `field:"required" json:"staticDir" yaml:"staticDir"`
	// Bucket options which is based on BucketProps.
	//
	// removalPolicy: @default RemovalPolicy.DESTROY
	// autoDeleteObjects @default true,.
	// Experimental.
	BucketOptions *awss3.BucketProps `field:"optional" json:"bucketOptions" yaml:"bucketOptions"`
	// CloudFront distribution default behavior options.
	// Experimental.
	DistributionDefaultBehaviorOptions *awscloudfront.AddBehaviorOptions `field:"optional" json:"distributionDefaultBehaviorOptions" yaml:"distributionDefaultBehaviorOptions"`
	// CloudFront distribution options.
	// Experimental.
	DistributionOptions *DistributionOptions `field:"optional" json:"distributionOptions" yaml:"distributionOptions"`
}

