package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewTaxRatesGetRequest() TaxRatesGetRequest {
	r := TaxRatesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxRatesGetRequest struct {
	client      *Client
	queryParams *TaxRatesGetRequestQueryParams
	pathParams  *TaxRatesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxRatesGetRequestBody
}

func (r TaxRatesGetRequest) NewQueryParams() *TaxRatesGetRequestQueryParams {
	return &TaxRatesGetRequestQueryParams{}
}

type TaxRatesGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p TaxRatesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxRatesGetRequest) QueryParams() *TaxRatesGetRequestQueryParams {
	return r.queryParams
}

func (r TaxRatesGetRequest) NewPathParams() *TaxRatesGetRequestPathParams {
	return &TaxRatesGetRequestPathParams{}
}

type TaxRatesGetRequestPathParams struct {
}

func (p *TaxRatesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxRatesGetRequest) PathParams() *TaxRatesGetRequestPathParams {
	return r.pathParams
}

func (r *TaxRatesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxRatesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxRatesGetRequest) Method() string {
	return r.method
}

func (r TaxRatesGetRequest) NewRequestBody() TaxRatesGetRequestBody {
	return TaxRatesGetRequestBody{}
}

type TaxRatesGetRequestBody struct {
}

func (r *TaxRatesGetRequest) RequestBody() *TaxRatesGetRequestBody {
	return nil
}

func (r *TaxRatesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxRatesGetRequest) SetRequestBody(body TaxRatesGetRequestBody) {
	r.requestBody = body
}

func (r *TaxRatesGetRequest) NewResponseBody() *TaxRatesGetResponseBody {
	return &TaxRatesGetResponseBody{}
}

type TaxRatesGetResponseBody TaxRates

func (r *TaxRatesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/taxRates/", r.PathParams())
	return &u
}

func (r *TaxRatesGetRequest) Do() (TaxRatesGetResponseBody, error, *http.Response) {
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

func (r *TaxRatesGetRequest) All() (TaxRatesGetResponseBody, error) {
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
