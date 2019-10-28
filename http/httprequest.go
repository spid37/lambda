package http

import (
	"context"
	"encoding/json"
)

// data available in request
// https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html

// EventHandler handler type for a http event lambda
type EventHandler func(context.Context, Request) (interface{}, error)

// Request -
type Request struct {
	ResourcePath   string           `json:"resourcePath"`
	HTTPMethod     string           `json:"httpMethod"`
	Identity       *RequestIdentity `json:"identity"`
	Params         json.RawMessage  `json:"params"`
	Query          json.RawMessage  `json:"query"`
	Body           json.RawMessage  `json:"body"`
	AuthorizerData *AuthorizerData  `json:"authorizer"`
}

// RequestIdentity -
type RequestIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId"`
	AccountID                     string `json:"accountId"`
	CognitoIdentityID             string `json:"cognitoIdentityId"`
	Caller                        string `json:"caller"`
	APIKey                        string `json:"apiKey"`
	SourceIP                      string `json:"sourceIp"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider"`
	UserArn                       string `json:"userArn"`
	UserAgent                     string `json:"userAgent"`
	User                          string `json:"user"`
}

// AuthorizerData is additional data stored on the response
type AuthorizerData struct {
	UserID string `json:"userId,omitempty"`
	Role   string `json:"role,omitempty"`
}
