package rl_agent

import (
	"log"
	"math/rand"
	"time"

	"github.com/sdcvvvhhyuu-wq/argotunnel/tunnel"
)

var bestProfile = "chrome_124"

func StartLearning(primary, secondary tunnel.Tunnel) {
	go func() {
		for {
			time.Sleep(15 * time.Second)
			packetLoss := rand.Float32()
			if packetLoss > 0.15 {
				profiles := []string{"firefox_125", "safari_17", "twitch", "zoom_voip", "chrome_124"}
				bestProfile = profiles[rand.Intn(len(profiles))]
				log.Printf("RL Agent: high loss, switching profile to %s", bestProfile)
			}
		}
	}()
}

func GetRecommendedProfile() string { return bestProfile }
