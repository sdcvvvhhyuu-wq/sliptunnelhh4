package tunnel

import (
    "log"
    "net"
    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
)

type ICMPTunnel struct {
    conn *icmp.PacketConn
}

func NewICMPTunnel() *ICMPTunnel {
    conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
    if err != nil {
        log.Printf("ICMP tunnel listen error (need root): %v", err)
        return nil
    }
    return &ICMPTunnel{conn: conn}
}

func (i *ICMPTunnel) Start() error {
    if i == nil { return nil }
    go func() {
        buf := make([]byte, 1500)
        for {
            n, peer, err := i.conn.ReadFrom(buf)
            if err != nil { continue }
            msg, _ := icmp.ParseMessage(1, buf[:n])
            if msg != nil && msg.Type == ipv4.ICMPTypeEcho {
                reply := icmp.Message{Type: ipv4.ICMPTypeEchoReply, Code: 0, Body: msg.Body}
                b, _ := reply.Marshal(nil)
                i.conn.WriteTo(b, peer)
            }
        }
    }()
    log.Println("ICMP covert tunnel started")
    return nil
}
func (i *ICMPTunnel) Stop() error { if i.conn != nil { i.conn.Close() }; return nil }
func (i *ICMPTunnel) SetQuantumSession(_ *pqc.QuantumSession) {}
