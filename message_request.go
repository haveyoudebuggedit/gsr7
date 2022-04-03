package psr7

type Request[T Request, R any] interface {
    Message[T, R]

    GetRequestTarget() string
    WithRequestTargetString(requestTarget string) (T, error)

    GetMethod() string
    WithMethod(method string) (T, error)

    GetURI() URI
    WithURI(uri URI, preserveHost bool) T
}
