// The CDK Construct Library of Astro
package construct


// The options for the LambdaAstroSite.
// Experimental.
type LambdaAstroSiteProps struct {
	// The server entry file, e.g. path.join(__dirname, "../server/entry.mjs").
	// Experimental.
	ServerEntry *string `field:"required" json:"serverEntry" yaml:"serverEntry"`
	// The directory of static files, e.g. path.join(__dirname, "../dist/client").
	// Experimental.
	StaticDir *string `field:"required" json:"staticDir" yaml:"staticDir"`
	// The options for the CloudFront distribution.
	//
	// Recommended to use CloudFront for production.
	// Experimental.
	CfOptions *CfOptions `field:"optional" json:"cfOptions" yaml:"cfOptions"`
	// HttpApi Gateway options.
	// Experimental.
	GwOptions *GwOptions `field:"optional" json:"gwOptions" yaml:"gwOptions"`
	// The server options.
	// Experimental.
	ServerOptions *ServerOptions `field:"optional" json:"serverOptions" yaml:"serverOptions"`
}

