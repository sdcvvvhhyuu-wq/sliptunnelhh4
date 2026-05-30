package obfs

import (
    "log"
    "math/rand"
    "sync"
    "time"
    "github.com/sdcvvvhhyuu-wq/argotunnel/domain_discovery"
    "github.com/sdcvvvhhyuu-wq/argotunnel/ip_scanner"
)

var (
    currentFront string
    mu           sync.RWMutex
    baseHosts    = []string{"cloudflare.com", "google.com", "amazon.com", "fastly.com", "akamai.com"}
    pool         []string
)

func init() {
    pool = baseHosts
    go dynamicUpdate()
}

func SetFrontingHost(host string) {
    mu.Lock(); currentFront = host; mu.Unlock()
}

func GetFrontingHost() string {
    mu.RLock(); defer mu.RUnlock()
    if currentFront != "" { return currentFront }
    if len(pool) > 0 { return pool[rand.Intn(len(pool))] }
    return "cloudflare.com"
}

func EnableAdaptiveFronting() {
    log.Println("Adaptive fronting enabled")
}

func dynamicUpdate() {
    for {
        time.Sleep(1 * time.Minute)
        newPool := baseHosts
        newPool = append(newPool, domain_discovery.GetDiscoveredDomains()...)
        for _, ip := range ip_scanner.GetCleanIPs() {
            newPool = append(newPool, ip.String())
        }
        mu.Lock()
        pool = newPool
        if len(pool) > 0 { currentFront = pool[rand.Intn(len(pool))] }
        mu.Unlock()
    }
}
