package d1

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	BaseURL = "https://api.cloudflare.com/client/v4/"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	accountId  string
	apiToken   string
	location   string
}

func NewClient(httpClient *http.Client, accountId, apiToken, location string) (*Client, error) {
	if accountId == "" {
		return nil, fmt.Errorf("d1: accountId is empty")
	}
	if apiToken == "" {
		return nil, fmt.Errorf("d1: apiToken is empty")
	}
	if location == "" {
		return nil, fmt.Errorf("d1: location is empty")
	}

	baseURL, err := url.Parse(BaseURL)

	if err != nil {
		return nil, fmt.Errorf("d1: NewClient: %w", err)
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		accountId:  accountId,
		apiToken:   apiToken,
		location:   location,
	}, nil
}
