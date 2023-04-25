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
type StaticAstroSite interface {
	AstroSiteConstruct
	// Experimental.
	BucketArn() *string
	// Experimental.
	BucketName() *string
	// Experimental.
	DistributionId() *string
	// Experimental.
	Domains() *[]*string
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

// The jsii proxy struct for StaticAstroSite
type jsiiProxy_StaticAstroSite struct {
	jsiiProxy_AstroSiteConstruct
}

func (j *jsiiProxy_StaticAstroSite) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticAstroSite) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticAstroSite) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticAstroSite) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticAstroSite) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewStaticAstroSite(scope constructs.Construct, id *string, props *StaticAstroSiteProps) StaticAstroSite {
	_init_.Initialize()

	if err := validateNewStaticAstroSiteParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StaticAstroSite{}

	_jsii_.Create(
		"@astrojs-aws/construct.StaticAstroSite",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStaticAstroSite_Override(s StaticAstroSite, scope constructs.Construct, id *string, props *StaticAstroSiteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@astrojs-aws/construct.StaticAstroSite",
		[]interface{}{scope, id, props},
		s,
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
func StaticAstroSite_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStaticAstroSite_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@astrojs-aws/construct.StaticAstroSite",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewBucket(scope constructs.Construct, whEnabled *bool, props *AssetsOptions) awss3.Bucket {
	if err := s.validateNewBucketParameters(scope, whEnabled, props); err != nil {
		panic(err)
	}
	var returns awss3.Bucket

	_jsii_.Invoke(
		s,
		"newBucket",
		[]interface{}{scope, whEnabled, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewDistribution(scope constructs.Construct, defaultBehavior *awscloudfront.BehaviorOptions, props *CfOptions) awscloudfront.Distribution {
	if err := s.validateNewDistributionParameters(scope, defaultBehavior, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.Distribution

	_jsii_.Invoke(
		s,
		"newDistribution",
		[]interface{}{scope, defaultBehavior, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewFunction(scope constructs.Construct, serverEntry *string, props *ServerOptions) awslambdanodejs.NodejsFunction {
	if err := s.validateNewFunctionParameters(scope, serverEntry, props); err != nil {
		panic(err)
	}
	var returns awslambdanodejs.NodejsFunction

	_jsii_.Invoke(
		s,
		"newFunction",
		[]interface{}{scope, serverEntry, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewHttpApiGatewayOrigin(httpApi awscdkapigatewayv2alpha.HttpApi) awscloudfrontorigins.HttpOrigin {
	if err := s.validateNewHttpApiGatewayOriginParameters(httpApi); err != nil {
		panic(err)
	}
	var returns awscloudfrontorigins.HttpOrigin

	_jsii_.Invoke(
		s,
		"newHttpApiGatewayOrigin",
		[]interface{}{httpApi},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewHttpApiGw(scope constructs.Construct, fn awslambdanodejs.NodejsFunction, props *GwOptions) awscdkapigatewayv2alpha.HttpApi {
	if err := s.validateNewHttpApiGwParameters(scope, fn, props); err != nil {
		panic(err)
	}
	var returns awscdkapigatewayv2alpha.HttpApi

	_jsii_.Invoke(
		s,
		"newHttpApiGw",
		[]interface{}{scope, fn, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) NewS3Origin(scope constructs.Construct, bucket awss3.Bucket) awscloudfrontorigins.S3Origin {
	if err := s.validateNewS3OriginParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns awscloudfrontorigins.S3Origin

	_jsii_.Invoke(
		s,
		"newS3Origin",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) ParseRoutesFromDir(dir *string, isCf *bool) *map[string]*string {
	if err := s.validateParseRoutesFromDirParameters(dir); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"parseRoutesFromDir",
		[]interface{}{dir, isCf},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) StrToRuntime(str *string) awslambda.Runtime {
	var returns awslambda.Runtime

	_jsii_.Invoke(
		s,
		"strToRuntime",
		[]interface{}{str},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StaticAstroSite) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

