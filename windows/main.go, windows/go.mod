package main
import ("fmt"; "os"; "os/signal"; "github.com/sdcvvvhhyuu-wq/sliptunnel/capabilities"; "github.com/sdcvvvhhyuu-wq/sliptunnel/engine")
func main() {
    fmt.Println("SlipTunnel for Windows")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go exec.Start()
    c := make(chan os.Signal, 1); signal.Notify(c, os.Interrupt); <-c
}
