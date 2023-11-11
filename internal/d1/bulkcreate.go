package d1

import (
	"context"
	"fmt"
	"time"
)

func (c *Client) BulkCreate(ctx context.Context, requests []CreateRequest) ([]*CreateResult, error) {
	n := 0
	retry := 0
	results := []*CreateResult{}

	for n < len(requests) {
		result, err := c.Create(ctx, requests[n].Name)

		if err != nil {
			retry += 1
			time.Sleep(10 * time.Second)

			if retry <= 3 {
				continue
			} else {
				return results, fmt.Errorf("d1: BulkCreate: %w", err)
			}
		}
		if !result.Success {
			retry += 1
			time.Sleep(5 * time.Second)

			if retry <= 3 {
				continue
			} else {
				return results, fmt.Errorf("d1: BulkCreate: %w", err)
			}
		}

		results = append(results, result.Result)

		retry = 0
		n += 1

		time.Sleep(250 * time.Millisecond)
	}

	return results, nil
}
