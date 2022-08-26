package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAlternativePaymentTypesGet(t *testing.T) {
	req := client.NewAlternativePaymentTypesGetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
