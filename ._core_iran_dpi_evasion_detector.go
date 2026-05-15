package iran_dpi_evasion

import (
	"crypto/rand"
	"fmt"
	"net"
	"sync"
	"time"
)

// IranDPIDetector detects and evades Iran's known DPI techniques
type IranDPIDetector struct {
	mu                 sync.RWMutex
	detectedSignatures map[string]DetectionSignature
	blockPatterns      []string
	activeEvasions     []EvasionTechnique
}

// DetectionSignature represents a known DPI signature
type DetectionSignature struct {
	Name        string
	Pattern     []byte
	Threshold   int
	Description string
}

// EvasionTechnique represents an evasion strategy
type EvasionTechnique struct {
	Name      string
	Active    bool
	Method    func() error
	Intensity int // 1-10, higher = more aggressive
}

// NewIranDPIDetector creates a new detector
func NewIranDPIDetector() *IranDPIDetector {
	detector := &IranDPIDetector{
		detectedSignatures: make(map[string]DetectionSignature),
		blockPatterns:      []string{},
		activeEvasions:     []EvasionTechnique{},
	}
	detector.initializeKnownPatterns()
	return detector
}

// initializeKnownPatterns sets up Iran's known DPI patterns
func (d *IranDPIDetector) initializeKnownPatterns() {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Known Iranian DPI signatures
	d.detectedSignatures["HTTPS_ClientHello"] = DetectionSignature{
		Name:        "HTTPS ClientHello Detection",
		Pattern:     []byte{0x16, 0x03, 0x01}, // TLS handshake
		Threshold:   1,
		Description: "Iran DPI detects standard TLS ClientHello patterns",
	}

	d.detectedSignatures["DNS_Poisoning"] = DetectionSignature{
		Name:        "DNS Poisoning",
		Pattern:     []byte{0x00, 0x01, 0x00, 0x00}, // DNS query
		Threshold:   1,
		Description: "Iran intercepts DNS queries for known sites",
	}

	d.detectedSignatures["SNI_Blocking"] = DetectionSignature{
		Name:        "SNI Blocking",
		Pattern:     []byte("SNI"),
		Threshold:   1,
		Description: "Iran blocks based on SNI (Server Name Indication)",
	}

	d.detectedSignatures["HTTP2_Framesize"] = DetectionSignature{
		Name:        "HTTP/2 Frame Size Detection",
		Pattern:     []byte{0x50, 0x52, 0x49}, // HTTP/2 PRI
		Threshold:   1,
		Description: "Iran detects HTTP/2 based on frame sizes",
	}

	d.detectedSignatures["Tor_Guard"] = DetectionSignature{
		Name:        "Tor Guard Detection",
		Pattern:     []byte{0x16, 0x03, 0x00}, // TLS v1.0 record
		Threshold:   5,
		Description: "Iran blocks based on Tor TLS patterns",
	}

	d.detectedSignatures["VPN_Packet_Pattern"] = DetectionSignature{
		Name:        "VPN Packet Pattern",
		Pattern:     []byte{}, // Variable patterns
		Threshold:   3,
		Description: "Iran detects repeating packet patterns from VPN",
	}

	d.blockPatterns = []string{
		"google.com", "facebook.com", "twitter.com", "youtube.com",
		"instagram.com", "whatsapp.com", "telegram.org", "medium.com",
		"reddit.com", "slack.com", "dropbox.com", "github.com",
	}
}

// DetectActiveFiltering checks for active DPI filtering
func (d *IranDPIDetector) DetectActiveFiltering() (bool, string) {
	testHosts := []string{
		"1.1.1.1:53",
		"8.8.8.8:53",
		"1.0.0.1:443",
	}

	for _, host := range testHosts {
		conn, err := net.DialTimeout("tcp", host, 2*time.Second)
		if err != nil {
			return true, fmt.Sprintf("DPI active: Cannot reach %s", host)
		}
		conn.Close()
	}

	return false, "No active DPI filtering detected"
}

// GeneratePolymorphicPayload creates changing payloads to evade DPI
func (d *IranDPIDetector) GeneratePolymorphicPayload(data []byte) []byte {
	randomOffset := make([]byte, randInt(1, 32))
	rand.Read(randomOffset)

	payload := make([]byte, len(data))
	copy(payload, data)

	timingBytes := make([]byte, randInt(4, 16))
	rand.Read(timingBytes)

	result := append(randomOffset, payload...)
	result = append(result, timingBytes...)

	return result
}

// EvadeWithFragmentation splits packets to evade DPI
func (d *IranDPIDetector) EvadeWithFragmentation(data []byte, fragmentSize int) [][]byte {
	var fragments [][]byte

	for i := 0; i < len(data); i += fragmentSize {
		end := i + fragmentSize
		if end > len(data) {
			end = len(data)
		}
		time.Sleep(time.Duration(randInt(1, 50)) * time.Millisecond)
		fragments = append(fragments, data[i:end])
	}

	return fragments
}

// Helper function
func randInt(min, max int) int {
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	value := int(randomBytes[0])%(max-min) + min
	return value
}
