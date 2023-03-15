package server

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"

	types "sqlbrite-server/gen-go/message_protocol"
	"sqlbrite-server/pkg/db"
)

type Server struct {
	AuthTokensPath string
	DB             *db.DBController
}

func (s *Server) HandleConnection(conn *websocket.Conn) (err error) {
	var writeMutex sync.Mutex

	for {
		messageType, message, readErr := conn.ReadMessage()
		if readErr != nil {
			return readErr
		}
		if messageType != websocket.TextMessage {
			return fmt.Errorf("invalid message type, must use text")
		}

		go func() {
			res := s.processMessage(message)

			writeMutex.Lock()
			conn.WriteMessage(websocket.TextMessage, res)
			writeMutex.Unlock()
		}()
	}
}

func (s *Server) processMessage(message []byte) (response []byte) {
	var operationMessage types.OperationMessage
	if err := json.Unmarshal(message, &operationMessage); err != nil {
		panic(err)
	}

	switch operationMessage.OperationType {
	case types.OperationType_EACH:
	case types.OperationType_ALL:
	case types.OperationType_GET:
		s.DB.Query(operationMessage.Sql)
	default:
		return
	}

	return
}
