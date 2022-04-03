package psr7

type ClientRequest interface {
    Request[ClientRequest, WritableStream]
}
