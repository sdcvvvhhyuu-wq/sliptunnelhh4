package ip_scanner

import (
    "log"
    "net"
    "sync"
    "time"
)

var (
    cidrs = []string{
        "104.16.0.0/12", "172.64.0.0/13",
        "13.32.0.0/15", "13.224.0.0/14",
        "151.101.0.0/16", "23.32.0.0/11",
    }
    cleanIPs []net.IP
    mu       sync.RWMutex
)

func GetCleanIPs() []net.IP {
    mu.RLock(); defer mu.RUnlock()
    res := make([]net.IP, len(cleanIPs))
    copy(res, cleanIPs)
    return res
}

func StartScanning() {
    log.Println("🔍 AI IP Scanner active – probing CDN ranges for clean IPs...")
    go func() {
        for {
            scanCIDRs()
            time.Sleep(10 * time.Minute)
        }
    }()
}

func scanCIDRs() {
    for _, cidr := range cidrs {
        _, ipnet, _ := net.ParseCIDR(cidr)
        for i := 0; i < 15; i++ {
            ip := randomIP(ipnet)
            if ip == nil { continue }
            if isClean(ip) {
                mu.Lock()
                cleanIPs = append(cleanIPs, ip)
                mu.Unlock()
                log.Printf("✅ Clean IP found: %s", ip)
            }
        }
    }
    mu.Lock()
    unique := map[string]net.IP{}
    for _, ip := range cleanIPs { unique[ip.String()] = ip }
    cleanIPs = nil
    for _, ip := range unique { cleanIPs = append(cleanIPs, ip) }
    mu.Unlock()
}

func randomIP(ipnet *net.IPNet) net.IP {
    ip := ipnet.IP.To4()
    if ip == nil { return nil }
    mask := ipnet.Mask
    randByte := func() byte { return byte(time.Now().UnixNano() % 255) }
    for i := range ip {
        ip[i] = (ip[i] & mask[i]) | (randByte() & ^mask[i])
    }
    return ip
}

func isClean(ip net.IP) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip.String(), "443"), 2*time.Second)
    if err != nil { return false }
    conn.Close()
    return true
}
