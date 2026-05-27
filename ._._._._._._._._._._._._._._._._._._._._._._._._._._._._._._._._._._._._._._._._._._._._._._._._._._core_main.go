package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/sdcvvvhhyuu-wq/argotunnel/capabilities"
    "github.com/sdcvvvhhyuu-wq/argotunnel/engine"
)

func main() {
    log.Println("ArgoTunnel Ultimate – AI Adaptive Anti‑Censorship System")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go func() {
        if err := exec.Start(); err != nil {
            log.Fatalf("Engine failed: %v", err)
        }
    }()
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
}
