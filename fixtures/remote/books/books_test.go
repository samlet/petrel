package books

import (
	"os"
	"testing"
	"zombiezen.com/go/capnproto2"
)

func TestBooks(t *testing.T) {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	book, err := NewRootBook(seg)
	if err != nil {
		panic(err)
	}
	book.SetTitle("War and Peace")
	book.SetPageCount(1440)

	// Write the message to stdout.
	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		panic(err)
	}
}
