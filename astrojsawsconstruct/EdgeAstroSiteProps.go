// The CDK Construct Library of Astro
package astrojsawsconstruct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The options for the EdgeAstroSite.
// Experimental.
type EdgeAstroSiteProps struct {
	// The server entry file, e.g. path.join(__dirname, "../server/entry.mjs").
	// Experimental.
	ServerEntry *string `field:"required" json:"serverEntry" yaml:"serverEntry"`
	// The directory of static files, e.g. path.join(__dirname, "../dist/client").
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
	// The server options.
	// Experimental.
	ServerOptions *ServerOptions `field:"optional" json:"serverOptions" yaml:"serverOptions"`
}

