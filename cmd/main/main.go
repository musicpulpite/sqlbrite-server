package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"sqlbrite-server/pkg/db"
	"sqlbrite-server/pkg/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func parseCMDArgs() (port int, dbPath string, authTokensPath string) {
	port = *flag.Int("port", 8080, "port for http server")
	dbPath = *flag.String("db-path", "sqlite3.db", "path to sqlite3 db")
	authTokensPath = *flag.String("auth-tokens-path", ".auth/auth_tokens.yml", "path to client authentication tokens")
	return
}

func main() {
	port, dbPath, authTokensPath := parseCMDArgs()

	dbController := db.DBController{
		Path: dbPath,
	}

	if err := dbController.Init(); err != nil {
		panic(err)
	}

	server := server.Server{
		AuthTokensPath: authTokensPath,
		DB:             &dbController,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		if err := server.HandleConnection(conn); err != nil {
			fmt.Errorf("Error handling ws connection: %s", err)
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
