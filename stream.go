package psr7

import "io"

type ReadableStream interface {
    io.ReadSeekCloser
}

type WritableStream interface {
    io.WriteCloser
}
