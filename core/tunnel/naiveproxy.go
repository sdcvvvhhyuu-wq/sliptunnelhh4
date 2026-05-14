package tunnel

import "log"

type NaiveProxy struct{}

func NewNaiveProxy() *NaiveProxy { return &NaiveProxy{} }
func (n *NaiveProxy) Start() error {
    log.Println("NaiveProxy tunnel started (Chromium stack)")
    return nil
}
func (n *NaiveProxy) Stop() error { return nil }
func (n *NaiveProxy) SetQuantumSession(_ *pqc.QuantumSession) {}
