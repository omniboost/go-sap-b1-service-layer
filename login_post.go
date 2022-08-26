package toast

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewLoginPostRequest() LoginPostRequest {
	r := LoginPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LoginPostRequest struct {
	client      *Client
	queryParams *LoginPostRequestQueryParams
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
	return LoginPostRequestBody{
		UserAccessType: "TOAST_MACHINE_CLIENT",
	}
}

type LoginPostRequestBody struct {
	ClientID       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	UserAccessType string `json:"userAccessType"`
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
	Class  string `json:"@class"`
	Token  Token  `json:"token"`
	Status string `json:"status"`
}

func (r *LoginPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/authentication/v1/authentication/login", r.PathParams())
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
	responseBody.Token.expiry = time.Now().Add(time.Duration(responseBody.Token.ExpiresIn) * time.Second)
	return *responseBody, err
}
