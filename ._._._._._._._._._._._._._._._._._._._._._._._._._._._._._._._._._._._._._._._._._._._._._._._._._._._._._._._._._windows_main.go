package main

import (
    "fmt"
    "os"
    "os/signal"
    "github.com/sdcvvvhhyuu-wq/argotunnel/capabilities"
    "github.com/sdcvvvhhyuu-wq/argotunnel/engine"
)

func main() {
    fmt.Println("ArgoTunnel Ultimate for Windows")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go exec.Start()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}
