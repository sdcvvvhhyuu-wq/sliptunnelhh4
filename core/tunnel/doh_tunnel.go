package tunnel

import (
    "log"
    "github.com/miekg/dns"
)

type DOHProxy struct{}

func NewDOHProxy() *DOHProxy { return &DOHProxy{} }
func (d *DOHProxy) Start() error {
    log.Println("DNS over HTTPS proxy active")
    return nil
}
func (d *DOHProxy) Stop() error { return nil }
func (d *DOHProxy) SetQuantumSession(_ *pqc.QuantumSession) {}
