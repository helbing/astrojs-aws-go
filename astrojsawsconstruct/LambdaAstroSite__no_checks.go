//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package astrojsawsconstruct

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaAstroSite) validateNewBucketParameters(scope constructs.Construct, staticDir *string, props *awss3.BucketProps) error {
	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewS3OriginParameters(scope constructs.Construct, bucket awss3.Bucket) error {
	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateParseRoutesFromDirParameters(dir *string) error {
	return nil
}

func validateLambdaAstroSite_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaAstroSiteParameters(scope constructs.Construct, id *string, props *LambdaAstroSiteProps) error {
	return nil
}

