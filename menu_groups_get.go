package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewMenuGroupsGetRequest() MenuGroupsGetRequest {
	r := MenuGroupsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuGroupsGetRequest struct {
	client      *Client
	queryParams *MenuGroupsGetRequestQueryParams
	pathParams  *MenuGroupsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuGroupsGetRequestBody
}

func (r MenuGroupsGetRequest) NewQueryParams() *MenuGroupsGetRequestQueryParams {
	return &MenuGroupsGetRequestQueryParams{}
}

type MenuGroupsGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p MenuGroupsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuGroupsGetRequest) QueryParams() *MenuGroupsGetRequestQueryParams {
	return r.queryParams
}

func (r MenuGroupsGetRequest) NewPathParams() *MenuGroupsGetRequestPathParams {
	return &MenuGroupsGetRequestPathParams{}
}

type MenuGroupsGetRequestPathParams struct {
}

func (p *MenuGroupsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MenuGroupsGetRequest) PathParams() *MenuGroupsGetRequestPathParams {
	return r.pathParams
}

func (r *MenuGroupsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuGroupsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuGroupsGetRequest) Method() string {
	return r.method
}

func (r MenuGroupsGetRequest) NewRequestBody() MenuGroupsGetRequestBody {
	return MenuGroupsGetRequestBody{}
}

type MenuGroupsGetRequestBody struct {
}

func (r *MenuGroupsGetRequest) RequestBody() *MenuGroupsGetRequestBody {
	return nil
}

func (r *MenuGroupsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuGroupsGetRequest) SetRequestBody(body MenuGroupsGetRequestBody) {
	r.requestBody = body
}

func (r *MenuGroupsGetRequest) NewResponseBody() *MenuGroupsGetResponseBody {
	return &MenuGroupsGetResponseBody{}
}

type MenuGroupsGetResponseBody MenuGroups

func (r *MenuGroupsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/menuGroups", r.PathParams())
	return &u
}

func (r *MenuGroupsGetRequest) Do() (MenuGroupsGetResponseBody, error, *http.Response) {
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

func (r *MenuGroupsGetRequest) All() (MenuGroupsGetResponseBody, error) {
	body, err, resp := r.Do()
	if err != nil {
		return body, err
	}

	concat := body
	next, err := r.client.GetNextPage(resp)
	if err != nil {
		return concat, err
	}

	for next != 0 {
		r.QueryParams().Page = next
		body, err, resp = r.Do()
		if err != nil {
			return concat, err
		}

		concat = append(concat, body...)
		next, err = r.client.GetNextPage(resp)
		if err != nil {
			return concat, err
		}
	}

	return concat, nil
}
