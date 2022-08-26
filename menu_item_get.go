package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewMenuItemGetRequest() MenuItemGetRequest {
	r := MenuItemGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuItemGetRequest struct {
	client      *Client
	queryParams *MenuItemGetRequestQueryParams
	pathParams  *MenuItemGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuItemGetRequestBody
}

func (r MenuItemGetRequest) NewQueryParams() *MenuItemGetRequestQueryParams {
	return &MenuItemGetRequestQueryParams{}
}

type MenuItemGetRequestQueryParams struct{}

func (p MenuItemGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuItemGetRequest) QueryParams() *MenuItemGetRequestQueryParams {
	return r.queryParams
}

func (r MenuItemGetRequest) NewPathParams() *MenuItemGetRequestPathParams {
	return &MenuItemGetRequestPathParams{}
}

type MenuItemGetRequestPathParams struct {
	GUID string `schema:"guid"`
}

func (p *MenuItemGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"guid": p.GUID,
	}
}

func (r *MenuItemGetRequest) PathParams() *MenuItemGetRequestPathParams {
	return r.pathParams
}

func (r *MenuItemGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuItemGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuItemGetRequest) Method() string {
	return r.method
}

func (r MenuItemGetRequest) NewRequestBody() MenuItemGetRequestBody {
	return MenuItemGetRequestBody{}
}

type MenuItemGetRequestBody struct {
}

func (r *MenuItemGetRequest) RequestBody() *MenuItemGetRequestBody {
	return nil
}

func (r *MenuItemGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuItemGetRequest) SetRequestBody(body MenuItemGetRequestBody) {
	r.requestBody = body
}

func (r *MenuItemGetRequest) NewResponseBody() *MenuItemGetResponseBody {
	return &MenuItemGetResponseBody{}
}

type MenuItemGetResponseBody MenuItem

func (r *MenuItemGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/menuItems/{{.guid}}", r.PathParams())
	return &u
}

func (r *MenuItemGetRequest) Do() (MenuItemGetResponseBody, error, *http.Response) {
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
