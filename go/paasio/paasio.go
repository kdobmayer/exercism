package paasio

import (
	"io"
	"sync"
)

type Wrapper struct {
	r io.Reader
	w io.Writer

	totalBytes int64
	nops       int
	mu         sync.Mutex
}

func (w *Wrapper) count(n int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.totalBytes += int64(n)
	w.nops++
}

func (w *Wrapper) Read(p []byte) (int, error) {
	b, err := w.r.Read(p)
	w.count(b)
	return b, err
}

func (w *Wrapper) ReadCount() (n int64, nops int) {
	return w.totalBytes, w.nops
}

func (w *Wrapper) Write(p []byte) (int, error) {
	b, err := w.w.Write(p)
	w.count(b)
	return b, err
}

func (w *Wrapper) WriteCount() (n int64, nops int) {
	return w.totalBytes, w.nops
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &Wrapper{r: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &Wrapper{w: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &Wrapper{r: rw, w: rw}
}