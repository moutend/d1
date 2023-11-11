package d1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type VerifyResponse struct {
	Result   VerifyResult `json:"result"`
	Success  bool         `json:"success"`
	Errors   []Error      `json:"errors"`
	Messages []Message    `json:"messages"`
}

type VerifyResult struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	ExpiresOn time.Time `json:"expires_on"`
}

func (c *Client) Verify(ctx context.Context) (*VerifyResponse, error) {
	endpoint, err := c.baseURL.Parse("./user/tokens/verify")

	if err != nil {
		return nil, fmt.Errorf("d1: Verify: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("d1: Verify: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("d1: Verify: %w", err)
	}

	defer res.Body.Close()

	var verifyResponse VerifyResponse

	if err := json.NewDecoder(res.Body).Decode(&verifyResponse); err != nil {
		return nil, fmt.Errorf("d1: Verify: %w", err)
	}

	return &verifyResponse, nil
}
