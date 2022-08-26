package toast_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	asperion "github.com/omniboost/go-sap-b1-service-layer"
)

var (
	client *asperion.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	toastRestaurantExternalID := os.Getenv("TOAST_RESTAURANT_EXTERNAL_ID")
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

	client = asperion.NewClient(nil)
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetToastRestaurantExternalID(toastRestaurantExternalID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
