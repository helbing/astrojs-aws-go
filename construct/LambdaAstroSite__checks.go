//go:build !no_runtime_type_checking

// The CDK Construct Library of Astro
package construct

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaAstroSite) validateNewBucketParameters(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if whEnabled == nil {
		return fmt.Errorf("parameter whEnabled is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewDistributionParameters(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if defaultBehavior == nil {
		return fmt.Errorf("parameter defaultBehavior is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(defaultBehavior, func() string { return "parameter defaultBehavior" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewFunctionParameters(scope constructs.Construct, serverEntry *string, props *ServerOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if serverEntry == nil {
		return fmt.Errorf("parameter serverEntry is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewHttpApiGatewayOriginParameters(httpApi awscdkapigatewayv2alpha.HttpApi) error {
	if httpApi == nil {
		return fmt.Errorf("parameter httpApi is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewHttpApiGwParameters(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if fn == nil {
		return fmt.Errorf("parameter fn is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateNewS3OriginParameters(scope constructs.Construct, bucket awss3.Bucket) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaAstroSite) validateParseRoutesFromDirParameters(dir *string) error {
	if dir == nil {
		return fmt.Errorf("parameter dir is required, but nil was provided")
	}

	return nil
}

func validateLambdaAstroSite_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaAstroSiteParameters(scope constructs.Construct, id *string, props *LambdaAstroSiteProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

