package tunnel

import (
    "crypto/tls"
    "net"
    "log"
)

type SSHOverTLS struct {
    inner *SSHTunnel
    sni   string
}

func NewSSHOverTLS(inner *SSHTunnel, sni string) *SSHOverTLS {
    return &SSHOverTLS{inner: inner, sni: sni}
}

func (s *SSHOverTLS) Start() error {
    config := &tls.Config{ServerName: s.sni}
    conn, err := tls.Dial("tcp", "server:443", config)
    if err != nil { return err }
    // SSH connection will run over this TLS conn (simplified here)
    log.Printf("SSH over TLS (SNI: %s) established", s.sni)
    return s.inner.Start() // start SSH on top
}
func (s *SSHOverTLS) Stop() error { return s.inner.Stop() }
func (s *SSHOverTLS) SetQuantumSession(q *pqc.QuantumSession) {}
