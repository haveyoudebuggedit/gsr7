package gsr7

type ServerResponse interface {
    response[ServerResponse, WritableStream]
}
