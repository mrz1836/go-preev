package preev

// newMockClient returns a client for mocking (using a custom HTTP interface)
func newMockClient(httpClient HTTPInterface) ClientInterface {
	return NewClient(nil, httpClient)
}
