// The CDK Construct Library of Astro
package astrojsawsconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/helbing/astrojs-aws-go/construct/astrojsawsconstruct/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/helbing/astrojs-aws-go/construct/astrojsawsconstruct/internal"
)

// The base class for all constructs.
// Experimental.
type AstroSiteConstruct interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// New bucket.
	// Experimental.
	NewBucket(scope constructs.Construct, staticDir *string, props *awss3.BucketProps) awss3.Bucket
	// New nodejs function.
	// Experimental.
	NewFunction(scope constructs.Construct, serverEntry *string, props *ServerOptions) awslambdanodejs.NodejsFunction
	// New S3 origin.
	// Experimental.
	NewS3Origin(scope constructs.Construct, bucket awss3.Bucket) awscloudfrontorigins.S3Origin
	// Parse routes from directory if the item is directory will parse to /item/* if the item is file will parse to /item.
	// Experimental.
	ParseRoutesFromDir(dir *string) *[]*string
	// Transform string to Runtime.
	// Experimental.
	StrToRuntime(str *string) awslambda.Runtime
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AstroSiteConstruct
type jsiiProxy_AstroSiteConstruct struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AstroSiteConstruct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAstroSiteConstruct(scope constructs.Construct, id *string) AstroSiteConstruct {
	_init_.Initialize()

	if err := validateNewAstroSiteConstructParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_AstroSiteConstruct{}

	_jsii_.Create(
		"@astrojs-aws/construct.AstroSiteConstruct",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewAstroSiteConstruct_Override(a AstroSiteConstruct, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@astrojs-aws/construct.AstroSiteConstruct",
		[]interface{}{scope, id},
		a,
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
func AstroSiteConstruct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAstroSiteConstruct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@astrojs-aws/construct.AstroSiteConstruct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) NewBucket(scope constructs.Construct, staticDir *string, props *awss3.BucketProps) awss3.Bucket {
	if err := a.validateNewBucketParameters(scope, staticDir, props); err != nil {
		panic(err)
	}
	var returns awss3.Bucket

	_jsii_.Invoke(
		a,
		"newBucket",
		[]interface{}{scope, staticDir, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) NewFunction(scope constructs.Construct, serverEntry *string, props *ServerOptions) awslambdanodejs.NodejsFunction {
	if err := a.validateNewFunctionParameters(scope, serverEntry, props); err != nil {
		panic(err)
	}
	var returns awslambdanodejs.NodejsFunction

	_jsii_.Invoke(
		a,
		"newFunction",
		[]interface{}{scope, serverEntry, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) NewS3Origin(scope constructs.Construct, bucket awss3.Bucket) awscloudfrontorigins.S3Origin {
	if err := a.validateNewS3OriginParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns awscloudfrontorigins.S3Origin

	_jsii_.Invoke(
		a,
		"newS3Origin",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) ParseRoutesFromDir(dir *string) *[]*string {
	if err := a.validateParseRoutesFromDirParameters(dir); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"parseRoutesFromDir",
		[]interface{}{dir},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) StrToRuntime(str *string) awslambda.Runtime {
	var returns awslambda.Runtime

	_jsii_.Invoke(
		a,
		"strToRuntime",
		[]interface{}{str},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AstroSiteConstruct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

