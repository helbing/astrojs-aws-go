package construct

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.AssetsOptions",
		reflect.TypeOf((*AssetsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@astrojs-aws/construct.AstroSiteConstruct",
		reflect.TypeOf((*AstroSiteConstruct)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "newBucket", GoMethod: "NewBucket"},
			_jsii_.MemberMethod{JsiiMethod: "newDistribution", GoMethod: "NewDistribution"},
			_jsii_.MemberMethod{JsiiMethod: "newFunction", GoMethod: "NewFunction"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGatewayOrigin", GoMethod: "NewHttpApiGatewayOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGw", GoMethod: "NewHttpApiGw"},
			_jsii_.MemberMethod{JsiiMethod: "newS3Origin", GoMethod: "NewS3Origin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseRoutesFromDir", GoMethod: "ParseRoutesFromDir"},
			_jsii_.MemberMethod{JsiiMethod: "strToRuntime", GoMethod: "StrToRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AstroSiteConstruct{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.CfOptions",
		reflect.TypeOf((*CfOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@astrojs-aws/construct.EdgeAstroSite",
		reflect.TypeOf((*EdgeAstroSite)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "domains", GoGetter: "Domains"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "newBucket", GoMethod: "NewBucket"},
			_jsii_.MemberMethod{JsiiMethod: "newDistribution", GoMethod: "NewDistribution"},
			_jsii_.MemberMethod{JsiiMethod: "newFunction", GoMethod: "NewFunction"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGatewayOrigin", GoMethod: "NewHttpApiGatewayOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGw", GoMethod: "NewHttpApiGw"},
			_jsii_.MemberMethod{JsiiMethod: "newS3Origin", GoMethod: "NewS3Origin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseRoutesFromDir", GoMethod: "ParseRoutesFromDir"},
			_jsii_.MemberMethod{JsiiMethod: "strToRuntime", GoMethod: "StrToRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EdgeAstroSite{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AstroSiteConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.EdgeAstroSiteProps",
		reflect.TypeOf((*EdgeAstroSiteProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.GwOptions",
		reflect.TypeOf((*GwOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@astrojs-aws/construct.LambdaAstroSite",
		reflect.TypeOf((*LambdaAstroSite)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "domains", GoGetter: "Domains"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "newBucket", GoMethod: "NewBucket"},
			_jsii_.MemberMethod{JsiiMethod: "newDistribution", GoMethod: "NewDistribution"},
			_jsii_.MemberMethod{JsiiMethod: "newFunction", GoMethod: "NewFunction"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGatewayOrigin", GoMethod: "NewHttpApiGatewayOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGw", GoMethod: "NewHttpApiGw"},
			_jsii_.MemberMethod{JsiiMethod: "newS3Origin", GoMethod: "NewS3Origin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseRoutesFromDir", GoMethod: "ParseRoutesFromDir"},
			_jsii_.MemberMethod{JsiiMethod: "strToRuntime", GoMethod: "StrToRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaAstroSite{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AstroSiteConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.LambdaAstroSiteProps",
		reflect.TypeOf((*LambdaAstroSiteProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.ServerOptions",
		reflect.TypeOf((*ServerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@astrojs-aws/construct.StaticAstroSite",
		reflect.TypeOf((*StaticAstroSite)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "domains", GoGetter: "Domains"},
			_jsii_.MemberMethod{JsiiMethod: "newBucket", GoMethod: "NewBucket"},
			_jsii_.MemberMethod{JsiiMethod: "newDistribution", GoMethod: "NewDistribution"},
			_jsii_.MemberMethod{JsiiMethod: "newFunction", GoMethod: "NewFunction"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGatewayOrigin", GoMethod: "NewHttpApiGatewayOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "newHttpApiGw", GoMethod: "NewHttpApiGw"},
			_jsii_.MemberMethod{JsiiMethod: "newS3Origin", GoMethod: "NewS3Origin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseRoutesFromDir", GoMethod: "ParseRoutesFromDir"},
			_jsii_.MemberMethod{JsiiMethod: "strToRuntime", GoMethod: "StrToRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StaticAstroSite{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AstroSiteConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@astrojs-aws/construct.StaticAstroSiteProps",
		reflect.TypeOf((*StaticAstroSiteProps)(nil)).Elem(),
	)
}
