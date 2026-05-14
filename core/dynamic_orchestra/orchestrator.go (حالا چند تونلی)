package dynamic_orchestra

import (
    "log"
    "math/rand"
    "time"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tunnel"
)

type Orchestrator struct {
    tunnels []tunnel.Tunnel
    activeIndex int
}

func NewOrchestrator(tunnels ...tunnel.Tunnel) *Orchestrator {
    return &Orchestrator{tunnels: tunnels}
}

func (o *Orchestrator) Start() {
    o.activeIndex = 0
    o.tunnels[o.activeIndex].Start()
    go func() {
        for {
            time.Sleep(30 * time.Second)
            loss := rand.Float32()
            if loss > 0.3 || !isAlive(o.tunnels[o.activeIndex]) {
                next := (o.activeIndex + 1) % len(o.tunnels)
                log.Printf("Orchestrator: switching from tunnel %d to %d", o.activeIndex, next)
                o.tunnels[o.activeIndex].Stop()
                o.tunnels[next].Start()
                o.activeIndex = next
            }
        }
    }()
}
func isAlive(t tunnel.Tunnel) bool { return true } // placeholder
