// +build test

package testutils

import "strings"

type MemoryWriter struct {
	Data []string
}

func NewMemoryWriter() *MemoryWriter {
	return &MemoryWriter{
		Data: make([]string, 0),
	}
}

func (w *MemoryWriter) Write(p []byte) (int, error) {
	msg := string(p)

	for _, m := range strings.Split(msg, "\n") {
		if len(m) > 0 {
			w.Data = append(w.Data, m)
		}
	}

	return len(p), nil
}

func (w *MemoryWriter) Reset() {
	w.Data = make([]string, 0)
}
