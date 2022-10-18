package sap

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sap-b1-service-layer/odata"
	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewItemsGetRequest() ItemsGetRequest {
	r := ItemsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
		path:    "/b1s/v1/Items",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ItemsGetRequest struct {
	client      *Client
	queryParams *ItemsGetRequestQueryParams
	path        string
	pathParams  *ItemsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ItemsGetRequestBody
}

func (r ItemsGetRequest) NewQueryParams() *ItemsGetRequestQueryParams {
	selectFields, _ := utils.Fields(&Item{})
	// expandFields := []string{}
	return &ItemsGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
		Count:  odata.NewCount(),
	}
}

type ItemsGetRequestQueryParams struct {
	Select *odata.Select `schema:"$select,omitempty"`
	// Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
	Count  *odata.Count  `schema:"$count,omitempty"`
}

func (p ItemsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ItemsGetRequest) QueryParams() *ItemsGetRequestQueryParams {
	return r.queryParams
}

func (r ItemsGetRequest) NewPathParams() *ItemsGetRequestPathParams {
	return &ItemsGetRequestPathParams{}
}

type ItemsGetRequestPathParams struct {
}

func (p *ItemsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ItemsGetRequest) Path() *string {
	return &r.path
}

func (r *ItemsGetRequest) PathParams() *ItemsGetRequestPathParams {
	return r.pathParams
}

func (r *ItemsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ItemsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ItemsGetRequest) Method() string {
	return r.method
}

func (r ItemsGetRequest) NewRequestBody() ItemsGetRequestBody {
	return ItemsGetRequestBody{}
}

type ItemsGetRequestBody struct{}

func (r *ItemsGetRequest) RequestBody() *ItemsGetRequestBody {
	return &r.requestBody
}

func (r *ItemsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ItemsGetRequest) SetRequestBody(body ItemsGetRequestBody) {
	r.requestBody = body
}

func (r *ItemsGetRequest) NewResponseBody() *ItemsGetResponseBody {
	return &ItemsGetResponseBody{}
}

type ItemsGetResponseBody struct {
	OdataMetadata string `json:"odata.metadata"`
	Values        Items  `json:"value"`
	OdataNextLink string `json:"odata.nextLink"`
}

func (r *ItemsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *ItemsGetRequest) Do() (ItemsGetResponseBody, error) {
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

func (r *ItemsGetRequest) All() (ItemsGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := ItemsGetResponseBody{}
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
