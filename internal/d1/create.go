package d1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreateRequest struct {
	Name         string `json:"name"`
	Experimental bool   `json:"experimental"`
	Location     string `json:"location"`
}

type CreateResponse struct {
	Result   *CreateResult `json:"result"`
	Errors   []Error       `json:"errors"`
	Messages []Message     `json:"messages"`
	Success  bool          `json:"success"`
}

type CreateResult struct {
	UUID            string    `json:"uuid"`
	Name            string    `json:"name"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedInResion string    `json:"created_in_region"`
	Version         string    `json:"version"`
}

func (c *Client) Create(ctx context.Context, name string) (*CreateResponse, error) {
	d1CreateRequest, err := json.Marshal(CreateRequest{
		Name:         name,
		Experimental: true,
		Location:     c.location,
	})

	if err != nil {
		return nil, fmt.Errorf("d1: Create: %w", err)
	}

	endpoint, err := c.baseURL.Parse(fmt.Sprintf("./accounts/%s/d1/database", c.accountId))

	if err != nil {
		return nil, fmt.Errorf("d1: Create: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewBuffer(d1CreateRequest))

	if err != nil {
		return nil, fmt.Errorf("d1: Create: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("d1: Create: %w", err)
	}

	defer res.Body.Close()

	var createResponse CreateResponse

	if err := json.NewDecoder(res.Body).Decode(&createResponse); err != nil {
		return nil, fmt.Errorf("d1: Create: %w", err)
	}

	return &createResponse, nil
}
