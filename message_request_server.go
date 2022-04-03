package psr7

type ServerRequest interface {
    Request[ServerRequest, ReadableStream]
}
