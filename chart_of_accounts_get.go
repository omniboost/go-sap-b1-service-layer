package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewChartOfAccountsGetRequest() ChartOfAccountsGetRequest {
	r := ChartOfAccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
		path:    "/b1s/v1/ChartOfAccounts",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ChartOfAccountsGetRequest struct {
	client      *Client
	queryParams *ChartOfAccountsGetRequestQueryParams
	path        string
	pathParams  *ChartOfAccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ChartOfAccountsGetRequestBody
}

func (r ChartOfAccountsGetRequest) NewQueryParams() *ChartOfAccountsGetRequestQueryParams {
	return &ChartOfAccountsGetRequestQueryParams{}
}

type ChartOfAccountsGetRequestQueryParams struct {
}

func (p ChartOfAccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ChartOfAccountsGetRequest) QueryParams() *ChartOfAccountsGetRequestQueryParams {
	return r.queryParams
}

func (r ChartOfAccountsGetRequest) NewPathParams() *ChartOfAccountsGetRequestPathParams {
	return &ChartOfAccountsGetRequestPathParams{}
}

type ChartOfAccountsGetRequestPathParams struct {
}

func (p *ChartOfAccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ChartOfAccountsGetRequest) Path() *string {
	return &r.path
}

func (r *ChartOfAccountsGetRequest) PathParams() *ChartOfAccountsGetRequestPathParams {
	return r.pathParams
}

func (r *ChartOfAccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ChartOfAccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ChartOfAccountsGetRequest) Method() string {
	return r.method
}

func (r ChartOfAccountsGetRequest) NewRequestBody() ChartOfAccountsGetRequestBody {
	return ChartOfAccountsGetRequestBody{}
}

type ChartOfAccountsGetRequestBody struct{}

func (r *ChartOfAccountsGetRequest) RequestBody() *ChartOfAccountsGetRequestBody {
	return &r.requestBody
}

func (r *ChartOfAccountsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ChartOfAccountsGetRequest) SetRequestBody(body ChartOfAccountsGetRequestBody) {
	r.requestBody = body
}

func (r *ChartOfAccountsGetRequest) NewResponseBody() *ChartOfAccountsGetResponseBody {
	return &ChartOfAccountsGetResponseBody{}
}

type ChartOfAccountsGetResponseBody struct{}

func (r *ChartOfAccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *ChartOfAccountsGetRequest) Do() (ChartOfAccountsGetResponseBody, error) {
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
