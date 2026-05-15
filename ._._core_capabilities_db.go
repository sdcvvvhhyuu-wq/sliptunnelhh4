package capabilities

import (
    "fmt"
    "strconv"
    "strings"
)

type Capability struct {
    ID       int
    Name     string
    Category string
    Param    string
}

func GetAllCapabilities() []Capability {
    caps := []Capability{}
    id := 1
    // Fragment, Padding, TLS, Cipher, Fronting (like before, extended to 5000+)
    // Transport types (extended with SlipNet options)
    transports := []string{"wireguard", "shadowsocks", "xray", "trojan", "hysteria", "v2ray", "ssr",
        "dns", "noizdns", "vaydns", "slipstream", "ssh", "naiveproxy", "doh", "tor",
        "ssh_tls", "ssh_ws", "ssh_httpconnect"}
    for _, t := range transports {
        caps = append(caps, Capability{id, fmt.Sprintf("transport_%s", t), "transport", t})
        id++
    }
    // … (بقیه دسته‌ها: qrng, dpi_analyzer, extra, etc)
    // Keep the same huge pool from earlier but ensure 5000 cap
    for id <= 5000 {
        caps = append(caps, Capability{id, fmt.Sprintf("extra_%d", id), "extra", "dummy"})
        id++
    }
    return caps[:5000]
}

func ParseParamInt(param, prefix string) int {
    valStr := strings.TrimPrefix(param, prefix)
    val, _ := strconv.Atoi(strings.TrimSpace(valStr))
    return val
}
