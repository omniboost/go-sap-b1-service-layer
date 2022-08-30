package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewLoginPostRequest() LoginPostRequest {
	r := LoginPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
		path:    "/b1s/v1/Login",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LoginPostRequest struct {
	client      *Client
	queryParams *LoginPostRequestQueryParams
	path        string
	pathParams  *LoginPostRequestPathParams
	method      string
	headers     http.Header
	requestBody LoginPostRequestBody
}

func (r LoginPostRequest) NewQueryParams() *LoginPostRequestQueryParams {
	return &LoginPostRequestQueryParams{}
}

type LoginPostRequestQueryParams struct {
}

func (p LoginPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LoginPostRequest) QueryParams() *LoginPostRequestQueryParams {
	return r.queryParams
}

func (r LoginPostRequest) NewPathParams() *LoginPostRequestPathParams {
	return &LoginPostRequestPathParams{}
}

type LoginPostRequestPathParams struct {
}

func (p *LoginPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LoginPostRequest) Path() *string {
	return &r.path
}

func (r *LoginPostRequest) PathParams() *LoginPostRequestPathParams {
	return r.pathParams
}

func (r *LoginPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LoginPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *LoginPostRequest) Method() string {
	return r.method
}

func (r LoginPostRequest) NewRequestBody() LoginPostRequestBody {
	return LoginPostRequestBody{}
}

type LoginPostRequestBody struct {
	Username  string `json:"UserName"`
	Password  string `json:"Password"`
	CompanyDB string `json:"CompanyDB"`
}

func (r *LoginPostRequest) RequestBody() *LoginPostRequestBody {
	return &r.requestBody
}

func (r *LoginPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *LoginPostRequest) SetRequestBody(body LoginPostRequestBody) {
	r.requestBody = body
}

func (r *LoginPostRequest) NewResponseBody() *LoginPostResponseBody {
	return &LoginPostResponseBody{}
}

type LoginPostResponseBody struct {
	OdataMetadata  string `json:"odata.metadata"`
	SessionID      string `json:"SessionId"`
	Version        string `json:"Version"`
	SessionTimeout int    `json:"SessionTimeout"`
}

func (r *LoginPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *LoginPostRequest) Do() (LoginPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
