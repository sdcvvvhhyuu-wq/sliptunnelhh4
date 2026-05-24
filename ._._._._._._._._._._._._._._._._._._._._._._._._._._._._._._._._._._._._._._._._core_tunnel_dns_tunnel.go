package tunnel

import "github.com/sdcvvvhhyuu-wq/sliptunnel/pqc"

type DNSTunnel struct{ mode string }
func NewDNSTunnel(mode string) *DNSTunnel { return &DNSTunnel{mode: mode} }
func (d *DNSTunnel) Start() error { return nil }
func (d *DNSTunnel) Stop() error  { return nil }
func (d *DNSTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
