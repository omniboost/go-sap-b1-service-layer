package sap

import (
	"context"
	"errors"
	"github.com/omniboost/go-omniboost-http-client/client"
	"github.com/omniboost/go-sap-b1-service-layer/odata"
	"github.com/omniboost/go-sap-b1-service-layer/utils"
	"net/http"
)

type (
	ItemGetAllRequest struct {
		client *SapB1Client

		Select *odata.Select `query:"$select,omitempty"`
		Skip   odata.Skip    `query:"$skip,omitempty"`
		Top    odata.Top     `query:"$top,omitempty"`
	}

	ItemGetAllResponse struct {
		Context  string `json:"@odata.context"`
		Values   Items  `json:"value"`
		NextLink string `json:"@odata.nextLink"`
	}
	ItemGetAllOption func(*ItemGetAllRequest)
)

func (i *ItemGetAllRequest) Method() string {
	return http.MethodGet
}

func (i *ItemGetAllRequest) PathTemplate() string {
	return "/b1s/v2/Items"
}

var _ client.Request = (*ItemGetAllRequest)(nil)

func ItemGetAllWithSelect(selectFields ...string) ItemGetAllOption {
	return func(r *ItemGetAllRequest) {
		for _, field := range selectFields {
			r.Select.Add(field)
		}
	}
}

func (s *SapB1Client) NewItemGetAllRequest(opts ...ItemGetAllOption) *ItemGetAllRequest {
	fields, _ := utils.Fields(&Item{})

	r := &ItemGetAllRequest{
		client: s,
		Select: odata.NewSelect(fields),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (i *ItemGetAllRequest) Do(ctx context.Context) (Items, error) {
	items := make(Items, 0)

	maxRequests := 2
	for {
		var resp ItemGetAllResponse
		err := i.client.Do(ctx, i, &resp)
		if err != nil {
			return nil, err
		}

		if len(resp.Values) == 0 {
			break
		}

		items = append(items, resp.Values...)

		skip, err := getSkip(resp.NextLink)
		if err != nil {
			return nil, err
		}
		if skip == 0 {
			break
		}
		i.Skip.Set(skip)

		maxRequests--
		if maxRequests == 0 {
			return nil, errors.New("max requests exceeded")
		}
	}

	return items, nil
}
