package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewRevenueCentersGetRequest() RevenueCentersGetRequest {
	r := RevenueCentersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RevenueCentersGetRequest struct {
	client      *Client
	queryParams *RevenueCentersGetRequestQueryParams
	pathParams  *RevenueCentersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody RevenueCentersGetRequestBody
}

func (r RevenueCentersGetRequest) NewQueryParams() *RevenueCentersGetRequestQueryParams {
	return &RevenueCentersGetRequestQueryParams{}
}

type RevenueCentersGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p RevenueCentersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RevenueCentersGetRequest) QueryParams() *RevenueCentersGetRequestQueryParams {
	return r.queryParams
}

func (r RevenueCentersGetRequest) NewPathParams() *RevenueCentersGetRequestPathParams {
	return &RevenueCentersGetRequestPathParams{}
}

type RevenueCentersGetRequestPathParams struct {
}

func (p *RevenueCentersGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RevenueCentersGetRequest) PathParams() *RevenueCentersGetRequestPathParams {
	return r.pathParams
}

func (r *RevenueCentersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RevenueCentersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *RevenueCentersGetRequest) Method() string {
	return r.method
}

func (r RevenueCentersGetRequest) NewRequestBody() RevenueCentersGetRequestBody {
	return RevenueCentersGetRequestBody{}
}

type RevenueCentersGetRequestBody struct {
}

func (r *RevenueCentersGetRequest) RequestBody() *RevenueCentersGetRequestBody {
	return nil
}

func (r *RevenueCentersGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *RevenueCentersGetRequest) SetRequestBody(body RevenueCentersGetRequestBody) {
	r.requestBody = body
}

func (r *RevenueCentersGetRequest) NewResponseBody() *RevenueCentersGetResponseBody {
	return &RevenueCentersGetResponseBody{}
}

type RevenueCentersGetResponseBody RevenueCenters

func (r *RevenueCentersGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/revenueCenters/", r.PathParams())
	return &u
}

func (r *RevenueCentersGetRequest) Do() (RevenueCentersGetResponseBody, error, *http.Response) {
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

func (r *RevenueCentersGetRequest) All() (RevenueCentersGetResponseBody, error) {
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
