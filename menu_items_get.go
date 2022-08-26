package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sap-b1-service-layer/utils"
)

func (c *Client) NewMenuItemsGetRequest() MenuItemsGetRequest {
	r := MenuItemsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuItemsGetRequest struct {
	client      *Client
	queryParams *MenuItemsGetRequestQueryParams
	pathParams  *MenuItemsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuItemsGetRequestBody
}

func (r MenuItemsGetRequest) NewQueryParams() *MenuItemsGetRequestQueryParams {
	return &MenuItemsGetRequestQueryParams{}
}

type MenuItemsGetRequestQueryParams struct {
	PageToken string `schema:"pageToken,omitempty"`
}

func (p MenuItemsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuItemsGetRequest) QueryParams() *MenuItemsGetRequestQueryParams {
	return r.queryParams
}

func (r MenuItemsGetRequest) NewPathParams() *MenuItemsGetRequestPathParams {
	return &MenuItemsGetRequestPathParams{}
}

type MenuItemsGetRequestPathParams struct {
}

func (p *MenuItemsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MenuItemsGetRequest) PathParams() *MenuItemsGetRequestPathParams {
	return r.pathParams
}

func (r *MenuItemsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuItemsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuItemsGetRequest) Method() string {
	return r.method
}

func (r MenuItemsGetRequest) NewRequestBody() MenuItemsGetRequestBody {
	return MenuItemsGetRequestBody{}
}

type MenuItemsGetRequestBody struct {
}

func (r *MenuItemsGetRequest) RequestBody() *MenuItemsGetRequestBody {
	return nil
}

func (r *MenuItemsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuItemsGetRequest) SetRequestBody(body MenuItemsGetRequestBody) {
	r.requestBody = body
}

func (r *MenuItemsGetRequest) NewResponseBody() *MenuItemsGetResponseBody {
	return &MenuItemsGetResponseBody{}
}

type MenuItemsGetResponseBody MenuItems

func (r *MenuItemsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/menuItems/", r.PathParams())
	return &u
}

func (r *MenuItemsGetRequest) Do() (MenuItemsGetResponseBody, error, *http.Response) {
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

func (r *MenuItemsGetRequest) All() (MenuItemsGetResponseBody, error) {
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
