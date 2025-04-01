package sap

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/omniboost/go-omniboost-http-client/client"
)

type (
	SapB1Client struct {
		client.Client
	}
)

func NewSapB1Client(opts ...client.Option) *SapB1Client {
	c := &SapB1Client{}
	opts = append([]client.Option{
		client.WithBaseURL(BaseURL),
		client.WithUserAgent(userAgent),
		client.WithUseCookies(true),
		client.WithParentClient(c),
	}, opts...)

	c.Client = client.NewClient(opts...)
	return c
}

func WithAuth(username, password, companyDB string) client.Option {
	hasRun := false
	m := sync.Mutex{}
	return client.WithPreflightAuth(func(req *http.Request, c client.Client) (*http.Request, error) {
		m.Lock()
		defer m.Unlock()

		if hasRun {
			return req, nil
		}

		sc := c.GetParentClient().(*SapB1Client)
		err := sc.NewLoginRequest(LoginWithCredentials(username, password, companyDB)).Do(req.Context())
		if err != nil {
			return nil, err
		}
		hasRun = true

		return req, nil
	})
}

func getSkip(nextLink string) (int, error) {
	if nextLink == "" {
		return 0, nil
	}

	u, err := url.Parse(nextLink)
	if err != nil {
		return 0, err
	}
	q := u.Query()

	skip := q.Get("$skip")
	if skip == "" {
		return 0, errors.New("missing $skip in next link")
	}

	return strconv.Atoi(skip)
}
