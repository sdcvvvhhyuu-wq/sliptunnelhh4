package tunnel

import (
    "fmt"
    "net"
    "net/url"
    "github.com/sdcvvvhhyuu-wq/sliptunnel/pqc"
)

type SSHOverHTTPConnect struct {
    inner    *SSHTunnel
    proxyURL *url.URL
}
func NewSSHOverHTTPConnect(inner *SSHTunnel, proxy string) *SSHOverHTTPConnect {
    u, _ := url.Parse(proxy)
    return &SSHOverHTTPConnect{inner, u}
}
func (s *SSHOverHTTPConnect) Start() error {
    conn, err := net.Dial("tcp", s.proxyURL.Host)
    if err != nil { return err }
    fmt.Fprintf(conn, "CONNECT %s HTTP/1.1\r\nHost: %s\r\n\r\n", "target:22", s.proxyURL.Host)
    return s.inner.Start()
}
func (s *SSHOverHTTPConnect) Stop() error { return s.inner.Stop() }
func (s *SSHOverHTTPConnect) SetQuantumSession(_ *pqc.QuantumSession) {}
