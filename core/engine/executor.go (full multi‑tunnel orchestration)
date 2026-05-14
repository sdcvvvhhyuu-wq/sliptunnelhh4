package engine

import (
    "log"
    "math/rand"
    "time"

    "github.com/sdcvvvhhyuu-wq/argotunnel/capabilities"
    "github.com/sdcvvvhhyuu-wq/argotunnel/domain_discovery"
    "github.com/sdcvvvhhyuu-wq/argotunnel/ip_scanner"
    "github.com/sdcvvvhhyuu-wq/argotunnel/obfs"
    "github.com/sdcvvvhhyuu-wq/argotunnel/pqc"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/ai_morph"
    "github.com/sdcvvvhhyuu-wq/argotunnel/active_shield"
    "github.com/sdcvvvhhyuu-wq/argotunnel/dynamic_orchestra"
    "github.com/sdcvvvhhyuu-wq/argotunnel/dpi_analyzer"
    "github.com/sdcvvvhhyuu-wq/argotunnel/rl_agent"
    "github.com/sdcvvvhhyuu-wq/argotunnel/gan_generator"
    "github.com/sdcvvvhhyuu-wq/argotunnel/quic_masq"
)

type Executor struct {
    caps           capabilities.CapabilitySet
    primary        tunnel.Tunnel
    fallbacks      []tunnel.Tunnel
}

func NewExecutor(c capabilities.CapabilitySet) *Executor {
    e := &Executor{caps: c}
    e.setupTunnels()
    return e
}

func (e *Executor) setupTunnels() {
    // Primary based on selected transport
    switch e.caps.Transport {
    case "wireguard":
        e.primary = tunnel.NewWireGuardTunnel()
    case "shadowsocks":
        e.primary = tunnel.NewShadowsocksTunnel()
    case "dns", "noizdns", "vaydns", "slipstream":
        e.primary = tunnel.NewDNSTunnel(e.caps.Transport)
    case "ssh":
        e.primary = tunnel.NewSSHTunnel()
    case "naiveproxy":
        e.primary = tunnel.NewNaiveProxy()
    case "doh":
        e.primary = tunnel.NewDOHProxy()
    case "tor":
        e.primary = tunnel.NewTorTunnel()
    default:
        e.primary = tunnel.NewWireGuardTunnel()
    }
    // Always add fallback tunnels
    e.fallbacks = []tunnel.Tunnel{
        tunnel.NewDNSTunnel("dns"),
        tunnel.NewICMPTunnel(),
        tunnel.NewSSHTunnel(),
    }
}

func (e *Executor) Start() error {
    domain_discovery.StartDiscovery()
    ip_scanner.StartScanning()

    // Enable fronting and obfuscation
    obfs.EnableAdaptiveFronting()
    if e.caps.UseUTLS { obfs.EnableUTLS(e.caps.UTLSProfile) }
    if e.caps.FragSize > 0 { obfs.EnableFragmentation(e.caps.FragSize) }

    // PQC key agreement
    if e.caps.PQCLevel > 0 {
        sess, err := pqc.NewQuantumSession(e.caps.PQCLevel)
        if err == nil {
            e.primary.SetQuantumSession(sess)
            for _, t := range e.fallbacks {
                t.SetQuantumSession(sess)
            }
            sess.StartKeyRollover()
        }
    }

    // AI traffic morph
    if e.caps.AIMorphEnabled { ai_morph.EnableTrafficMorphing() }
    active_shield.StartProbeShield()
    gan_generator.EnableGANMorphing(20 * time.Second)
    quic_masq.EnableQUICWrapper(e.primary)

    // Orchestrator with all tunnels
    all := append([]tunnel.Tunnel{e.primary}, e.fallbacks...)
    orch := dynamic_orchestra.NewOrchestrator(all...)
    orch.Start()

    // DPI analyzer + RL agent
    if e.caps.DPIAdaptive {
        dpi := dpi_analyzer.NewAnalyzer()
        go func() {
            for {
                time.Sleep(5 * time.Second)
                profile := rl_agent.GetRecommendedProfile()
                dpi.SetProfile(profile)
                ai_morph.SetActiveProfile(profile)
            }
        }()
    }

    log.Printf("SlipNet started with primary %s + %d fallback tunnels", e.caps.Transport, len(e.fallbacks))
    return e.primary.Start()
}
