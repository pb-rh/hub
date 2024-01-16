// Code generated by goa v3.14.6, DO NOT EDIT.
//
// admin service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Admin service
type Service interface {
	// Create or Update an agent user with required scopes
	UpdateAgent(context.Context, *UpdateAgentPayload) (res *UpdateAgentResult, err error)
	// Refresh the changes in config file
	RefreshConfig(context.Context, *RefreshConfigPayload) (res *RefreshConfigResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "hub"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "admin"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"UpdateAgent", "RefreshConfig"}

// RefreshConfigPayload is the payload type of the admin service RefreshConfig
// method.
type RefreshConfigPayload struct {
	// User JWT
	Token string
}

// RefreshConfigResult is the result type of the admin service RefreshConfig
// method.
type RefreshConfigResult struct {
	// Config file checksum
	Checksum string
}

// UpdateAgentPayload is the payload type of the admin service UpdateAgent
// method.
type UpdateAgentPayload struct {
	// User JWT
	Token string
	// Name of Agent
	Name string
	// Scopes required for Agent
	Scopes []string
}

// UpdateAgentResult is the result type of the admin service UpdateAgent method.
type UpdateAgentResult struct {
	// Agent JWT
	Token string
}

// MakeInvalidPayload builds a goa.ServiceError from an error.
func MakeInvalidPayload(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid-payload", false, false, false)
}

// MakeInvalidToken builds a goa.ServiceError from an error.
func MakeInvalidToken(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid-token", false, false, false)
}

// MakeInvalidScopes builds a goa.ServiceError from an error.
func MakeInvalidScopes(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid-scopes", false, false, false)
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal-error", false, false, false)
}
