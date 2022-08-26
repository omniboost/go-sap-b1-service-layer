package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewDiscountsGetRequest() DiscountsGetRequest {
	r := DiscountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DiscountsGetRequest struct {
	client      *Client
	queryParams *DiscountsGetRequestQueryParams
	pathParams  *DiscountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody DiscountsGetRequestBody
}

func (r DiscountsGetRequest) NewQueryParams() *DiscountsGetRequestQueryParams {
	return &DiscountsGetRequestQueryParams{}
}

type DiscountsGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p DiscountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DiscountsGetRequest) QueryParams() *DiscountsGetRequestQueryParams {
	return r.queryParams
}

func (r DiscountsGetRequest) NewPathParams() *DiscountsGetRequestPathParams {
	return &DiscountsGetRequestPathParams{}
}

type DiscountsGetRequestPathParams struct {
}

func (p *DiscountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DiscountsGetRequest) PathParams() *DiscountsGetRequestPathParams {
	return r.pathParams
}

func (r *DiscountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DiscountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DiscountsGetRequest) Method() string {
	return r.method
}

func (r DiscountsGetRequest) NewRequestBody() DiscountsGetRequestBody {
	return DiscountsGetRequestBody{}
}

type DiscountsGetRequestBody struct {
}

func (r *DiscountsGetRequest) RequestBody() *DiscountsGetRequestBody {
	return nil
}

func (r *DiscountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DiscountsGetRequest) SetRequestBody(body DiscountsGetRequestBody) {
	r.requestBody = body
}

func (r *DiscountsGetRequest) NewResponseBody() *DiscountsGetResponseBody {
	return &DiscountsGetResponseBody{}
}

type DiscountsGetResponseBody Discounts

func (r *DiscountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/discounts", r.PathParams())
	return &u
}

func (r *DiscountsGetRequest) Do() (DiscountsGetResponseBody, error, *http.Response) {
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
