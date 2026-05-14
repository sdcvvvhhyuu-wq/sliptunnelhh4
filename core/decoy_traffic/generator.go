package decoy_traffic

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"time"
)

// DecoyTrafficGenerator creates realistic dummy traffic
type DecoyTrafficGenerator struct {
	patterns  []TrafficPattern
	intensity int // 1-10
}

// TrafficPattern represents a traffic type
type TrafficPattern struct {
	Name         string
	PacketSize   int
	Interval     time.Duration
	Probability  float64
	Description  string
}

// NewDecoyTrafficGenerator creates a new generator
func NewDecoyTrafficGenerator(intensity int) *DecoyTrafficGenerator {
	return &DecoyTrafficGenerator{
		patterns: []TrafficPattern{
			{
				Name:        "HTTP_Browse",
				PacketSize:  randInt(512, 2048),
				Interval:    time.Duration(randInt(100, 500)) * time.Millisecond,
				Probability: 0.4,
				Description: "Simulates web browsing",
			},
			{
				Name:        "Video_Stream",
				PacketSize:  randInt(4096, 8192),
				Interval:    time.Duration(randInt(50, 200)) * time.Millisecond,
				Probability: 0.3,
				Description: "Simulates video streaming",
			},
			{
				Name:        "VoIP_Call",
				PacketSize:  randInt(160, 320),
				Interval:    time.Duration(20) * time.Millisecond,
				Probability: 0.2,
				Description: "Simulates VoIP call",
			},
			{
				Name:        "File_Transfer",
				PacketSize:  randInt(1024, 4096),
				Interval:    time.Duration(randInt(200, 1000)) * time.Millisecond,
				Probability: 0.1,
				Description: "Simulates file transfer",
			},
		},
		intensity: intensity,
	}
}

// GenerateHTTPTraffic creates realistic HTTP traffic
func (dtg *DecoyTrafficGenerator) GenerateHTTPTraffic() []byte {
	httpTemplates := []string{
		"GET / HTTP/1.1\r\nHost: example.com\r\n\r\n",
		"POST /api HTTP/1.1\r\nHost: api.example.com\r\n\r\n",
		"GET /images HTTP/1.1\r\nHost: cdn.example.com\r\n\r\n",
	}

	template := httpTemplates[randInt(0, len(httpTemplates))]
	traffic := make([]byte, len(template)+randInt(100, 500))
	copy(traffic, template)
	rand.Read(traffic[len(template):])

	return traffic
}

// GenerateVideoStreamTraffic simulates video streaming
func (dtg *DecoyTrafficGenerator) GenerateVideoStreamTraffic() []byte {
	// Simulate MP4/WebM chunks
	chunking := []byte{0xFF, 0xD8, 0xFF, 0xE0} // JPEG magic
	padding := make([]byte, randInt(2000, 8000))
	rand.Read(padding)

	return append(chunking, padding...)
}

// GenerateVoIPTraffic simulates VOIP packets
func (dtg *DecoyTrafficGenerator) GenerateVoIPTraffic() []byte {
	// Simulate G.711 codec packets
	voipPacket := make([]byte, 160) // Standard VoIP packet size
	rand.Read(voipPacket)
	return voipPacket
}

// GeneratePattern creates traffic following a pattern
func (dtg *DecoyTrafficGenerator) GeneratePattern(pattern TrafficPattern) []byte {
	data := make([]byte, pattern.PacketSize)
	io.ReadFull(rand.Reader, data)
	return data
}

// GenerateDecoySequence creates a sequence of decoy packets
func (dtg *DecoyTrafficGenerator) GenerateDecoySequence(duration time.Duration) [][]byte {
	var sequence [][]byte
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		pattern := dtg.selectPattern()
		packet := dtg.GeneratePattern(pattern)
		sequence = append(sequence, packet)

		time.Sleep(pattern.Interval)
	}

	return sequence
}

// selectPattern randomly selects a pattern based on probability
func (dtg *DecoyTrafficGenerator) selectPattern() TrafficPattern {
	r := rand.Float64()
	cumProb := 0.0

	for _, p := range dtg.patterns {
		cumProb += p.Probability
		if r <= cumProb {
			return p
		}
	}

	return dtg.patterns[len(dtg.patterns)-1]
}

// GetTrafficMix returns mix of different traffic types
func (dtg *DecoyTrafficGenerator) GetTrafficMix() map[string]int {
	mix := make(map[string]int)
	totalPackets := 100 * dtg.intensity

	for _, p := range dtg.patterns {
		mix[p.Name] = int(math.Round(float64(totalPackets) * p.Probability))
	}

	return mix
}

// DistributeTraffic spreads traffic across time
func (dtg *DecoyTrafficGenerator) DistributeTraffic(packets int, duration time.Duration) []time.Duration {
	timings := make([]time.Duration, packets)
	interval := duration / time.Duration(packets)

	for i := 0; i < packets; i++ {
		jitter := time.Duration(randInt(-50, 50)) * time.Millisecond
		timings[i] = interval + jitter
	}

	return timings
}

// Helper function
func randInt(min, max int) int {
	b := make([]byte, 4)
	rand.Read(b)
	value := int(b[0])%(max-min) + min
	return value
}

// DescribePattern returns description of traffic pattern
func (dtg *DecoyTrafficGenerator) DescribePattern(name string) string {
	for _, p := range dtg.patterns {
		if p.Name == name {
			return fmt.Sprintf("%s: %s (Size: %d bytes, Interval: %v)",
				p.Name, p.Description, p.PacketSize, p.Interval)
		}
	}
	return "Pattern not found"
}
