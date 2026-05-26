package tunnel

import (
    "net"
    "golang.org/x/crypto/ssh"
)

type SSHTunnel struct {
    payload []byte
}

func NewSSHTunnel() *SSHTunnel { return &SSHTunnel{} }
func (s *SSHTunnel) EnablePayloadInjection() { s.payload = []byte{0x16, 0x03, 0x01, 0x00} /* TLS handshake fake */ }
func (s *SSHTunnel) Start() error {
    config := &ssh.ClientConfig{User: "root", Auth: []ssh.AuthMethod{ssh.Password("pass")}}
    _, err := ssh.Dial("tcp", "example.com:22", config)
    if err != nil { return err }
    return nil
}
func (s *SSHTunnel) Stop() error { return nil }
func (s *SSHTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
