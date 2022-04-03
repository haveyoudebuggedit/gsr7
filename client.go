package gsr7

type Client interface {
    Request(request ClientRequest) (ClientResponse, error)
}
