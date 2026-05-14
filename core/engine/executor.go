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
    "github.com/sdcvvvhhyuu-wq/argotunnel/dns_tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/icmp_tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/dpi_analyzer"
    "github.com/sdcvvvhhyuu-wq/argotunnel/rl_agent"
    "github.com/sdcvvvhhyuu-wq/argotunnel/gan_generator"
    "github.com/sdcvvvhhyuu-wq/argotunnel/quic_masq"
)

type Executor struct {
    caps          capabilities.CapabilitySet
    frontingHosts []string
    cleanIPs      []string
}

func NewExecutor(c capabilities.CapabilitySet) *Executor {
    return &Executor{caps: c}
}

func (e *Executor) Start() error {
    domain_discovery.StartDiscovery()
    ip_scanner.StartScanning()

    baseHosts := []string{
        "cloudflare.com", "google.com", "amazon.com", "fastly.com",
        "akamai.com", "microsoft.com", "azure.com", "facebook.com",
        "twitter.com", "dropbox.com", "godaddy.com", "namecheap.com",
        "digitalocean.com", "linode.com", "vultr.com", "cdn77.com",
        "stackpath.com", "cachefly.com", "bunnycdn.com", "gcorelabs.com",
    }
    e.frontingHosts = baseHosts

    go func() {
        for {
            time.Sleep(30 * time.Second)
            discovered := domain_discovery.GetDiscoveredDomains()
            e.frontingHosts = append(baseHosts, discovered...)
            e.cleanIPs = nil
            for _, ip := range ip_scanner.GetCleanIPs() {
                e.cleanIPs = append(e.cleanIPs, ip.String())
            }
            choices := append(e.frontingHosts, e.cleanIPs...)
            if len(choices) > 0 {
                chosen := choices[rand.Intn(len(choices))]
                obfs.SetFrontingHost(chosen)
            }
        }
    }()

    log.Printf("⚡ ArgoTunnel booting with %d AI/Quantum capabilities", len(e.caps.ActiveIDs))

    var primary, secondary tunnel.Tunnel
    if e.caps.Transport == "shadowsocks" {
        primary = tunnel.NewShadowsocksTunnel()
        secondary = tunnel.NewWireGuardTunnel()
    } else {
        primary = tunnel.NewWireGuardTunnel()
        secondary = tunnel.NewShadowsocksTunnel()
    }

    if e.caps.UseUTLS {
        obfs.EnableUTLS(e.caps.UTLSProfile)
    }
    if e.caps.FragSize > 0 {
        obfs.EnableFragmentation(e.caps.FragSize)
    }
    if e.caps.DPIAdaptive {
        dpiA := dpi_analyzer.NewAnalyzer()
        go func() {
            for {
                time.Sleep(5 * time.Second)
                bestProfile := rl_agent.GetRecommendedProfile()
                dpiA.SetProfile(bestProfile)
                ai_morph.SetActiveProfile(bestProfile)
            }
        }()
    }
    if e.caps.AIMorphEnabled {
        ai_morph.EnableTrafficMorphing()
    }
    active_shield.StartProbeShield()
    if e.caps.PQCLevel > 0 {
        sess, err := pqc.NewQuantumSession(e.caps.PQCLevel)
        if err == nil {
            primary.SetQuantumSession(sess)
            secondary.SetQuantumSession(sess)
            sess.RegisterTunnel(primary)
            sess.StartKeyRollover()
        }
    }
    if e.caps.UseDNSTunnel {
        dns_tunnel.NewDNSTunnel("tunnel.example.com", nil).Start()
    }
    if e.caps.UseICMPTunnel {
        icmp_tunnel.NewICMPTunnel().Start()
    }
    gan_generator.EnableGANMorphing(20 * time.Second)
    quic_masq.EnableQUICWrapper(primary)
    dynamic_orchestra.NewOrchestrator(primary, secondary).Start()

    return primary.Start()
}
