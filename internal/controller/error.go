package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type errMsg struct {
	Msg string `json:"error"`
}

func newErrMsg(msg string) errMsg {
	return errMsg{
		Msg: msg,
	}
}

func writeMsgJson(w io.Writer, msg string) {
	enc := json.NewEncoder(w)
	err := enc.Encode(newErrMsg(msg))
	if err != nil {
		m := fmt.Sprintf("Failed to encode error msg into json: %v", err)
		log.Println(m)
	}
}
