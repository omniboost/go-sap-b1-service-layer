package sap

import (
	"context"
	"github.com/omniboost/go-omniboost-http-client/client"
	"net/http"
)

type (
	DeliveryNoteCreateRequest struct {
		client *SapB1Client

		Note any
	}

	DeliveryNoteCreateResponse struct {
		Context  string        `json:"@odata.context"`
		Value    DeliveryNotes `json:"value"`
		NextLink string        `json:"@odata.nextLink"`
	}

	DeliveryNoteCreateOption func(*DeliveryNoteCreateRequest)
)

var _ client.RequestWithBody = (*DeliveryNoteCreateRequest)(nil)

func DeliveryNoteCreateWithFullDeliveryNote(deliveryNote *DeliveryNote) DeliveryNoteCreateOption {
	return func(r *DeliveryNoteCreateRequest) {
		r.Note = deliveryNote
	}
}

func DeliveryNoteCreateWithSimplifiedDeliveryNote(deliveryNote *SimplifiedDeliveryNote) DeliveryNoteCreateOption {
	return func(r *DeliveryNoteCreateRequest) {
		r.Note = deliveryNote
	}
}

func (s *SapB1Client) NewDeliveryNoteCreateRequest(opts ...DeliveryNoteCreateOption) *DeliveryNoteCreateRequest {
	r := &DeliveryNoteCreateRequest{
		client: s,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *DeliveryNoteCreateRequest) Method() string {
	return http.MethodPost
}

func (r *DeliveryNoteCreateRequest) PathTemplate() string {
	return "/b1s/v2/DeliveryNotes"
}

func (r *DeliveryNoteCreateRequest) Body() any {
	return r.Note
}

func (r *DeliveryNoteCreateRequest) Do(ctx context.Context) (*DeliveryNote, error) {
	var resp DeliveryNote
	err := r.client.Do(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
