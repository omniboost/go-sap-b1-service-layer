package sap_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalEntriesGet(t *testing.T) {
	req := client.NewJournalEntriesGetRequest()
	// req.QueryParams().Count.Set(20)
	// req.QueryParams().Filter.Set("ItemCode eq '80304018'")
	req.QueryParams().Filter.Set("contains(ReferenceDate,'2023')")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
