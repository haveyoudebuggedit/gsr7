package gsr7

type URI interface {
    String() string

    GetScheme() string
    GetAuthority() string
    GetUserInfo() string
    GetHost() string
    GetPort() *uint16
    GetPath() string
    GetQuery() string
    GetFragment() string

    WithScheme(scheme string) (URI, error)
    WithUserInfo(user, password string) URI
    WithHost(host string) (URI, error)
    WithPort(port *uint16) URI
    WithPath(path string) (URI, error)
    WithQuery(query string) (URI, error)
    WithFragment(fragment string) (URI, error)
}
