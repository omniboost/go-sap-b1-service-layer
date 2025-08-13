package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewWarehouseStockPostRequest() WarehouseStockPostRequest {
	r := WarehouseStockPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
		path:    "/b1s/v1/U_Z3_GM_ECO_WHSSTOCK",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type WarehouseStockPostRequest struct {
	client      *Client
	queryParams *WarehouseStockPostRequestQueryParams
	path        string
	pathParams  *WarehouseStockPostRequestPathParams
	method      string
	headers     http.Header
	requestBody WarehouseStockPostRequestBody
}

func (r WarehouseStockPostRequest) NewQueryParams() *WarehouseStockPostRequestQueryParams {
	return &WarehouseStockPostRequestQueryParams{}
}

type WarehouseStockPostRequestQueryParams struct {
}

func (p WarehouseStockPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *WarehouseStockPostRequest) QueryParams() *WarehouseStockPostRequestQueryParams {
	return r.queryParams
}

func (r WarehouseStockPostRequest) NewPathParams() *WarehouseStockPostRequestPathParams {
	return &WarehouseStockPostRequestPathParams{}
}

type WarehouseStockPostRequestPathParams struct {
}

func (p *WarehouseStockPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *WarehouseStockPostRequest) Path() *string {
	return &r.path
}

func (r *WarehouseStockPostRequest) PathParams() *WarehouseStockPostRequestPathParams {
	return r.pathParams
}

func (r *WarehouseStockPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *WarehouseStockPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *WarehouseStockPostRequest) Method() string {
	return r.method
}

func (r WarehouseStockPostRequest) NewRequestBody() WarehouseStockPostRequestBody {
	return WarehouseStockPostRequestBody{}
}

type WarehouseStockPostRequestBody WarehouseStock

func (r *WarehouseStockPostRequest) RequestBody() *WarehouseStockPostRequestBody {
	return &r.requestBody
}

func (r *WarehouseStockPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *WarehouseStockPostRequest) SetRequestBody(body WarehouseStockPostRequestBody) {
	r.requestBody = body
}

func (r *WarehouseStockPostRequest) NewResponseBody() *WarehouseStockPostResponseBody {
	return &WarehouseStockPostResponseBody{}
}

type WarehouseStockPostResponseBody WarehouseStock

func (r *WarehouseStockPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *WarehouseStockPostRequest) Do() (WarehouseStockPostResponseBody, error) {
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
