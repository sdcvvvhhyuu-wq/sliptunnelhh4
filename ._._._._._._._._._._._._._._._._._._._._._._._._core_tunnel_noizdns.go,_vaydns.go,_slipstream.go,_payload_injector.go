package tunnel
import ("log"; "github.com/sdcvvvhhyuu-wq/sliptunnel/pqc")
type NoizDNSTunnel struct{}
func NewNoizDNSTunnel() *NoizDNSTunnel { return &NoizDNSTunnel{} }
func (n *NoizDNSTunnel) Start() error { log.Println("NoizDNS started"); return nil }
func (n *NoizDNSTunnel) Stop() error { return nil }
func (n *NoizDNSTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}

type VayDNSTunnel struct{}
func NewVayDNSTunnel() *VayDNSTunnel { return &VayDNSTunnel{} }
func (v *VayDNSTunnel) Start() error { log.Println("VayDNS started"); return nil }
func (v *VayDNSTunnel) Stop() error { return nil }
func (v *VayDNSTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}

type SlipStreamTunnel struct{}
func NewSlipStreamTunnel() *SlipStreamTunnel { return &SlipStreamTunnel{} }
func (s *SlipStreamTunnel) Start() error { log.Println("SlipStream started"); return nil }
func (s *SlipStreamTunnel) Stop() error { return nil }
func (s *SlipStreamTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}

type PayloadInjector struct {
    inner   Tunnel
    payload []byte
}
func NewPayloadInjector(inner Tunnel, payload []byte) *PayloadInjector { return &PayloadInjector{inner, payload} }
func (p *PayloadInjector) Start() error { log.Printf("Payload inject: %d bytes", len(p.payload)); return p.inner.Start() }
func (p *PayloadInjector) Stop() error { return p.inner.Stop() }
func (p *PayloadInjector) SetQuantumSession(q *pqc.QuantumSession) {}
