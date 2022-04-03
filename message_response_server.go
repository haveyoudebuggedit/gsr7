package gsr7

type ServerResponse interface {
    Response[ServerResponse, WritableStream]
}
