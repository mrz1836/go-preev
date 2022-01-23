/*
Package preev is the unofficial golang implementation for the Preev.pro API

Example:

```
// Create a client:
client := preev.NewClient(nil, nil)

// Get pairs:
pairs, _ := client.GetPairs()

fmt.Println("Found Active Pair(s):", pairs.BsvUsd.Name)
```
*/
package preev

import (
	"context"
	"io/ioutil"
	"net/http"
)

// NewClient creates a new client for Preev requests
func NewClient(clientOptions *Options, customHTTPClient HTTPInterface) ClientInterface {
	return createClient(clientOptions, customHTTPClient)
}

// request is a generic request wrapper that can be used without constraints
func (c *Client) request(ctx context.Context, url string) (response string, err error) {

	// Store for debugging purposes
	c.lastRequest.Method = http.MethodGet
	c.lastRequest.URL = url

	// Start the request
	var request *http.Request
	if request, err = http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil,
	); err != nil {
		return
	}

	// Change the header (user agent is in case they block default Go user agents)
	request.Header.Set("User-Agent", c.UserAgent())

	// Fire the http request
	var resp *http.Response
	if resp, err = c.httpClient.Do(request); err != nil {
		if resp != nil {
			c.lastRequest.StatusCode = resp.StatusCode
		}
		return
	}

	// Close the response body
	defer func() {
		_ = resp.Body.Close()
	}()

	// Set the status
	c.lastRequest.StatusCode = resp.StatusCode

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Return the raw JSON response
	response = string(body)
	return
}
