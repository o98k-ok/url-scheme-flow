package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type WebSkt struct {
	conn  *websocket.Conn
	mType string
}

func NewCommandSocket(url string) (*WebSkt, error) {
	con, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &WebSkt{
		conn:  con,
		mType: "command",
	}, nil
}

func (w *WebSkt) Do(msg string) error {
	simpleReq := map[string]string{
		w.mType: msg,
	}

	dat, err := json.Marshal(simpleReq)
	if err != nil {
		return err
	}

	return w.conn.WriteMessage(websocket.TextMessage, dat)
}

func (w *WebSkt) Close() {
	w.conn.Close()
}
