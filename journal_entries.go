package sap

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sap-b1-service-layer/odata"
	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewJournalEntriesGetRequest() JournalEntriesGetRequest {
	r := JournalEntriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
		path:    "/b1s/v1/JournalEntries",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntriesGetRequest struct {
	client      *Client
	queryParams *JournalEntriesGetRequestQueryParams
	path        string
	pathParams  *JournalEntriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntriesGetRequestBody
}

func (r JournalEntriesGetRequest) NewQueryParams() *JournalEntriesGetRequestQueryParams {
	selectFields, _ := utils.Fields(&Item{})
	// expandFields := []string{}
	return &JournalEntriesGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
		Count:  odata.NewCount(),
	}
}

type JournalEntriesGetRequestQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	// Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
	Count  *odata.Count  `schema:"$count,omitempty"`
}

func (p JournalEntriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntriesGetRequest) QueryParams() *JournalEntriesGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntriesGetRequest) NewPathParams() *JournalEntriesGetRequestPathParams {
	return &JournalEntriesGetRequestPathParams{}
}

type JournalEntriesGetRequestPathParams struct {
}

func (p *JournalEntriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalEntriesGetRequest) Path() *string {
	return &r.path
}

func (r *JournalEntriesGetRequest) PathParams() *JournalEntriesGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntriesGetRequest) Method() string {
	return r.method
}

func (r JournalEntriesGetRequest) NewRequestBody() JournalEntriesGetRequestBody {
	return JournalEntriesGetRequestBody{}
}

type JournalEntriesGetRequestBody struct{}

func (r *JournalEntriesGetRequest) RequestBody() *JournalEntriesGetRequestBody {
	return &r.requestBody
}

func (r *JournalEntriesGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalEntriesGetRequest) SetRequestBody(body JournalEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntriesGetRequest) NewResponseBody() *JournalEntriesGetResponseBody {
	return &JournalEntriesGetResponseBody{}
}

type JournalEntriesGetResponseBody struct {
	OdataMetadata string         `json:"odata.metadata"`
	Values        JournalEntries `json:"value"`
	OdataNextLink string         `json:"odata.nextLink"`
}

func (r *JournalEntriesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *JournalEntriesGetRequest) Do() (JournalEntriesGetResponseBody, error) {
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

func (r *JournalEntriesGetRequest) All() (JournalEntriesGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := JournalEntriesGetResponseBody{}
	concat.Values = append(concat.Values, resp.Values...)

	for resp.OdataNextLink != "" {
		u, err := url.Parse(resp.OdataNextLink)
		if err != nil {
			return resp, err
		}

		skip, err := strconv.Atoi(u.Query().Get("$skip"))
		if err != nil {
			return resp, err
		}

		r.QueryParams().Skip.Set(skip)
		resp, err = r.Do()
		if err != nil {
			return resp, err
		}

		concat.Values = append(concat.Values, resp.Values...)
	}

	return concat, nil
}
