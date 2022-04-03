package psr7

type ClientResponse interface {
    Response[ClientResponse, ReadableStream]
}
