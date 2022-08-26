package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMenuItemsGet(t *testing.T) {
	req := client.NewMenuItemsGetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
