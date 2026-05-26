package dynamic_orchestra

import (
    "log"
    "math/rand"
    "time"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tunnel"
)

type Orchestrator struct {
    tunnels []tunnel.Tunnel
    active  int
}

func NewOrchestrator(tuns ...tunnel.Tunnel) *Orchestrator {
    return &Orchestrator{tunnels: tuns}
}

func (o *Orchestrator) Start() {
    o.active = 0
    o.tunnels[o.active].Start()
    go func() {
        for {
            time.Sleep(30 * time.Second)
            // simulate switching based on random loss
            if rand.Float32() > 0.8 {
                next := (o.active + 1) % len(o.tunnels)
                log.Printf("Orchestrator: switching from %d to %d", o.active, next)
                o.tunnels[o.active].Stop()
                o.tunnels[next].Start()
                o.active = next
            }
        }
    }()
}
