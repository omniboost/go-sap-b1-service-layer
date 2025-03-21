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
	DeliveryNoteGetAllRequest struct {
		client *SapB1Client

		Select *odata.Select `query:"$select,omitempty"`
		Filter *odata.Filter `query:"$filter,omitempty"`
		Top    odata.Top     `query:"$top,omitempty"`
		Skip   odata.Skip    `query:"$skip,omitempty"`
	}

	DeliveryNoteGetAllResponse struct {
		Context  string        `json:"@odata.context"`
		Value    DeliveryNotes `json:"value"`
		NextLink string        `json:"@odata.nextLink"`
	}

	DeliveryNoteGetAllOption func(*DeliveryNoteGetAllRequest)
)

var _ client.Request = (*DeliveryNoteGetAllRequest)(nil)

func DeliveryNoteGetAllWithSelect(selectFields ...string) DeliveryNoteGetAllOption {
	return func(r *DeliveryNoteGetAllRequest) {
		for _, field := range selectFields {
			r.Select.Add(field)
		}
	}
}

func DeliveryNoteGetAllWithFilter(filter string) DeliveryNoteGetAllOption {
	return func(r *DeliveryNoteGetAllRequest) {
		r.Filter.Set(filter)
	}
}

func (s *SapB1Client) NewDeliveryNoteGetAllRequest(opts ...DeliveryNoteGetAllOption) *DeliveryNoteGetAllRequest {
	fields, _ := utils.Fields(&DeliveryNote{})

	r := &DeliveryNoteGetAllRequest{
		client: s,
		Select: odata.NewSelect(fields),
		Filter: odata.NewFilter(),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *DeliveryNoteGetAllRequest) Method() string {
	return http.MethodGet
}

func (r *DeliveryNoteGetAllRequest) PathTemplate() string {
	return "/b1s/v2/DeliveryNotes"
}

func (r *DeliveryNoteGetAllRequest) Do(ctx context.Context) (DeliveryNotes, error) {
	deliveryNotes := make(DeliveryNotes, 0)

	maxRequests := 10_000
	for {
		var resp DeliveryNoteGetAllResponse
		err := r.client.Do(ctx, r, &resp)
		if err != nil {
			return nil, err
		}

		deliveryNotes = append(deliveryNotes, resp.Value...)

		if resp.NextLink == "" {
			break
		}

		skip, err := getSkip(resp.NextLink)
		if err != nil {
			return nil, err
		}
		if skip == 0 {
			break
		}
		r.Skip.Set(skip)

		maxRequests--
		if maxRequests == 0 {
			return nil, errors.New("max requests exceeded")
		}

		maxRequests--
		if maxRequests == 0 {
			break
		}
	}

	return deliveryNotes, nil
}
