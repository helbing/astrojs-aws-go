// The CDK Construct Library of Astro
package construct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/helbing/astrojs-aws-go/construct/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type EdgeAstroSite interface {
	AstroSiteConstruct
	// Experimental.
	BucketArn() *string
	// Experimental.
	BucketName() *string
	// Experimental.
	DistributionId() *string
	// Experimental.
	Domains() *[]*string
	// Experimental.
	FunctionArn() *string
	// Experimental.
	FunctionName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// New bucket.
	// Experimental.
	NewBucket(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) awss3.Bucket
	// New CloudFront distribution.
	// Experimental.
	NewDistribution(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) awscloudfront.Distribution
	// New nodejs function.
	// Experimental.
	NewFunction(scope constructs.Construct, serverEntry *string, props *ServerOptions) awslambdanodejs.NodejsFunction
	// New HttpApi Gateway origin.
	// Experimental.
	NewHttpApiGatewayOrigin(httpApi awscdkapigatewayv2alpha.HttpApi) awscloudfrontorigins.HttpOrigin
	// New HttpApi Gateway.
	// Experimental.
	NewHttpApiGw(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) awscdkapigatewayv2alpha.HttpApi
	// New S3 origin.
	// Experimental.
	NewS3Origin(scope constructs.Construct, bucket awss3.Bucket) awscloudfrontorigins.S3Origin
	// Parse routes from directory.
	//
	// if the item is directory will parse to {"/item/*": "/item/*"} or {"/item/{proxy+}": "/item/{proxy}"}
	// if the item is file will parse to {"/item": "/item"}.
	// Experimental.
	ParseRoutesFromDir(dir *string, isCf *bool) *map[string]*string
	// Transform string to Runtime.
	// Experimental.
	StrToRuntime(str *string) awslambda.Runtime
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgeAstroSite
type jsiiProxy_EdgeAstroSite struct {
	jsiiProxy_AstroSiteConstruct
}

func (j *jsiiProxy_EdgeAstroSite) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeAstroSite) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewEdgeAstroSite(scope constructs.Construct, id *string, props *EdgeAstroSiteProps) EdgeAstroSite {
	_init_.Initialize()

	if err := validateNewEdgeAstroSiteParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgeAstroSite{}

	_jsii_.Create(
		"@astrojs-aws/construct.EdgeAstroSite",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEdgeAstroSite_Override(e EdgeAstroSite, scope constructs.Construct, id *string, props *EdgeAstroSiteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@astrojs-aws/construct.EdgeAstroSite",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func EdgeAstroSite_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgeAstroSite_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@astrojs-aws/construct.EdgeAstroSite",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewBucket(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) awss3.Bucket {
	if err := e.validateNewBucketParameters(scope, whEnabled, props); err != nil {
		panic(err)
	}
	var returns awss3.Bucket

	_jsii_.Invoke(
		e,
		"newBucket",
		[]interface{}{scope, whEnabled, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewDistribution(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) awscloudfront.Distribution {
	if err := e.validateNewDistributionParameters(scope, defaultBehavior, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.Distribution

	_jsii_.Invoke(
		e,
		"newDistribution",
		[]interface{}{scope, defaultBehavior, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewFunction(scope constructs.Construct, serverEntry *string, props *ServerOptions) awslambdanodejs.NodejsFunction {
	if err := e.validateNewFunctionParameters(scope, serverEntry, props); err != nil {
		panic(err)
	}
	var returns awslambdanodejs.NodejsFunction

	_jsii_.Invoke(
		e,
		"newFunction",
		[]interface{}{scope, serverEntry, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewHttpApiGatewayOrigin(httpApi awscdkapigatewayv2alpha.HttpApi) awscloudfrontorigins.HttpOrigin {
	if err := e.validateNewHttpApiGatewayOriginParameters(httpApi); err != nil {
		panic(err)
	}
	var returns awscloudfrontorigins.HttpOrigin

	_jsii_.Invoke(
		e,
		"newHttpApiGatewayOrigin",
		[]interface{}{httpApi},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewHttpApiGw(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) awscdkapigatewayv2alpha.HttpApi {
	if err := e.validateNewHttpApiGwParameters(scope, fn, props); err != nil {
		panic(err)
	}
	var returns awscdkapigatewayv2alpha.HttpApi

	_jsii_.Invoke(
		e,
		"newHttpApiGw",
		[]interface{}{scope, fn, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) NewS3Origin(scope constructs.Construct, bucket awss3.Bucket) awscloudfrontorigins.S3Origin {
	if err := e.validateNewS3OriginParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns awscloudfrontorigins.S3Origin

	_jsii_.Invoke(
		e,
		"newS3Origin",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) ParseRoutesFromDir(dir *string, isCf *bool) *map[string]*string {
	if err := e.validateParseRoutesFromDirParameters(dir); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"parseRoutesFromDir",
		[]interface{}{dir, isCf},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) StrToRuntime(str *string) awslambda.Runtime {
	var returns awslambda.Runtime

	_jsii_.Invoke(
		e,
		"strToRuntime",
		[]interface{}{str},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeAstroSite) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

