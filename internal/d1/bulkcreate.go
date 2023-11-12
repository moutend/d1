package d1

import (
	"context"
	"fmt"
	"time"
)

func (c *Client) BulkCreate(ctx context.Context, requests []CreateRequest) ([]*CreateResult, error) {
	const maxRetryCount = 5

	n := 0
	retry := 0
	results := []*CreateResult{}

	for n < len(requests) {
		if requests[n].Name == "" {
			return results, fmt.Errorf("d1: BulkCreate: database name is empty")
		}

		c.debug.Printf("d1: BulkCreate: creating: database name: %s: retry: %d\n", requests[n].Name, retry)

		result, err := c.Create(ctx, requests[n].Name)

		if err != nil {
			c.debug.Printf("d1: BulkCreate: received unexpected response: database name: %s: %s", requests[n].Name, err)

			retry += 1
			time.Sleep(time.Minute)

			if retry <= maxRetryCount {
				continue
			} else {
				return results, fmt.Errorf("d1: BulkCreate: received unexpected response: database name: %s: %w", requests[n].Name, err)
			}
		}
		if !result.Success {
			c.debug.Printf("d1: BulkCreate: received error response: database name: %s: code: %d: message: %s\n", requests[n].Name, result.Errors[0].Code, result.Errors[0].Message)

			retry += 1
			time.Sleep(30 * time.Second)

			switch result.Errors[0].Code {
			case 7502:
				// The error code 7502 means already exists, so that continue creating databases.
				results = append(results, &CreateResult{
					Name: requests[n].Name,
				})

				retry = 0
				n += 1

				continue
			}
			if retry <= maxRetryCount {
				continue
			} else {
				return results, fmt.Errorf("d1: BulkCreate: received error response: database name: %s: code: %d: message: %s", requests[n].Name, result.Errors[0].Code, result.Errors[0].Message)
			}
		}

		results = append(results, result.Result)

		retry = 0
		n += 1

		time.Sleep(500 * time.Millisecond)
	}

	return results, nil
}
