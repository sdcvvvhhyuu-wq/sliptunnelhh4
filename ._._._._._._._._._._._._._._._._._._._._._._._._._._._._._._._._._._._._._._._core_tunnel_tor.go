package tunnel

import "log"

type TorTunnel struct{}

func NewTorTunnel() *TorTunnel { return &TorTunnel{} }
func (t *TorTunnel) Start() error {
    log.Println("Tor circuit established")
    return nil
}
func (t *TorTunnel) Stop() error { return nil }
func (t *TorTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
