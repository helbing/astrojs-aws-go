//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package construct

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EdgeAstroSite) validateNewBucketParameters(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateNewDistributionParameters(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateNewHttpApiGatewayOriginParameters(httpApi awscdkapigatewayv2alpha.HttpApi) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateNewHttpApiGwParameters(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateNewS3OriginParameters(scope constructs.Construct, bucket awss3.Bucket) error {
	return nil
}

func (e *jsiiProxy_EdgeAstroSite) validateParseRoutesFromDirParameters(dir *string) error {
	return nil
}

func validateEdgeAstroSite_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEdgeAstroSiteParameters(scope constructs.Construct, id *string, props *EdgeAstroSiteProps) error {
	return nil
}

