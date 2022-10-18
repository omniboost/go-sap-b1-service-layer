package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewMetadataGetRequest() MetadataGetRequest {
	r := MetadataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
		path:    "/b1s/v1/$metadata",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataGetRequest struct {
	client      *Client
	queryParams *MetadataGetRequestQueryParams
	path        string
	pathParams  *MetadataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MetadataGetRequestBody
}

func (r MetadataGetRequest) NewQueryParams() *MetadataGetRequestQueryParams {
	return &MetadataGetRequestQueryParams{}
}

type MetadataGetRequestQueryParams struct {
}

func (p MetadataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataGetRequest) QueryParams() *MetadataGetRequestQueryParams {
	return r.queryParams
}

func (r MetadataGetRequest) NewPathParams() *MetadataGetRequestPathParams {
	return &MetadataGetRequestPathParams{}
}

type MetadataGetRequestPathParams struct {
}

func (p *MetadataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MetadataGetRequest) Path() *string {
	return &r.path
}

func (r *MetadataGetRequest) PathParams() *MetadataGetRequestPathParams {
	return r.pathParams
}

func (r *MetadataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetadataGetRequest) Method() string {
	return r.method
}

func (r MetadataGetRequest) NewRequestBody() MetadataGetRequestBody {
	return MetadataGetRequestBody{}
}

type MetadataGetRequestBody struct{}

func (r *MetadataGetRequest) RequestBody() *MetadataGetRequestBody {
	return &r.requestBody
}

func (r *MetadataGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *MetadataGetRequest) SetRequestBody(body MetadataGetRequestBody) {
	r.requestBody = body
}

func (r *MetadataGetRequest) NewResponseBody() *MetadataGetResponseBody {
	return &MetadataGetResponseBody{}
}

type MetadataGetResponseBody struct{}

func (r *MetadataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *MetadataGetRequest) Do() (MetadataGetResponseBody, error) {
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
