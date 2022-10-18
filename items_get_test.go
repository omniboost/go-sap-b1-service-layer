package sap_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestItemsGet(t *testing.T) {
	req := client.NewItemsGetRequest()
	// req.QueryParams().Count.Set(20)
	// req.QueryParams().Filter.Set("ItemCode eq '80304018'")
	// req.QueryParams().Filter.Set("ItemCode eq '19200001'")
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
