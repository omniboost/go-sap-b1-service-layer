package sap_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalEntriesPost(t *testing.T) {
	req := client.NewJournalEntriesPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
