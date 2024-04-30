package wmio

import "bytes"

// heaps.go module contains data structures for encoding/decoding
// heap references, representing the heaps as a stack of individual
// heaps, for illustration, while decoding class object, we allocate a
// new heap and within the decoding process we allocate another heap
// for the methods. When we are done with methods encoding, we pop the
// heap and continue to work on class heap.

// Heaps represent the stack of heaps, top-level (last) element
// represents the current heap.
type Heaps struct {
	hs []*Heap
}

// SetHeap function sets the current (last) heap element buffer
// to the bytes `h`.
func (r *Heaps) SetHeap(h []byte) {
	r.hs[len(r.hs)-1].buf = bytes.NewBuffer(h)
}

// HeapBuffer function returns the top-level (last) heap
// bytes.Buffer.
func (r *Heaps) HeapBuffer() *bytes.Buffer {
	return r.hs[len(r.hs)-1].buf
}

// Heap function returns the top-level (last) heap bytes.
func (r *Heaps) Heap() []byte {
	return r.HeapBuffer().Bytes()
}

// Head function returns the list of references on the top-level
// (last) heap.
func (r *Heaps) Head() []*HeapRef {
	return r.hs[len(r.hs)-1].refs
}

// Push function allocates a new heap on the heap stack.
func (r *Heaps) Push() {
	r.hs = append(r.hs, &Heap{buf: bytes.NewBuffer(nil)})
}

// Pop function removes the top-level (last) heap from the
// heap stack.
func (r *Heaps) Pop() {
	r.hs = r.hs[:len(r.hs)-1]
}

// Append function appends the reference to the top-level (last)
// heap.
func (r *Heaps) Append(ref *HeapRef) {
	r.hs[len(r.hs)-1].refs = append(r.hs[len(r.hs)-1].refs, ref)
}

// Truncate function removes the list of references from the top-level
// (last) heap but leaves the heap as-is.
func (r *Heaps) Truncate() []*HeapRef {
	head := r.hs[len(r.hs)-1].refs
	r.hs[len(r.hs)-1].refs = nil
	return head
}

// The heap stack element.
type Heap struct {
	// The encoding buffer.
	buf *bytes.Buffer
	// The list of references to encode.
	refs []*HeapRef
}

// The heap reference.
type HeapRef struct {
	// The offset on the heap.
	Offset uint32
	// The value to encode/decode.
	Value any
	// The debug information associated with the
	// reference.
	Debug []string
}
