package gsr7

type ClientRequest interface {
    request[ClientRequest, WritableStream]
}
