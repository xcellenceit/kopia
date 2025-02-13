package gather

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGatherWriteBuffer(t *testing.T) {
	// reset for testing
	all := &chunkAllocator{
		chunkSize: 100,
	}

	w := NewWriteBuffer()
	w.alloc = all

	defer w.Close()

	w.Append([]byte("hello "))
	fmt.Fprintf(w, "world!")

	if got, want := w.ToByteSlice(), []byte("hello world!"); !bytes.Equal(got, want) {
		t.Errorf("invaldi bytes %v, want %v", string(got), string(want))
	}

	if got, want := len(w.inner.Slices), 1; got != want {
		t.Errorf("invalid number of slices %v, want %v", got, want)
	}

	w.Append(bytes.Repeat([]byte("x"), all.chunkSize))

	if got, want := w.Length(), all.chunkSize+12; got != want {
		t.Errorf("invalid length: %v, want %v", got, want)
	}

	// one more slice was allocated
	if got, want := len(w.inner.Slices), 2; got != want {
		t.Errorf("invalid number of slices %v, want %v", got, want)
	}

	// write to fill the remainder of 2nd slice
	w.Append(bytes.Repeat([]byte("x"), all.chunkSize-12))

	// still 2 slices
	if got, want := len(w.inner.Slices), 2; got != want {
		t.Errorf("invalid number of slices %v, want %v", got, want)
	}

	// one more byte allocates new slice
	w.Append([]byte("x"))

	// still 3 slices
	if got, want := len(w.inner.Slices), 3; got != want {
		t.Errorf("invalid number of slices %v, want %v", got, want)
	}

	w.Reset()
}

func TestGatherDefaultWriteBuffer(t *testing.T) {
	var w WriteBuffer

	// one more byte allocates new slice
	w.Append([]byte("x"))

	if got, want := len(w.inner.Slices), 1; got != want {
		t.Errorf("invalid number of slices %v, want %v", got, want)
	}
}

func TestGatherWriteBufferContig(t *testing.T) {
	var w WriteBuffer
	defer w.Close()

	// allocate more than contig allocator can provide
	theCap := contiguousAllocator.chunkSize + 10
	b := w.MakeContiguous(theCap)
	require.Equal(t, theCap, len(b))
	require.Equal(t, theCap, cap(b))
}
