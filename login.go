package sap

import (
	"context"
	"github.com/omniboost/go-omniboost-http-client/client"
	"net/http"
)

type (
	LoginRequest struct {
		client *SapB1Client
		body   LoginRequestBody
	}
	LoginRequestBody struct {
		CompanyDB string `json:"CompanyDB"`
		UserName  string `json:"UserName"`
		Password  string `json:"Password"`
	}

	LoginRequestOpts func(r *LoginRequest)
)

var _ client.RequestWithBody = (*LoginRequest)(nil)
var _ client.RequestWithAuthPreference = (*LoginRequest)(nil)

func (r *LoginRequest) Method() string {
	return http.MethodPost
}

func (r *LoginRequest) PathTemplate() string {
	return "/b1s/v2/Login"
}

func (r *LoginRequest) Body() any {
	return r.body
}

func (r *LoginRequest) SkipAuth() bool {
	return true
}

func LoginWithCredentials(username, password, companyDB string) LoginRequestOpts {
	return func(r *LoginRequest) {
		r.body.UserName = username
		r.body.Password = password
		r.body.CompanyDB = companyDB
	}
}

func (s *SapB1Client) NewLoginRequest(opts ...LoginRequestOpts) *LoginRequest {
	req := &LoginRequest{
		client: s,
		body:   LoginRequestBody{},
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

func (r *LoginRequest) Do(ctx context.Context) error {
	return r.client.Do(ctx, r, nil)
}
