# 🚀 SlipTunnel - Quick Start Guide

## ⚡ 30-Second Setup

```bash
# 1. Clone repository
git clone https://github.com/sdcvvvhhyuu-wq/sliptunnel.git
cd sliptunnel

# 2. Build core
cd core
go mod tidy
go build

# 3. Build Android APK
cd ../android
./gradlew assembleRelease

# 4. Done! Check output
ls ../android/app/build/outputs/apk/release/
```

## 📋 What Got Extracted?

From **vip.txt** (26,349 lines):
- ✅ **69+ source files** across 8 directories
- ✅ **35+ tunnel protocols** (DNS, SSH, Tor, etc.)
- ✅ **Advanced AI modules** (traffic morphing, DPI detection)
- ✅ **Multi-platform support** (Android, iOS, Windows, Linux)
- ✅ **Automated CI/CD** (GitHub Actions ready)

## 🏗️ Project Structure

```
core/                  # Go backend (36 files)
├── tunnel/           # DNS, SSH, Tor, NaiveProxy, DoH, ICMP
├── obfs/             # ECH, Fronting, Fragmentation, uTLS
├── ai_morph/         # AI traffic morphing
├── dpi_analyzer/     # DPI detection
└── pqc/              # Post-quantum crypto

android/             # Android app (25 files)
├── app/build.gradle.kts
└── src/main/
    ├── kotlin/      # UI (Jetpack Compose)
    └── java/        # VPN Service

ios/                 # iOS app (9 files)
├── ArgoTunnel/      # Main app
└── PacketTunnelProvider/  # VPN extension

windows/linux/scanner-cli/openwrt/  # Other platforms
```

## 🎯 Key Commands

### Build All Platforms
```bash
# Android
cd android && ./gradlew assembleRelease

# CLI Scanner
cd scanner-cli && go build -o ../bin/scanner

# Windows/Linux
cd windows && go build
cd ../linux && go build
```

### Create Release
```bash
git tag v1.0.0-alpha
git push origin --tags
# GitHub Actions automatically builds all platforms
```

### Verify Installation
```bash
./validate_extraction.sh
```

## 🔧 Core Modules

### Tunneling
- DNS (DNSTT, NoizDNS, VayDNS, SlipStream)
- SSH (4 variants with TLS/WebSocket/HTTP)
- Tor (with bridge support)
- NaiveProxy, DoH, ICMP, WireGuard

### Anti-DPI
- ECH (Encrypted Client Hello)
- Domain Fronting
- Packet Fragmentation
- Traffic Noise

### AI/ML
- Traffic Morphing
- DPI Analysis
- Reinforcement Learning
- GAN Traffic Generation

## 📱 Mobile Apps

### Android
- Kotlin + Jetpack Compose UI
- VPN Service + Quick Tile
- OTA auto-updates
- Tunnel selection UI

### iOS
- Swift implementation
- PacketTunnel extension
- Network extension support
- Always-on VPN capability

## 🔐 Security Features

- Post-Quantum Cryptography (Kyber)
- Automatic tunnel selection
- Multi-tunnel failover
- DPI evasion techniques
- Real-time threat detection

## 📊 Build Requirements

| Tool | Version |
|------|---------|
| Go | 1.22+ |
| Java | 17+ |
| Android SDK | 31+ |
| Gradle | 8.0+ (auto-download) |

## ✅ Documentation

| File | Purpose |
|------|---------|
| EXTRACTION_COMPLETE.md | Full guide (English) |
| EXTRACTION_SUMMARY.md | Project overview |
| PROJECT_STATUS.md | Module inventory |
| EXTRACTED_README.md | Structure reference |
| This file | Quick reference |

## 🚨 First Time Building?

1. **Install dependencies**
   ```bash
   go mod tidy
   ./gradlew wrapper --gradle-version=8.0
   ```

2. **Set up Android SDK**
   ```bash
   export ANDROID_SDK_ROOT=$HOME/Android/Sdk
   ```

3. **Create GitHub secret**
   ```bash
   gh secret set GH_TOKEN --body "$(gh auth token)"
   ```

4. **Test build**
   ```bash
   cd core && go build -o ../bin/test
   ```

## 🎓 Module Details

### Example: Building DNS Tunnel
```go
// File: core/tunnel/dns_tunnel.go
// Implements DNSTT, NoizDNS, VayDNS protocols
// Auto-selects based on network conditions
```

### Example: Android UI
```kotlin
// File: android/app/src/main/kotlin/com/argotunnel/MainActivity.kt
// Jetpack Compose UI with tunnel selection
// Real-time connection status display
```

## 🔗 Important Links

- **Repository**: https://github.com/sdcvvvhhyuu-wq/sliptunnel
- **Actions**: https://github.com/sdcvvvhhyuu-wq/sliptunnel/actions
- **Releases**: https://github.com/sdcvvvhhyuu-wq/sliptunnel/releases

## ❓ Troubleshooting

```bash
# Verify extraction
find . -name "*.go" | wc -l          # Should show 36+

# Check Go modules
cd core && go mod verify              # Must pass

# Verify gradle
cd ../android && ./gradlew check      # Must pass

# View git log
git log --oneline                     # Should show 3+ commits
```

## 🎯 Next Steps

1. ✅ **Understand** the extracted structure
2. ✅ **Build** each component
3. ✅ **Test** functionality
4. ✅ **Deploy** to your infrastructure
5. ✅ **Monitor** CI/CD pipeline

---

**Status**: ✅ READY TO BUILD  
**Version**: v1.0.0-alpha  
**Last Updated**: 2026-05-14  

Get started now! 🚀
