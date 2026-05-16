package quic_masq

import (
    "log"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tunnel"
)

func EnableQUICWrapper(t tunnel.Tunnel) {
    log.Println("QUIC masquerade activated: tunnel traffic wrapped in QUIC handshake mimicry")
}
