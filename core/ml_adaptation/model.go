package ml_adaptation

import (
	"fmt"
	"sync"
	"time"
)

// AdaptationModel learns filtering patterns and adapts in real-time
type AdaptationModel struct {
	mu               sync.RWMutex
	detectionEvents  []DetectionEvent
	evasionStrategies []StrategyPerformance
	currentStrategy  string
	learningRate     float64
}

// DetectionEvent records when detection occurs
type DetectionEvent struct {
	Timestamp   time.Time
	FilterType  string
	Severity    int // 1-10
	TunnelType  string
	FailureRate float64
}

// StrategyPerformance tracks how well strategies work
type StrategyPerformance struct {
	Name            string
	SuccessRate     float64
	LastUsed        time.Time
	FailureCount    int
	SuccessCount    int
	AverageLatency  float64
	Effectiveness   float64 // 0-1
}

// NewAdaptationModel creates a new ML model
func NewAdaptationModel() *AdaptationModel {
	return &AdaptationModel{
		detectionEvents: []DetectionEvent{},
		evasionStrategies: []StrategyPerformance{
			{
				Name:         "Polymorphic",
				SuccessRate:  0.85,
				Effectiveness: 0.85,
			},
			{
				Name:         "Fragmentation",
				SuccessRate:  0.72,
				Effectiveness: 0.72,
			},
			{
				Name:         "Timing",
				SuccessRate:  0.68,
				Effectiveness: 0.68,
			},
			{
				Name:         "Decoy",
				SuccessRate:  0.65,
				Effectiveness: 0.65,
			},
			{
				Name:         "Protocol Morphing",
				SuccessRate:  0.90,
				Effectiveness: 0.90,
			},
		},
		currentStrategy: "Polymorphic",
		learningRate:    0.01,
	}
}

// RecordDetection records a detection event
func (am *AdaptationModel) RecordDetection(event DetectionEvent) {
	am.mu.Lock()
	defer am.mu.Unlock()

	event.Timestamp = time.Now()
	am.detectionEvents = append(am.detectionEvents, event)

	// Update based on event
	am.updateStrategies()
}

// RecordSuccess records a successful tunnel connection
func (am *AdaptationModel) RecordSuccess(strategyName string, latency float64) {
	am.mu.Lock()
	defer am.mu.Unlock()

	for i := range am.evasionStrategies {
		if am.evasionStrategies[i].Name == strategyName {
			am.evasionStrategies[i].SuccessCount++
			am.evasionStrategies[i].LastUsed = time.Now()
			am.evasionStrategies[i].AverageLatency = 
				(am.evasionStrategies[i].AverageLatency + latency) / 2
			am.updateSuccessRate(i)
		}
	}
}

// RecordFailure records a failed connection attempt
func (am *AdaptationModel) RecordFailure(strategyName string) {
	am.mu.Lock()
	defer am.mu.Unlock()

	for i := range am.evasionStrategies {
		if am.evasionStrategies[i].Name == strategyName {
			am.evasionStrategies[i].FailureCount++
			am.updateSuccessRate(i)
		}
	}
}

// updateSuccessRate updates strategy success rate
func (am *AdaptationModel) updateSuccessRate(index int) {
	total := am.evasionStrategies[index].SuccessCount + am.evasionStrategies[index].FailureCount
	if total > 0 {
		am.evasionStrategies[index].SuccessRate = 
			float64(am.evasionStrategies[index].SuccessCount) / float64(total)
		am.evasionStrategies[index].Effectiveness = 
			am.evasionStrategies[index].SuccessRate * 
			(1.0 - (am.evasionStrategies[index].AverageLatency / 10000.0))
	}
}

// updateStrategies adapts strategies based on detections
func (am *AdaptationModel) updateStrategies() {
	// Analyze recent detections
	if len(am.detectionEvents) < 2 {
		return
	}

	recent := am.detectionEvents[len(am.detectionEvents)-1]

	// Adjust strategies based on filter type
	switch recent.FilterType {
	case "TLS_Fingerprint":
		am.boostStrategy("Polymorphic", 0.05)
	case "SNI_Blocking":
		am.boostStrategy("Protocol Morphing", 0.05)
	case "DNS_Poisoning":
		am.boostStrategy("Fragmentation", 0.03)
	case "Packet_Pattern":
		am.boostStrategy("Timing", 0.04)
	}
}

// boostStrategy increases effectiveness of a strategy
func (am *AdaptationModel) boostStrategy(name string, boost float64) {
	for i := range am.evasionStrategies {
		if am.evasionStrategies[i].Name == name {
			am.evasionStrategies[i].Effectiveness += boost
			if am.evasionStrategies[i].Effectiveness > 1.0 {
				am.evasionStrategies[i].Effectiveness = 1.0
			}
		}
	}
}

// SelectBestStrategy selects optimal strategy
func (am *AdaptationModel) SelectBestStrategy() string {
	am.mu.RLock()
	defer am.mu.RUnlock()

	bestIdx := 0
	bestEffectiveness := am.evasionStrategies[0].Effectiveness

	for i := 1; i < len(am.evasionStrategies); i++ {
		if am.evasionStrategies[i].Effectiveness > bestEffectiveness {
			bestEffectiveness = am.evasionStrategies[i].Effectiveness
			bestIdx = i
		}
	}

	am.currentStrategy = am.evasionStrategies[bestIdx].Name
	return am.currentStrategy
}

// GetStrategyStats returns statistics about strategies
func (am *AdaptationModel) GetStrategyStats() map[string]interface{} {
	am.mu.RLock()
	defer am.mu.RUnlock()

	stats := make(map[string]interface{})
	for _, s := range am.evasionStrategies {
		stats[s.Name] = map[string]interface{}{
			"success_rate":  s.SuccessRate,
			"effectiveness": s.Effectiveness,
			"avg_latency":   s.AverageLatency,
			"success_count": s.SuccessCount,
			"failure_count": s.FailureCount,
		}
	}

	return stats
}

// GetDetectionTrends analyzes trends in detections
func (am *AdaptationModel) GetDetectionTrends() map[string]int {
	am.mu.RLock()
	defer am.mu.RUnlock()

	trends := make(map[string]int)
	for _, event := range am.detectionEvents {
		trends[event.FilterType]++
	}

	return trends
}

// PredictNextFilter predicts likely next filter based on patterns
func (am *AdaptationModel) PredictNextFilter() string {
	am.mu.RLock()
	defer am.mu.RUnlock()

	if len(am.detectionEvents) < 3 {
		return "Unknown"
	}

	// Simple prediction: most common recent filter
	recent := am.detectionEvents[len(am.detectionEvents)-3:]
	filterCounts := make(map[string]int)

	for _, event := range recent {
		filterCounts[event.FilterType]++
	}

	maxCount := 0
	prediction := "Unknown"
	for filter, count := range filterCounts {
		if count > maxCount {
			maxCount = count
			prediction = filter
		}
	}

	return prediction
}

// GetRecommendation provides recommendation for user
func (am *AdaptationModel) GetRecommendation() string {
	am.mu.RLock()
	defer am.mu.RUnlock()

	prediction := am.PredictNextFilter()
	current := am.currentStrategy

	recommendation := fmt.Sprintf(
		"Current Strategy: %s\nPredicted Filter: %s\nRecommended Action: Use best available tunnel",
		current, prediction,
	)

	return recommendation
}
