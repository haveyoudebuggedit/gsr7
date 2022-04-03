package gsr7

type ClientResponse interface {
    Response[ClientResponse, ReadableStream]
}
