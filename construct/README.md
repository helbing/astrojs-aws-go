# Astro Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## Introduction

[Astro](https://astro.build/) is the all-in-one web framework designed for speed. This library is supported modes, `static`, `lambda` and `edge`, to deploy astro website to AWS.

## Usage

### Static Hosting

![Static Hosting](https://raw.githubusercontent.com/helbing/astrojs-aws/main/docs/images/static-hosting.png)

For developing, it's not recommended to deploy with CloudFront, because setup CloudFront is really too slow, at least takes 5min.

```go
import "github.com/helbing/astrojs-aws-go/construct"


site := construct.NewStaticAstroSite(this, jsii.String("Site"), &StaticAstroSiteProps{
	SiteDir: jsii.String("/path/to/dist"),
})

awscdk.NewCfnOutput(this, jsii.String("Domains"), &CfnOutputProps{
	Value: site.Domains.join(jsii.String(", ")),
})
```

It's recommended to deploy with CloudFront in production environment.

```go
import "github.com/helbing/astrojs-aws-go/construct"


site := construct.NewStaticAstroSite(this, jsii.String("Site"), &StaticAstroSiteProps{
	SiteDir: jsii.String("/path/to/dist"),
	CfOptions: &CfOptions{
		Domain: jsii.String("example.com"),
		CertificateArn: jsii.String("arn:aws:acm:us-east-1:xxx-xxx-xxx"),
	},
})

awscdk.NewCfnOutput(this, jsii.String("Domains"), &CfnOutputProps{
	Value: site.Domains.join(jsii.String(", ")),
})
```

### Lambda Hosting

![Lambda Hosting](https://raw.githubusercontent.com/helbing/astrojs-aws/main/docs/images/lambda-hosting.png)

```go
import "github.com/helbing/astrojs-aws-go/construct"


site := construct.NewLambdaAstroSite(this, jsii.String("Site"), &LambdaAstroSiteProps{
	ServerEntry: jsii.String("/path/to/server/entry.mjs"),
	StaticDir: jsii.String("/path/to/client"),
})

awscdk.NewCfnOutput(this, jsii.String("Domains"), &CfnOutputProps{
	Value: site.Domains.join(jsii.String(", ")),
})
```

Deploy with CloudFront.

```go
import "github.com/helbing/astrojs-aws-go/construct"


site := construct.NewLambdaAstroSite(this, jsii.String("Site"), &LambdaAstroSiteProps{
	ServerEntry: jsii.String("/path/to/server/entry.mjs"),
	StaticDir: jsii.String("/path/to/client"),
	CfOptions: &CfOptions{
		Domain: jsii.String("example.com"),
		CertificateArn: jsii.String("arn:aws:acm:us-east-1:xxx-xxx-xxx"),
	},
})

awscdk.NewCfnOutput(this, jsii.String("Domains"), &CfnOutputProps{
	Value: site.Domains.join(jsii.String(", ")),
})
```

### Edge Hosting

![Edge Hosting](https://raw.githubusercontent.com/helbing/astrojs-aws/main/docs/images/edge-hosting.png)

As we known that, edge function working in Edge node. But for developing, setup CloudFront is really too slow. We can only deploy the Lambda function, and use [AWS SAM](https://aws.amazon.com/serverless/sam/) for testing and debuging.

```go
import "github.com/helbing/astrojs-aws-go/construct"


construct.NewEdgeAstroSite(this, jsii.String("Site"), &EdgeAstroSiteProps{
	ServerEntry: jsii.String("/path/to/server/entry.mjs"),
	StaticDir: jsii.String("/path/to/client"),
	OnlyLambda: jsii.Boolean(true),
})
```

Deploy to production environment.

```go
import "github.com/helbing/astrojs-aws-go/construct"


site := construct.NewEdgeAstroSite(this, jsii.String("Site"), &EdgeAstroSiteProps{
	ServerEntry: jsii.String("/path/to/server/entry.mjs"),
	StaticDir: jsii.String("/path/to/client"),
	CfOptions: &CfOptions{
		Domain: jsii.String("example.com"),
		CertificateArn: jsii.String("arn:aws:acm:us-east-1:xxx-xxx-xxx"),
	},
})

awscdk.NewCfnOutput(this, jsii.String("Domains"), &CfnOutputProps{
	Value: site.Domains.join(jsii.String(", ")),
})
```
