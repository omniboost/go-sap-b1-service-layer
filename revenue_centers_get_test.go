package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRevenueCentersGet(t *testing.T) {
	req := client.NewRevenueCentersGetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
