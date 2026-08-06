package sap_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestChartOfAccountsGet(t *testing.T) {
	req := client.NewChartOfAccountsGetRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
