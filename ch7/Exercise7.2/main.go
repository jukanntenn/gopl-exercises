package main

import (
	"bytes"
	"fmt"
	"io"
)

type WriterWrapper struct {
	w io.Writer
	n int64
}

func (w *WriterWrapper) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.n += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := WriterWrapper{
		w: w,
		n: 0,
	}
	return &wrapper, &wrapper.n
}

func main() {
	var buf *bytes.Buffer
	w, n := CountingWriter(buf)
	fmt.Fprint(w, []byte("Hello"))
	fmt.Println(*n)

	fmt.Fprint(w, []byte(" World!"))
	fmt.Println(*n)
}
