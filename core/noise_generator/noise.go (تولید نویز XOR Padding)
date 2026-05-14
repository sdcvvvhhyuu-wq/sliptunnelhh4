package noise_generator

import (
    "crypto/rand"
    "log"
    "time"
)

func StartNoiseGenerator(min, max int, interval time.Duration) {
    go func() {
        for {
            size := min + int(time.Now().UnixNano())%(max-min)
            garbage := make([]byte, size)
            rand.Read(garbage)
            log.Printf("Noise generator: sent %d bytes of random padding", size)
            time.Sleep(interval)
        }
    }()
}
