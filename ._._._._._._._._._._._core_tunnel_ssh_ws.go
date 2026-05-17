package tunnel

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "github.com/sdcvvvhhyuu-wq/sliptunnel/pqc"
)

type SSHOverWebSocket struct{ inner *SSHTunnel }

func NewSSHOverWebSocket(inner *SSHTunnel) *SSHOverWebSocket { return &SSHOverWebSocket{inner} }

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (s *SSHOverWebSocket) Start() error {
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil { return }
        defer conn.Close()
        log.Println("WebSocket SSH tunnel established")
    })
    return nil
}
func (s *SSHOverWebSocket) Stop() error { return nil }
func (s *SSHOverWebSocket) SetQuantumSession(_ *pqc.QuantumSession) {}
