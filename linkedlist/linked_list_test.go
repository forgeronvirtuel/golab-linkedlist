package linkedlist

import (
	"bytes"
	"math/rand"
	"testing"
)

const dataSize = 1_000_000

func generateBytes() []byte {
	token := make([]byte, 4)
	rand.Read(token)
	return token
}

func TestStringLinkedListItem_Add(t *testing.T) {
	var ll *BytesLinkedListItem

	ll, err := ll.Add([]byte("A"))
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(ll.msg, []byte("A")) != 0 {
		t.Errorf("Expected message to be A, got %s", string(ll.msg))
	}

	ll, err = ll.Add([]byte("B"))
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(ll.msg, []byte("B")) != 0 {
		t.Errorf("Expected message to be B, got %s", string(ll.msg))
	}
}

func BenchmarkStringLinkedListItem_Add(b *testing.B) {
	var ll *BytesLinkedListItem
	var err error

	data := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = generateBytes()
	}

	ll, _ = ll.Add(data[0])

	b.StartTimer()
	for i := 1; i < dataSize; i++ {
		ll, err = ll.Add(data[i])
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
}

func TestBytesLinkedListItem_Length(t *testing.T) {
	var ll *BytesLinkedListItem
	if ll.Length() != 0 {
		t.Errorf("Length %d, expected 0", ll.Length())
	}

	ll, err := ll.Add(generateBytes())
	if err != nil {
		t.Fatal(err)
	}

	if ll.Length() != 1 {
		t.Errorf("Got length %d, expected 1", ll.Length())
	}

	ll, err = ll.Add(generateBytes())
	if err != nil {
		t.Fatal(err)
	}

	if ll.Length() != 2 {
		t.Errorf("Got length %d, expected 2", ll.Length())
	}
}
