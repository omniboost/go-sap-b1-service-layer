package sap_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalEntriesPost(t *testing.T) {
	req := client.NewJournalEntriesPostRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
