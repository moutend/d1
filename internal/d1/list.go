package d1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ListResponse struct {
	Result     []*ListResult  `json:"result"`
	ResultInfo ListResultInfo `json:"result_info"`
	Success    bool           `json:"success"`
	Errors     []Error        `json:"errors"`
}

type ListResult struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Version   string    `json:"version"`
}

type ListResultInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

func (c *Client) List(ctx context.Context) ([]*ListResult, error) {
	endpoint, err := c.baseURL.Parse(fmt.Sprintf("./accounts/%s/d1/database", c.accountId))

	if err != nil {
		return nil, fmt.Errorf("d1: List: %w", err)
	}

	const pageSize = 100

	page := 1
	results := []*ListResult{}

	for len(results)%pageSize == 0 {
		values := url.Values{}

		values.Set("per_page", fmt.Sprint(pageSize))
		values.Set("page", fmt.Sprint(page))

		endpoint.RawQuery = values.Encode()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)

		if err != nil {
			return results, fmt.Errorf("d1: List: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

		res, err := c.httpClient.Do(req)

		if err != nil {
			return results, fmt.Errorf("d1: List: %w", err)
		}

		defer res.Body.Close()

		page += 1

		var listResponse ListResponse

		if err := json.NewDecoder(res.Body).Decode(&listResponse); err != nil {
			return results, fmt.Errorf("d1: List: %w", err)
		}

		results = append(results, listResponse.Result...)

		if len(listResponse.Result) < pageSize {
			break
		}

		time.Sleep(250 * time.Millisecond)
	}

	return results, nil
}
