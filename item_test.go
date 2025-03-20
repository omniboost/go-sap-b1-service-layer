package sap_test

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/omniboost/go-sap-b1-service-layer"
	"testing"
)

func TestItemGetAllRequest(t *testing.T) {
	spew.Dump(sabClient.NewItemGetAllRequest(
		sap.ItemGetAllWithSelect("ItemCode", "ItemName"),
	).Do(context.Background()))
}
