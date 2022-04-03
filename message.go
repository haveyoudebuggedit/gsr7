package psr7

type Message[T any, R any] interface {
    GetProtocolVersion() string
    WithProtocolVersion(version string) T

    GetHeaders() [][]string

    HasHeader(name string) bool
    GetHeader(name string) []string
    GetHeaderLine(name string) string

    WithHeader(name, value string) (T, error)
    WithHeaderValues(name, value []string) (T, error)
    WithAddedHeader(name, value string) (T, error)
    WithoutHeader(name string) T

    GetBody() R
    WithBody(R) (ClientRequest, error)
}
