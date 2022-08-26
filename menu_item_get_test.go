package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMenuItemGet(t *testing.T) {
	req := client.NewMenuItemGetRequest()
	req.PathParams().GUID = "b3550ab4-269b-4f48-8027-4ac1b2369802"
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
