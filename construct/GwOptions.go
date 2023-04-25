// The CDK Construct Library of Astro
package construct

import (
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

// The options for the CloudFront distribution.
// Experimental.
type GwOptions struct {
	// OIDC scopes attached to the gateway.
	// Experimental.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Authorizer to applied to the gateway.
	// Experimental.
	Authorizer awscdkapigatewayv2alpha.IHttpRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Specifies a CORS configuration for an API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
	//
	// Experimental.
	Cors *awscdkapigatewayv2alpha.CorsPreflightOptions `field:"optional" json:"cors" yaml:"cors"`
}

