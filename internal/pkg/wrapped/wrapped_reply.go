package wrapped

import (
	"encoding/json"
	"io"
)

// A WrappedReply represents the response that should be returned when the
// client has requested data be wrapped and meta information included
type WrappedReply struct {
	Message string      `json:"message"`
	Self    string      `json:"self"`
	Data    interface{} `json:"data,omitempty"`
	Hints   string      `json:"hints,omitempty"`
}

// DecodeWrappedReplyFromReader decodes JSON from a Reader into a
// WrappedReply
func DecodeWrappedReplyFromReader(r io.Reader) (WrappedReply, error) {
	var wr WrappedReply
	err := json.NewDecoder(r).Decode(&wr)
	return wr, err
}
