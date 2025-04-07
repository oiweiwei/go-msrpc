package internal

import (
	"context"
	"fmt"
	"io"
	"time"
)

type RawConn interface {
	io.ReadWriteCloser
}

// BufferedConn is a raw connection wrapper to optimize the short
// read header operations.
type BufferedConn struct {
	// The wrapped raw connection.
	RawConn
	// cur for current position, total is buffer start position.
	cur, total []byte
}

func (conn *BufferedConn) Resized(sz int) *BufferedConn {
	if sz > len(conn.total) {
		conn.total = append(conn.total, make([]byte, sz-len(conn.total))...)
	}
	conn.cur, conn.total = nil, conn.total[:sz]
	return conn
}

// NewBufferedConn function returns the new buffered connection.
func NewBufferedConn(cc RawConn, sz int) *BufferedConn {
	return &BufferedConn{RawConn: cc, cur: nil, total: make([]byte, sz)}
}

func (conn *BufferedConn) MustWithin(ctx context.Context, b []byte, f func([]byte) (int, error), timeout time.Duration) error {

	result := make(chan error)
	go func() {
		result <- conn.Must(b, f)
	}()

	timer := NewTimer()
	defer timer.Stop()

	select {
	case <-timer.After(timeout):
		timer.Clear()
		return context.DeadlineExceeded
	case <-ctx.Done():
		return ctx.Err()
	case err := <-result:
		return err
	}
}

func (conn *BufferedConn) Must(b []byte, f func([]byte) (int, error)) error {
	n, err := f(b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return fmt.Errorf("must: trailing bytes")
	}
	return nil
}

func (conn *BufferedConn) Write(b []byte) (int, error) {

	want := len(b)

	for len(b) > 0 {
		n, err := conn.RawConn.Write(b)
		if err != nil {
			return n, err
		}
		b = b[n:]
	}

	return want, nil
}

// Read function reads the data into buffer `b` and optionally fetches
// the next data block.
func (conn *BufferedConn) Read(b []byte) (int, error) {

	want := len(b)

	if len(b) <= len(conn.cur) {
		// copy current to bytes and advance conn.cur to b lenght.
		conn.cur = conn.cur[copy(b, conn.cur):]
		return want, nil
	}

	if len(conn.cur) > 0 {
		// copy all bytes to b.
		b = b[copy(b, conn.cur):]
	}

	for len(b) > 0 {
		// reset current buffer.
		conn.cur = conn.total

		// read available data.
		n, err := conn.RawConn.Read(conn.cur)
		if err != nil {
			return n, err
		}

		// limit current buffer to read chunk.
		conn.cur = conn.cur[:n]
		// copy data to buffer.
		n = copy(b, conn.cur)
		// advance the buffer.
		conn.cur, b = conn.cur[n:], b[n:]
	}

	return want, nil
}
