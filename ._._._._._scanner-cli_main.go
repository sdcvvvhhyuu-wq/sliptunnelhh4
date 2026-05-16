package main

import (
    "flag"
    "fmt"
    "net"
    "time"
)

func main() {
    cidr := flag.String("cidr", "104.16.0.0/12", "CIDR range to scan")
    port := flag.String("port", "443", "port to check")
    timeout := flag.Duration("timeout", 500*time.Millisecond, "dial timeout")
    flag.Parse()

    _, ipnet, _ := net.ParseCIDR(*cidr)
    for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
        if isClean(ip, *port, *timeout) {
            fmt.Println(ip)
        }
    }
}

func inc(ip net.IP) {
    for j := len(ip)-1; j>=0; j-- {
        ip[j]++
        if ip[j] > 0 { break }
    }
}

func isClean(ip net.IP, port string, timeout time.Duration) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip.String(), port), timeout)
    if err == nil { conn.Close(); return true }
    return false
}
