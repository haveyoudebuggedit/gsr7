package gsr7

type response[T message, R any] interface {
    message[T, R, ResponseCookie]

    GetStatusCode() uint16
    WithStatus(code uint16, reasonPhrase string) (T, error)
    GetReasonPhrase() string
}
