package obfs

import "log"

var currentFrontingHost string

func SetFrontingHost(host string) {
    currentFrontingHost = host
    log.Printf("Fronting host set to: %s", host)
}

func GetFrontingHost() string {
    return currentFrontingHost
}
