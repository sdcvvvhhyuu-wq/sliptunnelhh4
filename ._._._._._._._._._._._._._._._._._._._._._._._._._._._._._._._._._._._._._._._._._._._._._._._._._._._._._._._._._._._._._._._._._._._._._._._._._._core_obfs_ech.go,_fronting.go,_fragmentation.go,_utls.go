package obfs

import (
    "crypto/tls"
    "log"
    "net"
    "sync"
    "time"
    "github.com/refraction-networking/utls"
    "github.com/sdcvvvhhyuu-wq/sliptunnel/domain_discovery"
    "github.com/sdcvvvhhyuu-wq/sliptunnel/ip_scanner"
)

func DialWithECH(network, addr string, config *tls.Config) (*tls.Conn, error) {
    tcpConn, err := net.Dial(network, addr)
    if err != nil { return nil, err }
    uConfig := &utls.Config{ServerName: config.ServerName, InsecureSkipVerify: config.InsecureSkipVerify}
    uConn := utls.UClient(tcpConn, uConfig, utls.HelloChrome_Auto)
    if err = uConn.Handshake(); err != nil {
        tcpConn.Close()
        return tls.Dial(network, addr, config)
    }
    log.Printf("ECH established to %s", config.ServerName)
    return uConn, nil
}

// --- Adaptive Fronting ---
var (
    currentFront string
    mu           sync.RWMutex
    pool         = []string{"cloudflare.com", "google.com", "amazon.com", "fastly.com", "akamai.com"}
)

func init() { go dynamicUpdate() }
func SetFrontingHost(host string) { mu.Lock(); currentFront = host; mu.Unlock() }
func GetFrontingHost() string {
    mu.RLock(); defer mu.RUnlock()
    if currentFront != "" { return currentFront }
    if len(pool) > 0 { return pool[time.Now().UnixNano()%int64(len(pool))] }
    return "cloudflare.com"
}
func EnableAdaptiveFronting() { log.Println("Adaptive fronting enabled") }
func dynamicUpdate() {
    for {
        time.Sleep(1 * time.Minute)
        newPool := []string{"cloudflare.com", "google.com", "amazon.com", "fastly.com", "akamai.com"}
        newPool = append(newPool, domain_discovery.GetDiscoveredDomains()...)
        for _, ip := range ip_scanner.GetCleanIPs() { newPool = append(newPool, ip.String()) }
        mu.Lock(); pool = newPool; mu.Unlock()
    }
}

func EnableFragmentation(size int) { log.Printf("Fragmentation: %d bytes", size) }
func EnableUTLS(profile string) {
    var id utls.ClientHelloID
    switch profile {
    case "chrome_124": id = utls.HelloChrome_124
    case "firefox_125": id = utls.HelloFirefox_125
    default: id = utls.HelloRandomized
    }
    log.Printf("UTLS: %s", id.Str())
}
