package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewRestaurantGetRequest() RestaurantGetRequest {
	r := RestaurantGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RestaurantGetRequest struct {
	client      *Client
	queryParams *RestaurantGetRequestQueryParams
	pathParams  *RestaurantGetRequestPathParams
	method      string
	headers     http.Header
	requestBody RestaurantGetRequestBody
}

func (r RestaurantGetRequest) NewQueryParams() *RestaurantGetRequestQueryParams {
	return &RestaurantGetRequestQueryParams{}
}

type RestaurantGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p RestaurantGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RestaurantGetRequest) QueryParams() *RestaurantGetRequestQueryParams {
	return r.queryParams
}

func (r RestaurantGetRequest) NewPathParams() *RestaurantGetRequestPathParams {
	return &RestaurantGetRequestPathParams{}
}

type RestaurantGetRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *RestaurantGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *RestaurantGetRequest) PathParams() *RestaurantGetRequestPathParams {
	return r.pathParams
}

func (r *RestaurantGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RestaurantGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *RestaurantGetRequest) Method() string {
	return r.method
}

func (r RestaurantGetRequest) NewRequestBody() RestaurantGetRequestBody {
	return RestaurantGetRequestBody{}
}

type RestaurantGetRequestBody struct {
}

func (r *RestaurantGetRequest) RequestBody() *RestaurantGetRequestBody {
	return nil
}

func (r *RestaurantGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *RestaurantGetRequest) SetRequestBody(body RestaurantGetRequestBody) {
	r.requestBody = body
}

func (r *RestaurantGetRequest) NewResponseBody() *RestaurantGetResponseBody {
	return &RestaurantGetResponseBody{}
}

type RestaurantGetResponseBody Restaurant

func (r *RestaurantGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/restaurants/v1/restaurants/{{.GUID}}", r.PathParams())
	return &u
}

func (r *RestaurantGetRequest) Do() (RestaurantGetResponseBody, error, *http.Response) {
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
