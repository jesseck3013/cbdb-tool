package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

const INTERNAL_SERVER_ERROR = "Internal Server Error"

type ErrMsg struct {
	Msg string `json:"error"`
}

func NewErrMsg(msg string) ErrMsg {
	return ErrMsg{
		Msg: msg,
	}
}

func writeMsgJson(w io.Writer, msg string) {
	enc := json.NewEncoder(w)
	err := enc.Encode(NewErrMsg(msg))
	if err != nil {
		m := fmt.Sprintf("Failed to encode error msg into json: %v", err)
		log.Println(m)
	}
}
