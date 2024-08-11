package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg) // json.stringify
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"` // look at it in the `method` field in the json data
}

func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, errors.New("did not find separator")
	}

	// Content-Length: <number>
	contenetLengthBytes := header[len("Content-Length: "):]
	contenetLength, err := strconv.Atoi(string(contenetLengthBytes))
	if err != nil {
		return 0, err
	}

	// TODO: we will use this later
	_ = content

  var BaseMessage
    if err := json.Unmarhall(content[:contenetLength], &BaseMessage); err != nil {
    return 0, err
  }

	return contenetLength, nil
}
