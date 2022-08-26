package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewSalesCategoriesGetRequest() SalesCategoriesGetRequest {
	r := SalesCategoriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesCategoriesGetRequest struct {
	client      *Client
	queryParams *SalesCategoriesGetRequestQueryParams
	pathParams  *SalesCategoriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesCategoriesGetRequestBody
}

func (r SalesCategoriesGetRequest) NewQueryParams() *SalesCategoriesGetRequestQueryParams {
	return &SalesCategoriesGetRequestQueryParams{}
}

type SalesCategoriesGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p SalesCategoriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesCategoriesGetRequest) QueryParams() *SalesCategoriesGetRequestQueryParams {
	return r.queryParams
}

func (r SalesCategoriesGetRequest) NewPathParams() *SalesCategoriesGetRequestPathParams {
	return &SalesCategoriesGetRequestPathParams{}
}

type SalesCategoriesGetRequestPathParams struct {
}

func (p *SalesCategoriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesCategoriesGetRequest) PathParams() *SalesCategoriesGetRequestPathParams {
	return r.pathParams
}

func (r *SalesCategoriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesCategoriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesCategoriesGetRequest) Method() string {
	return r.method
}

func (r SalesCategoriesGetRequest) NewRequestBody() SalesCategoriesGetRequestBody {
	return SalesCategoriesGetRequestBody{}
}

type SalesCategoriesGetRequestBody struct {
}

func (r *SalesCategoriesGetRequest) RequestBody() *SalesCategoriesGetRequestBody {
	return nil
}

func (r *SalesCategoriesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesCategoriesGetRequest) SetRequestBody(body SalesCategoriesGetRequestBody) {
	r.requestBody = body
}

func (r *SalesCategoriesGetRequest) NewResponseBody() *SalesCategoriesGetResponseBody {
	return &SalesCategoriesGetResponseBody{}
}

type SalesCategoriesGetResponseBody SalesCategories

func (r *SalesCategoriesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/salesCategories/", r.PathParams())
	return &u
}

func (r *SalesCategoriesGetRequest) Do() (SalesCategoriesGetResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	err = r.client.InitToken(req)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	return *responseBody, err, resp
}

func (r *SalesCategoriesGetRequest) All() (SalesCategoriesGetResponseBody, error) {
	body, err, resp := r.Do()
	if err != nil {
		return body, err
	}

	concat := body
	token, err := r.client.GetPageToken(resp)
	if err != nil {
		return concat, err
	}

	for token != "" {
		r.QueryParams().PageToken = token
		body, err, resp = r.Do()
		if err != nil {
			return concat, err
		}

		concat = append(concat, body...)
		token, err = r.client.GetPageToken(resp)
		if err != nil {
			return concat, err
		}
	}

	return concat, nil
}
