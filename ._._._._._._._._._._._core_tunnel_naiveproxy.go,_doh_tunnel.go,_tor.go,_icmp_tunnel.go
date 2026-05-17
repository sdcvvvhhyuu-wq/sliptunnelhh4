package tunnel
import ("log"; "github.com/sdcvvvhhyuu-wq/sliptunnel/pqc")
type NaiveProxy struct{}
func NewNaiveProxy() *NaiveProxy { return &NaiveProxy{} }
func (n *NaiveProxy) Start() error { log.Println("NaiveProxy started"); return nil }
func (n *NaiveProxy) Stop() error { return nil }
func (n *NaiveProxy) SetQuantumSession(_ *pqc.QuantumSession) {}

type DOHProxy struct{}
func NewDOHProxy() *DOHProxy { return &DOHProxy{} }
func (d *DOHProxy) Start() error { log.Println("DoH started"); return nil }
func (d *DOHProxy) Stop() error { return nil }
func (d *DOHProxy) SetQuantumSession(_ *pqc.QuantumSession) {}

type TorTunnel struct{}
func NewTorTunnel() *TorTunnel { return &TorTunnel{} }
func (t *TorTunnel) Start() error { log.Println("Tor started"); return nil }
func (t *TorTunnel) Stop() error { return nil }
func (t *TorTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}

type ICMPTunnel struct{}
func NewICMPTunnel() *ICMPTunnel { return &ICMPTunnel{} }
func (i *ICMPTunnel) Start() error { log.Println("ICMP covert started"); return nil }
func (i *ICMPTunnel) Stop() error { return nil }
func (i *ICMPTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
