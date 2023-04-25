// The CDK Construct Library of Astro
package construct


// The options for the EdgeAstroSite.
// Experimental.
type EdgeAstroSiteProps struct {
	// The server entry file, e.g. path.join(__dirname, "../server/entry.mjs").
	// Experimental.
	ServerEntry *string `field:"required" json:"serverEntry" yaml:"serverEntry"`
	// The directory of static files, e.g. path.join(__dirname, "../dist/client").
	// Experimental.
	StaticDir *string `field:"required" json:"staticDir" yaml:"staticDir"`
	// The options for the CloudFront distribution.
	//
	// CloudFront is required, unless `onlyLambda` is true.
	// Experimental.
	CfOptions *CfOptions `field:"optional" json:"cfOptions" yaml:"cfOptions"`
	// Only deploy the lambda function for testing, no S3 Bucket and CloudFront.
	//
	// Edge function only works in CloudFront, but it really deploy too slow.
	// Experimental.
	OnlyLambda *bool `field:"optional" json:"onlyLambda" yaml:"onlyLambda"`
	// The server options.
	// Experimental.
	ServerOptions *ServerOptions `field:"optional" json:"serverOptions" yaml:"serverOptions"`
}

