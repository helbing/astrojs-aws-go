//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package construct

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StaticAstroSite) validateNewBucketParameters(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewDistributionParameters(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewHttpApiGatewayOriginParameters(httpApi awscdkapigatewayv2alpha.HttpApi) error {
	return nil
}

func (s *jsiiProxy_StaticAstroSite) validateNewHttpApiGwParameters(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) error {
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

