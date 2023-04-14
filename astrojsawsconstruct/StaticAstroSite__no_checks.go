//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package astrojsawsconstruct

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StaticAstroSite) validateNewBucketParameters(scope constructs.Construct, staticDir *string, props *awss3.BucketProps) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewS3OriginParameters(scope constructs.Construct, bucket awss3.Bucket) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateParseRoutesFromDirParameters(dir *string) error {
	return nil
}

func validateStaticAstroSite_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStaticAstroSiteParameters(scope constructs.Construct, id *string, props *StaticAstroSiteProps) error {
	return nil
}

