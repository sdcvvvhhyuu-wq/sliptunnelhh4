package tunnel

import (
    "crypto/tls"
    "net"
)

type SSHOverTLS struct {
    inner Tunnel
    sni   string
}

func NewSSHOverTLS(inner Tunnel, sni string) *SSHOverTLS { return &SSHOverTLS{inner, sni} }
func (s *SSHOverTLS) Start() error {
    conn, err := tls.Dial("tcp", "server:443", &tls.Config{ServerName: s.sni})
    if err != nil { return err }
    // Tunnel SSH connection over this conn
    return nil
}
func (s *SSHOverTLS) Stop() error { return nil }
func (s *SSHOverTLS) SetQuantumSession(q *pqc.QuantumSession) {}
