//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package construct

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AstroSiteConstruct) validateNewBucketParameters(scope constructs.Construct, staticDir *string, props *awss3.BucketProps) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewS3OriginParameters(scope constructs.Construct, bucket awss3.Bucket) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateParseRoutesFromDirParameters(dir *string) error {
	return nil
}

func validateAstroSiteConstruct_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAstroSiteConstructParameters(scope constructs.Construct, id *string) error {
	return nil
}

