package gsr7

type request[T request, R any] interface {
    message[T, R, RequestCookie]

    GetRequestTarget() string
    WithRequestTargetString(requestTarget string) (T, error)

    GetMethod() string
    WithMethod(method string) (T, error)

    GetURI() URI
    WithURI(uri URI, preserveHost bool) T
}
