package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-toast"
)

func TestOrdersBulkGet(t *testing.T) {
	req := client.NewOrdersBulkGetRequest()
	req.QueryParams().StartDate = toast.DateTime{time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().EndDate = toast.DateTime{time.Date(2022, 6, 13, 0, 0, 0, 0, time.UTC)}
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
