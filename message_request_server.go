package gsr7

type ServerRequest interface {
    request[ServerRequest, ReadableStream]
}
