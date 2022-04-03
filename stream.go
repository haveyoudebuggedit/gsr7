package gsr7

import "io"

type ReadableStream interface {
    io.ReadSeekCloser

    String() string
    Bytes() []byte
}

type WritableStream interface {
    io.WriteCloser
}
