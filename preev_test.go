package preev

// newMockClient returns a client for mocking
func newMockClient(httpClient httpInterface) *Client {
	client := NewClient(nil)
	client.httpClient = httpClient
	return client
}
