package sap_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoginPost(t *testing.T) {
	req := client.NewLoginPostRequest()
	req.RequestBody().Username = client.Username()
	req.RequestBody().Password = client.Password()
	req.RequestBody().CompanyDB = client.CompanyDB()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
