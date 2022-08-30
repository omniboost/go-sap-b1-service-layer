package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewBusinessPartnersPostRequest() BusinessPartnersPostRequest {
	r := BusinessPartnersPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
		path:    "/b1s/v1/BusinessPartners",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BusinessPartnersPostRequest struct {
	client      *Client
	queryParams *BusinessPartnersPostRequestQueryParams
	path        string
	pathParams  *BusinessPartnersPostRequestPathParams
	method      string
	headers     http.Header
	requestBody BusinessPartnersPostRequestBody
}

func (r BusinessPartnersPostRequest) NewQueryParams() *BusinessPartnersPostRequestQueryParams {
	return &BusinessPartnersPostRequestQueryParams{}
}

type BusinessPartnersPostRequestQueryParams struct {
}

func (p BusinessPartnersPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BusinessPartnersPostRequest) QueryParams() *BusinessPartnersPostRequestQueryParams {
	return r.queryParams
}

func (r BusinessPartnersPostRequest) NewPathParams() *BusinessPartnersPostRequestPathParams {
	return &BusinessPartnersPostRequestPathParams{}
}

type BusinessPartnersPostRequestPathParams struct {
}

func (p *BusinessPartnersPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BusinessPartnersPostRequest) Path() *string {
	return &r.path
}

func (r *BusinessPartnersPostRequest) PathParams() *BusinessPartnersPostRequestPathParams {
	return r.pathParams
}

func (r *BusinessPartnersPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BusinessPartnersPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *BusinessPartnersPostRequest) Method() string {
	return r.method
}

func (r BusinessPartnersPostRequest) NewRequestBody() BusinessPartnersPostRequestBody {
	return BusinessPartnersPostRequestBody{}
}

type BusinessPartnersPostRequestBody BusinessPartner

func (r *BusinessPartnersPostRequest) RequestBody() *BusinessPartnersPostRequestBody {
	return &r.requestBody
}

func (r *BusinessPartnersPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BusinessPartnersPostRequest) SetRequestBody(body BusinessPartnersPostRequestBody) {
	r.requestBody = body
}

func (r *BusinessPartnersPostRequest) NewResponseBody() *BusinessPartnersPostResponseBody {
	return &BusinessPartnersPostResponseBody{}
}

type BusinessPartnersPostResponseBody BusinessPartner

func (r *BusinessPartnersPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *BusinessPartnersPostRequest) Do() (BusinessPartnersPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = r.client.InitSession(req)
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
