package gsr7

type Response[ResponseType any, BodyType any] interface {
	Message[ResponseType, BodyType, ResponseCookie]

	GetStatusCode() uint16
	WithStatusCode(code uint16) (ResponseType, error)
	WithStatus(code uint16, reasonPhrase string) (ResponseType, error)
	GetReasonPhrase() string
}
