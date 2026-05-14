package engine

import (
    "log"
    "net"
    "time"
)

func StartInternetHealthCheck() {
    go func() {
        for {
            time.Sleep(10 * time.Second)
            if !isInternetReachable() {
                log.Println("❌ Internet down! Activating air-gap protocols (ICMP/DNS covert)")
            }
        }
    }()
}

func isInternetReachable() bool {
    conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 2*time.Second)
    if err != nil { return false }
    conn.Close()
    return true
}
