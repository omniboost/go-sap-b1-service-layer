package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewAlternativePaymentTypesGetRequest() AlternativePaymentTypesGetRequest {
	r := AlternativePaymentTypesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AlternativePaymentTypesGetRequest struct {
	client      *Client
	queryParams *AlternativePaymentTypesGetRequestQueryParams
	pathParams  *AlternativePaymentTypesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AlternativePaymentTypesGetRequestBody
}

func (r AlternativePaymentTypesGetRequest) NewQueryParams() *AlternativePaymentTypesGetRequestQueryParams {
	return &AlternativePaymentTypesGetRequestQueryParams{}
}

type AlternativePaymentTypesGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p AlternativePaymentTypesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AlternativePaymentTypesGetRequest) QueryParams() *AlternativePaymentTypesGetRequestQueryParams {
	return r.queryParams
}

func (r AlternativePaymentTypesGetRequest) NewPathParams() *AlternativePaymentTypesGetRequestPathParams {
	return &AlternativePaymentTypesGetRequestPathParams{}
}

type AlternativePaymentTypesGetRequestPathParams struct {
}

func (p *AlternativePaymentTypesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AlternativePaymentTypesGetRequest) PathParams() *AlternativePaymentTypesGetRequestPathParams {
	return r.pathParams
}

func (r *AlternativePaymentTypesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AlternativePaymentTypesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AlternativePaymentTypesGetRequest) Method() string {
	return r.method
}

func (r AlternativePaymentTypesGetRequest) NewRequestBody() AlternativePaymentTypesGetRequestBody {
	return AlternativePaymentTypesGetRequestBody{}
}

type AlternativePaymentTypesGetRequestBody struct {
}

func (r *AlternativePaymentTypesGetRequest) RequestBody() *AlternativePaymentTypesGetRequestBody {
	return nil
}

func (r *AlternativePaymentTypesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AlternativePaymentTypesGetRequest) SetRequestBody(body AlternativePaymentTypesGetRequestBody) {
	r.requestBody = body
}

func (r *AlternativePaymentTypesGetRequest) NewResponseBody() *AlternativePaymentTypesGetResponseBody {
	return &AlternativePaymentTypesGetResponseBody{}
}

type AlternativePaymentTypesGetResponseBody AlternativePaymentTypes

func (r *AlternativePaymentTypesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/alternatePaymentTypes/", r.PathParams())
	return &u
}

func (r *AlternativePaymentTypesGetRequest) Do() (AlternativePaymentTypesGetResponseBody, error, *http.Response) {
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

func (r *AlternativePaymentTypesGetRequest) All() (AlternativePaymentTypesGetResponseBody, error) {
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
