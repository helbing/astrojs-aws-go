//go:build no_runtime_type_checking

// The CDK Construct Library of Astro
package construct

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AstroSiteConstruct) validateNewBucketParameters(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewDistributionParameters(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewHttpApiGatewayOriginParameters(httpApi awscdkapigatewayv2alpha.HttpApi) error {
	return nil
}

func (a *jsiiProxy_AstroSiteConstruct) validateNewHttpApiGwParameters(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) error {
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

