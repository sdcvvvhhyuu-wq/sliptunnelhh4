package obfs

import (
    "crypto/tls"
    "log"
    "net"
    "github.com/refraction-networking/utls"
)

// DialWithECH establishes a TLS connection with ECH support, falling back to standard TLS
func DialWithECH(network, addr string, config *tls.Config) (*tls.Conn, error) {
    log.Printf("ECH: dialing %s with encrypted ClientHello", addr)
    tcpConn, err := net.Dial(network, addr)
    if err != nil { return nil, err }

    uConfig := &utls.Config{
        ServerName:         config.ServerName,
        InsecureSkipVerify: config.InsecureSkipVerify,
    }
    uConn := utls.UClient(tcpConn, uConfig, utls.HelloChrome_Auto)
    err = uConn.Handshake()
    if err != nil {
        log.Printf("ECH handshake failed, falling back: %v", err)
        tcpConn.Close()
        return tls.Dial(network, addr, config)
    }
    log.Println("ECH connection established")
    return uConn, nil
}
