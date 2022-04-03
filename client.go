package gsr7

// Client defines the methods for a HTTP client.
type Client interface {
	Request(request ClientRequest) (ClientResponse, error)
}
