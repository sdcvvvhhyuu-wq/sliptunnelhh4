package domain_discovery

import (
    "log"
    "math/rand"
    "net"
    "sync"
    "time"
)

var (
    baseCandidates = []string{
        "cloudflare.com","google.com","amazon.com","fastly.com",
        "akamai.com","microsoft.com","azure.com","facebook.com",
        "twitter.com","dropbox.com","godaddy.com","namecheap.com",
        "digitalocean.com","linode.com","vultr.com","cdn77.com",
        "stackpath.com","cachefly.com","bunnycdn.com","gcorelabs.com",
    }
    discoveredDomains []string
    mu                sync.RWMutex
)

func GetDiscoveredDomains() []string {
    mu.RLock(); defer mu.RUnlock()
    res := make([]string, len(discoveredDomains))
    copy(res, discoveredDomains)
    return res
}

func StartDiscovery() {
    log.Println("🌐 AI Domain Discovery started...")
    go func() {
        for {
            candidate := generateCandidate()
            if isReachable(candidate) {
                addDiscovered(candidate)
            }
            time.Sleep(60 * time.Second)
        }
    }()
}

func generateCandidate() string {
    base := baseCandidates[rand.Intn(len(baseCandidates))]
    prefixes := []string{"www","cdn","edge","static","assets","api","img","v"}
    return prefixes[rand.Intn(len(prefixes))] + "." + base
}

func isReachable(domain string) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(domain, "443"), 3*time.Second)
    if err != nil { return false }
    conn.Close()
    return true
}

func addDiscovered(domain string) {
    mu.Lock(); defer mu.Unlock()
    for _, d := range discoveredDomains {
        if d == domain { return }
    }
    discoveredDomains = append(discoveredDomains, domain)
    log.Printf("✅ New domain discovered: %s", domain)
}
