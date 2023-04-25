// The CDK Construct Library of Astro
package construct

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The options for the Assets.
// Experimental.
type AssetsOptions struct {
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	// Experimental.
	Cors *[]*awss3.CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// Error document for the website.
	// Experimental.
	Errorhtml *string `field:"optional" json:"errorhtml" yaml:"errorhtml"`
	// Index document for the website.
	// Experimental.
	Indexhtml *string `field:"optional" json:"indexhtml" yaml:"indexhtml"`
}

