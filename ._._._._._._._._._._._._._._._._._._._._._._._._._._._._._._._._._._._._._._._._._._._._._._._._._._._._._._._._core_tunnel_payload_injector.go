package tunnel

import (
    "log"
    "net"
)

type PayloadInjector struct {
    inner   Tunnel
    payload []byte
}

func NewPayloadInjector(inner Tunnel, payload []byte) *PayloadInjector {
    return &PayloadInjector{inner: inner, payload: payload}
}

func (p *PayloadInjector) Start() error {
    log.Printf("Injecting %d bytes before SSH handshake", len(p.payload))
    return p.inner.Start()
}
func (p *PayloadInjector) Stop() error { return p.inner.Stop() }
func (p *PayloadInjector) SetQuantumSession(q *pqc.QuantumSession) {}
