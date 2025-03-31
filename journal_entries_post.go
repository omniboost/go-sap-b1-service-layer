package sap

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewJournalEntriesPostRequest() JournalEntriesPostRequest {
	r := JournalEntriesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
		path:    "/b1s/v1/JournalEntries",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntriesPostRequest struct {
	client      *Client
	queryParams *JournalEntriesPostRequestQueryParams
	path        string
	pathParams  *JournalEntriesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntriesPostRequestBody
}

func (r JournalEntriesPostRequest) NewQueryParams() *JournalEntriesPostRequestQueryParams {
	return &JournalEntriesPostRequestQueryParams{}
}

type JournalEntriesPostRequestQueryParams struct {
}

func (p JournalEntriesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntriesPostRequest) QueryParams() *JournalEntriesPostRequestQueryParams {
	return r.queryParams
}

func (r JournalEntriesPostRequest) NewPathParams() *JournalEntriesPostRequestPathParams {
	return &JournalEntriesPostRequestPathParams{}
}

type JournalEntriesPostRequestPathParams struct {
}

func (p *JournalEntriesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalEntriesPostRequest) Path() *string {
	return &r.path
}

func (r *JournalEntriesPostRequest) PathParams() *JournalEntriesPostRequestPathParams {
	return r.pathParams
}

func (r *JournalEntriesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntriesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntriesPostRequest) Method() string {
	return r.method
}

func (r JournalEntriesPostRequest) NewRequestBody() JournalEntriesPostRequestBody {
	return JournalEntriesPostRequestBody{}
}

type JournalEntriesPostRequestBody JournalEntry

func (r *JournalEntriesPostRequest) RequestBody() *JournalEntriesPostRequestBody {
	return &r.requestBody
}

func (r *JournalEntriesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalEntriesPostRequest) SetRequestBody(body JournalEntriesPostRequestBody) {
	r.requestBody = body
}

func (r *JournalEntriesPostRequest) NewResponseBody() *JournalEntriesPostResponseBody {
	return &JournalEntriesPostResponseBody{}
}

type JournalEntriesPostResponseBody JournalEntry

func (r *JournalEntriesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *JournalEntriesPostRequest) Do() (JournalEntriesPostResponseBody, error) {
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
