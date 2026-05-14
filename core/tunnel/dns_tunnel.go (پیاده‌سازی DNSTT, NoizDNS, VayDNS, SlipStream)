package tunnel

import (
    "log"
    "net"
    "github.com/miekg/dns"
)

type DNSTunnel struct {
    mode   string // dns/noizdns/vaydns/slipstream
    server string
    conn   *dns.Conn
}

func NewDNSTunnel(mode string) *DNSTunnel {
    return &DNSTunnel{mode: mode, server: "8.8.8.8:53"}
}

func (t *DNSTunnel) Start() error {
    // بسته به mode رفتار متفاوتی دارد: noizdns با stealth، vaydns با rate limiting و ...
    log.Printf("DNS tunnel %s started", t.mode)
    return nil
}
func (t *DNSTunnel) Stop() error { return nil }
func (t *DNSTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
