package sap_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	sap "github.com/omniboost/go-sap-b1-service-layer"
)

var (
	client *sap.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	username := os.Getenv("B1_USERNAME")
	password := os.Getenv("B1_PASSWORD")
	companyDB := os.Getenv("B1_COMPANY_DB")
	debug := os.Getenv("DEBUG")
	if err != nil {
		log.Fatal(err)
	}
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	httpClient := &http.Client{Transport: trans}

	client = sap.NewClient(httpClient)
	client.SetUsername(username)
	client.SetPassword(password)
	client.SetCompanyDB(companyDB)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
