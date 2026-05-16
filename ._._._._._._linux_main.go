package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "github.com/sdcvvvhhyuu-wq/argotunnel/capabilities"
    "github.com/sdcvvvhhyuu-wq/argotunnel/engine"
)

func main() {
    fmt.Println("ArgoTunnel Ultimate for Linux")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go exec.Start()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
}
