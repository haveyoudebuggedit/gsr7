package gsr7

type ServerRequest interface {
    Request[ServerRequest, ReadableStream]
}
