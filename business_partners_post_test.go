package sap_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestBusinessPartnersPost(t *testing.T) {
	req := client.NewBusinessPartnersPostRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
