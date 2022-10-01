package linkedlist

import "fmt"

type LinkedListError struct {
	msg       string
	operation string
}

func (err *LinkedListError) Error() string {
	return fmt.Sprintf("%s: %s", err.operation, err.msg)
}

type BytesLinkedListItem struct {
	msg []byte
	ll  *BytesLinkedListItem
}

func (ll *BytesLinkedListItem) Add(msg []byte) (*BytesLinkedListItem, error) {
	if ll == nil {
		return &BytesLinkedListItem{
			msg: msg,
			ll:  nil,
		}, nil
	}

	// defensive add
	if ll.ll != nil {
		return nil, &LinkedListError{
			msg:       "Already an item present",
			operation: "Add",
		}
	}
	return &BytesLinkedListItem{msg: msg, ll: ll}, nil
}

func (ll *BytesLinkedListItem) Length() int {
	if ll == nil {
		return 0
	}
	var i int
	for i = 1; ll.ll != nil; i++ {
		ll = ll.ll
	}
	return i
}
