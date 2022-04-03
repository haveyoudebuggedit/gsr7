package gsr7

type Version interface {
    String() string
    Major() uint8
    Minor() uint8
}

type Message[MessageType any, BodyType any, CookieType Cookie] interface {
    GetProtocolVersion() Version
    WithProtocolVersion(version Version) MessageType

    GetHeaders() [][]string

    HasHeader(name string) bool
    GetHeader(name string) []string
    GetHeaderLine(name string) string

    WithHeader(name, value string) (MessageType, error)
    WithHeaderValues(name, value []string) (MessageType, error)
    WithAddedHeader(name, value string) (MessageType, error)
    WithoutHeader(name string) MessageType

    GetBody() BodyType
    WithBody(BodyType) (ClientRequest, error)

    GetCookies() []CookieType
    WithCookie(cookie CookieType) MessageType
    WithCookies(cookies []CookieType) MessageType
}
