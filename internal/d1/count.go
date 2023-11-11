package d1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) Count(ctx context.Context) (int, error) {
	endpoint, err := c.baseURL.Parse(fmt.Sprintf("./accounts/%s/d1/database", c.accountId))

	if err != nil {
		return 0, fmt.Errorf("d1: Count: %w", err)
	}

	values := url.Values{}

	values.Set("per_page", "1")
	values.Set("page", "1")

	endpoint.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)

	if err != nil {
		return 0, fmt.Errorf("d1: Count: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)

	if err != nil {
		return 0, fmt.Errorf("d1: Count: %w", err)
	}

	defer res.Body.Close()

	var listResponse ListResponse

	if err := json.NewDecoder(res.Body).Decode(&listResponse); err != nil {
		return 0, fmt.Errorf("d1: Count: %w", err)
	}

	return listResponse.ResultInfo.TotalCount, nil
}
