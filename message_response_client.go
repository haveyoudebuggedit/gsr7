package gsr7

type ClientResponse interface {
    response[ClientResponse, ReadableStream]
}
