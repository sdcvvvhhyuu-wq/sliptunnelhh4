package auto_fallback

import (
	"sync"
	"time"
)

// FallbackChain manages automatic tunnel switching
type FallbackChain struct {
	tunnels    []TunnelOption
	current    int
	mu         sync.RWMutex
	lastSwitch time.Time
}

// TunnelOption represents a tunnel configuration
type TunnelOption struct {
	Name        string
	Type        string // DNS, SSH, Tor, DoH, ICMP, etc.
	Endpoint    string
	Port        int
	Priority    int
	IsActive    bool
	FailCount   int
	LastSuccess time.Time
}

// NewFallbackChain creates a new fallback chain
func NewFallbackChain() *FallbackChain {
	return &FallbackChain{
		tunnels: []TunnelOption{
			{
				Name:     "DNS-DNSTT",
				Type:     "DNS",
				Endpoint: "1.1.1.1",
				Port:     53,
				Priority: 1,
				IsActive: true,
			},
			{
				Name:     "SSH-WebSocket",
				Type:     "SSH",
				Endpoint: "ssh-gateway.example.com",
				Port:     443,
				Priority: 2,
				IsActive: true,
			},
			{
				Name:     "TOR-Bridge",
				Type:     "Tor",
				Endpoint: "tor-bridge.onion",
				Port:     9050,
				Priority: 3,
				IsActive: true,
			},
			{
				Name:     "DoH-Cloudflare",
				Type:     "DoH",
				Endpoint: "1.1.1.1",
				Port:     443,
				Priority: 4,
				IsActive: true,
			},
			{
				Name:     "ICMP-Tunnel",
				Type:     "ICMP",
				Endpoint: "8.8.8.8",
				Port:     0,
				Priority: 5,
				IsActive: true,
			},
		},
		current: 0,
	}
}

// GetActive returns the currently active tunnel
func (fc *FallbackChain) GetActive() TunnelOption {
	fc.mu.RLock()
	defer fc.mu.RUnlock()
	return fc.tunnels[fc.current]
}

// MarkFailed marks tunnel as failed
func (fc *FallbackChain) MarkFailed(index int) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if index >= 0 && index < len(fc.tunnels) {
		fc.tunnels[index].FailCount++
		if fc.tunnels[index].FailCount > 3 {
			fc.tunnels[index].IsActive = false
		}
	}
}

// MarkSuccess records successful tunnel usage
func (fc *FallbackChain) MarkSuccess(index int) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if index >= 0 && index < len(fc.tunnels) {
		fc.tunnels[index].FailCount = 0
		fc.tunnels[index].LastSuccess = time.Now()
		fc.tunnels[index].IsActive = true
	}
}

// SwitchTunnel switches to next available tunnel
func (fc *FallbackChain) SwitchTunnel() bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	startIdx := fc.current
	for i := 0; i < len(fc.tunnels); i++ {
		nextIdx := (startIdx + i + 1) % len(fc.tunnels)
		if fc.tunnels[nextIdx].IsActive {
			fc.current = nextIdx
			fc.lastSwitch = time.Now()
			return true
		}
	}

	return false
}

// GetAllActive returns all active tunnels
func (fc *FallbackChain) GetAllActive() []TunnelOption {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	var active []TunnelOption
	for _, t := range fc.tunnels {
		if t.IsActive {
			active = append(active, t)
		}
	}
	return active
}

// AddTunnel adds a new tunnel option
func (fc *FallbackChain) AddTunnel(tunnel TunnelOption) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.tunnels = append(fc.tunnels, tunnel)
}

// RemoveTunnel removes a tunnel by index
func (fc *FallbackChain) RemoveTunnel(index int) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if index >= 0 && index < len(fc.tunnels) {
		fc.tunnels = append(fc.tunnels[:index], fc.tunnels[index+1:]...)
		return true
	}
	return false
}

// ResetFailures resets failure counts for all tunnels
func (fc *FallbackChain) ResetFailures() {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	for i := range fc.tunnels {
		fc.tunnels[i].FailCount = 0
		fc.tunnels[i].IsActive = true
	}
}

// GetStats returns statistics about tunnels
func (fc *FallbackChain) GetStats() map[string]interface{} {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	active := 0
	for _, t := range fc.tunnels {
		if t.IsActive {
			active++
		}
	}

	return map[string]interface{}{
		"total_tunnels":    len(fc.tunnels),
		"active_tunnels":   active,
		"current_tunnel":   fc.tunnels[fc.current].Name,
		"last_switch_time": fc.lastSwitch,
	}
}
