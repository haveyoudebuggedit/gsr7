package gsr7

type ClientRequest interface {
    Request[ClientRequest, WritableStream]
}
