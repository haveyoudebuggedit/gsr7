package gsr7

type Request[RequestType any, BodyType any] interface {
	Message[RequestType, BodyType, RequestCookie]

	GetRequestTarget() string
	WithRequestTargetString(requestTarget string) (RequestType, error)

	GetMethod() string
	WithMethod(method string) RequestType
	WithMethodE(method string) (RequestType, error)

	GetURI() URI
	WithURI(uri URI) RequestType
	WithURIPreserveHost(uri URI) RequestType
}
