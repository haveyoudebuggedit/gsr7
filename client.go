package gsr7

// Client defines the methods for a HTTP client.
type Client interface {
	// Request sends a HTTP request and returns a response. An error is returned if a connection error happens. No
	// error is returned if the server responds with a non-200 status code.
	Request(request ClientRequest) (ClientResponse, error)
}
