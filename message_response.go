package psr7

type Response[T Message] interface {
    Message[T]

    GetStatusCode() uint16
    WithStatus(code uint16, reasonPhrase string) (T, error)
    GetReasonPhrase() string
}
