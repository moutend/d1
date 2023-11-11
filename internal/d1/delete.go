package d1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteResponse struct {
	Errors  []Error `json:"errors"`
	Success bool    `json:"success"`
}

func (c *Client) Delete(ctx context.Context, uuid string) (*DeleteResponse, error) {
	endpoint, err := c.baseURL.Parse(fmt.Sprintf("./accounts/%s/d1/database/%s", c.accountId, uuid))

	if err != nil {
		return nil, fmt.Errorf("d1: Delete: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("d1: Delete: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("d1: Delete: %w", err)
	}

	defer res.Body.Close()

	var deleteResponse DeleteResponse

	if err := json.NewDecoder(res.Body).Decode(&deleteResponse); err != nil {
		return nil, fmt.Errorf("d1: Delete: %w", err)
	}

	return &deleteResponse, nil
}
