package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewDiningOptionsGetRequest() DiningOptionsGetRequest {
	r := DiningOptionsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DiningOptionsGetRequest struct {
	client      *Client
	queryParams *DiningOptionsGetRequestQueryParams
	pathParams  *DiningOptionsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody DiningOptionsGetRequestBody
}

func (r DiningOptionsGetRequest) NewQueryParams() *DiningOptionsGetRequestQueryParams {
	return &DiningOptionsGetRequestQueryParams{}
}

type DiningOptionsGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p DiningOptionsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DiningOptionsGetRequest) QueryParams() *DiningOptionsGetRequestQueryParams {
	return r.queryParams
}

func (r DiningOptionsGetRequest) NewPathParams() *DiningOptionsGetRequestPathParams {
	return &DiningOptionsGetRequestPathParams{}
}

type DiningOptionsGetRequestPathParams struct {
}

func (p *DiningOptionsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DiningOptionsGetRequest) PathParams() *DiningOptionsGetRequestPathParams {
	return r.pathParams
}

func (r *DiningOptionsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DiningOptionsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DiningOptionsGetRequest) Method() string {
	return r.method
}

func (r DiningOptionsGetRequest) NewRequestBody() DiningOptionsGetRequestBody {
	return DiningOptionsGetRequestBody{}
}

type DiningOptionsGetRequestBody struct {
}

func (r *DiningOptionsGetRequest) RequestBody() *DiningOptionsGetRequestBody {
	return nil
}

func (r *DiningOptionsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DiningOptionsGetRequest) SetRequestBody(body DiningOptionsGetRequestBody) {
	r.requestBody = body
}

func (r *DiningOptionsGetRequest) NewResponseBody() *DiningOptionsGetResponseBody {
	return &DiningOptionsGetResponseBody{}
}

type DiningOptionsGetResponseBody DiningOptions

func (r *DiningOptionsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/diningOptions/", r.PathParams())
	return &u
}

func (r *DiningOptionsGetRequest) Do() (DiningOptionsGetResponseBody, error, *http.Response) {
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

func (r *DiningOptionsGetRequest) All() (DiningOptionsGetResponseBody, error) {
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
