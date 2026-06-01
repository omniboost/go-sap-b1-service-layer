package sap

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sap-b1-service-layer/odata"
	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewBusinessPartnersGetRequest() BusinessPartnersGetRequest {
	r := BusinessPartnersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
		path:    "/b1s/v1/BusinessPartners",
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BusinessPartnersGetRequest struct {
	client      *Client
	queryParams *BusinessPartnersGetRequestQueryParams
	path        string
	pathParams  *BusinessPartnersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody BusinessPartnersGetRequestBody
}

func (r BusinessPartnersGetRequest) NewQueryParams() *BusinessPartnersGetRequestQueryParams {
	return &BusinessPartnersGetRequestQueryParams{
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
		Count:  odata.NewCount(),
	}
}

type BusinessPartnersGetRequestQueryParams struct {
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
	Count  *odata.Count  `schema:"$count,omitempty"`
}

func (p BusinessPartnersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BusinessPartnersGetRequest) QueryParams() *BusinessPartnersGetRequestQueryParams {
	return r.queryParams
}

func (r BusinessPartnersGetRequest) NewPathParams() *BusinessPartnersGetRequestPathParams {
	return &BusinessPartnersGetRequestPathParams{}
}

type BusinessPartnersGetRequestPathParams struct {
}

func (p *BusinessPartnersGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BusinessPartnersGetRequest) Path() *string {
	return &r.path
}

func (r *BusinessPartnersGetRequest) PathParams() *BusinessPartnersGetRequestPathParams {
	return r.pathParams
}

func (r *BusinessPartnersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BusinessPartnersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BusinessPartnersGetRequest) Method() string {
	return r.method
}

func (r BusinessPartnersGetRequest) NewRequestBody() BusinessPartnersGetRequestBody {
	return BusinessPartnersGetRequestBody{}
}

type BusinessPartnersGetRequestBody struct{}

func (r *BusinessPartnersGetRequest) RequestBody() *BusinessPartnersGetRequestBody {
	return &r.requestBody
}

func (r *BusinessPartnersGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BusinessPartnersGetRequest) SetRequestBody(body BusinessPartnersGetRequestBody) {
	r.requestBody = body
}

func (r *BusinessPartnersGetRequest) NewResponseBody() *BusinessPartnersGetResponseBody {
	return &BusinessPartnersGetResponseBody{}
}

type BusinessPartnersGetResponseBody struct {
	OdataMetadata string           `json:"odata.metadata"`
	Values        BusinessPartners `json:"value"`
	OdataNextLink string           `json:"odata.nextLink"`
}

func (r *BusinessPartnersGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL(*r.Path(), r.PathParams())
	return &u
}

func (r *BusinessPartnersGetRequest) Do() (BusinessPartnersGetResponseBody, error) {
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = r.client.InitSession(req)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *BusinessPartnersGetRequest) All() (BusinessPartnersGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := BusinessPartnersGetResponseBody{}
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
