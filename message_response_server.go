package psr7

type ServerResponse interface {
    Response[ServerResponse, WritableStream]
}
