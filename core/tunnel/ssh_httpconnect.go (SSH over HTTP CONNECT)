package tunnel

import (
    "fmt"
    "net"
    "net/http"
    "net/url"
    "log"
)

type SSHOverHTTPConnect struct {
    inner    *SSHTunnel
    proxyURL *url.URL
}

func NewSSHOverHTTPConnect(inner *SSHTunnel, proxy string) *SSHOverHTTPConnect {
    u, _ := url.Parse(proxy)
    return &SSHOverHTTPConnect{inner: inner, proxyURL: u}
}

func (s *SSHOverHTTPConnect) Start() error {
    // Dial to proxy and send CONNECT
    conn, err := net.Dial("tcp", s.proxyURL.Host)
    if err != nil { return err }
    fmt.Fprintf(conn, "CONNECT %s HTTP/1.1\r\nHost: %s\r\n\r\n", "demohost.com:22", s.proxyURL.Host)
    // Read response, then run SSH on conn
    log.Println("SSH over HTTP CONNECT started")
    return s.inner.Start()
}
func (s *SSHOverHTTPConnect) Stop() error { return s.inner.Stop() }
func (s *SSHOverHTTPConnect) SetQuantumSession(q *pqc.QuantumSession) {}
