package tunnel

import (
    "golang.org/x/crypto/ssh"
)

type SSHTunnel struct {
    client *ssh.Client
}

func NewSSHTunnel() *SSHTunnel { return &SSHTunnel{} }

func (s *SSHTunnel) Start() error {
    config := &ssh.ClientConfig{
        User: "root",
        Auth: []ssh.AuthMethod{ssh.Password("password")},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", "example.com:22", config)
    if err != nil { return err }
    s.client = client
    return nil
}
func (s *SSHTunnel) Stop() error {
    if s.client != nil { return s.client.Close() }
    return nil
}
func (s *SSHTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
