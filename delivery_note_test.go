package sap_test

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/omniboost/go-sap-b1-service-layer"
	"testing"
)

func TestDeliveryNoteGetAllRequest(t *testing.T) {
	spew.Dump(sabClient.NewDeliveryNoteGetAllRequest(
		sap.DeliveryNoteGetAllWithFilter("DocDate eq '2024-09-23'"),
	).Do(context.Background()))
}
