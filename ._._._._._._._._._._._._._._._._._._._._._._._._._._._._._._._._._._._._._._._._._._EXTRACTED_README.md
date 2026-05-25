# ArgoTunnel/SlipTunnel - استخراج شده از vip.txt

✅ **استخراج کامل و خودکار انجام شد**

## 📊 خلاصه استخراج

- **تعداد فایل‌های استخراج‌شده**: 92 فایل
- **منبع**: vip.txt (26,349 خط)
- **وضعیت**: آماده برای ساخت و استقرار

## 📁 ساختار فایل‌ها

```
sliptunnel/
├── .github/workflows/
│   └── build-all.yml          # CI/CD - ساخت خودکار
├── android/                    # برنامه Android
│   ├── build.gradle.kts
│   ├── settings.gradle.kts
│   └── app/src/main/
│       ├── AndroidManifest.xml
│       ├── kotlin/com/argotunnel/
│       └── res/
├── core/                       # هسته Go
│   ├── main.go
│   ├── go.mod
│   ├── engine/
│   ├── tunnel/                 # ماژول‌های تونل
│   │   ├── tunnel.go
│   │   ├── dns_tunnel.go
│   │   ├── ssh.go
│   │   ├── naiveproxy.go
│   │   ├── tor.go
│   │   └── ...
│   ├── obfs/                   # مخفی‌سازی
│   │   ├── fragmentation.go
│   │   ├── utls.go
│   │   └── fronting.go
│   ├── pqc/                    # رمزنگاری کوانتومی
│   │   └── kyber.go
│   ├── ai_morph/               # هوش مصنوعی برای تطابق ترافیک
│   ├── dpi_analyzer/           # تحلیل DPI
│   └── ...
├── ios/                        # برنامه iOS
├── windows/                    # نسخه Windows
├── linux/                      # نسخه Linux
├── scanner-cli/                # ابزار اسکن CLI
├── openwrt/                    # پشتیبانی OpenWrt
├── .gitignore
├── Makefile
└── README.md
```

## 🔧 ماژول‌های اصلی

### 1. تونل‌ها (Tunnels)
- **DNS Tunnel** (DNSTT, NoizDNS, VayDNS, SlipStream)
- **SSH** (SSH، SSH+TLS، SSH+WebSocket، SSH+HTTP CONNECT)
- **Tor**
- **NaiveProxy**
- **DoH Tunnel**
- **ICMP Tunnel**

### 2. مخفی‌سازی (Obfuscation)
- **ECH (Encrypted Client Hello)**
- **Fronting (Domain Fronting)**
- **Fragmentation** (تکه‌کاری ترافیک)
- **uTLS** (TLS Fingerprint Spoofing)

### 3. قابلیت‌های پیشرفته
- **PQC (Post-Quantum Cryptography)** - Kyber
- **AI Traffic Morphing** - تطابق هوشمند ترافیک
- **Active Shield** - دفاع فعال
- **Dynamic Orchestration** - ترکیب تونل‌های چند‌گانه
- **DPI Analyzer** - تشخیص DPI
- **QRNG** - تولید اعداد تصادفی کوانتومی
- **RL Agent** - بهینه‌سازی یادگیری تقویتی
- **GAN Generator** - تولید ترافیک مصنوعی
- **QUIC Masquerade** - نقاب‌گذاری QUIC

## 🚀 نحوه شروع

### پیش‌نیازها
```bash
# Go 1.22+
go version

# Android SDK (برای ساخت Android)
export ANDROID_SDK_ROOT=~/Android/Sdk

# Java 17+
java -version
```

### ساخت هسته Go
```bash
cd core
go mod tidy
go build -o ../bin/sliptunnel-core
```

### ساخت APK (Android)
```bash
cd android
./gradlew assembleRelease
```

### ساخت اسکنر CLI
```bash
cd scanner-cli
go build -o ../bin/scanner
```

## 🔄 CI/CD خودکار

فایل `.github/workflows/build-all.yml` تمام پلتفرم‌ها را ساخته‌ی:
- ✅ Android (APK با امضای Release)
- ✅ iOS (Xarchive)
- ✅ Windows (EXE)
- ✅ Linux (Binary)
- ✅ OpenWrt (IPK)

**نحوه استفاده:**
1. Push یا tag بسازید: `git tag v1.0.0`
2. GitHub Actions خودکار شروع می‌شود
3. نسخه‌های جدید در Releases موجود می‌شوند

## 📝 نکات مهم

### Go Module Dependencies
فایل‌های `go.mod` در هر ماژول باید شامل dependencies زیر باشند:
- `github.com/xtls/xray-core`
- `github.com/hiddify/Hiddify-Manager`
- `github.com/klauspost/compress` (برای compression)
- و دیگر ...

### Android Dependencies
- Kotlin 1.9+
- Jetpack Compose
- Android SDK Level 31+

### Secret Variables
برای CI/CD نیاز دارید:
- `GH_TOKEN` - دسترسی repo و workflow

## ✅ بررسی صحت

```bash
# بررسی ساختار فایل‌ها
find . -name "*.go" | wc -l
find . -name "*.kt" | wc -l

# بررسی Go modules
cd core && go mod verify

# بررسی gradle
cd android && ./gradlew check
```

## 📌 وضعیت

| بخش | وضعیت |
|-----|------|
| استخراج فایل‌ها | ✅ کامل |
| ساختار پروژه | ✅ آماده |
| CI/CD | ✅ موجود |
| توثیق | ✅ کامل |
| آماده برای ساخت | ✅ بله |

---

**تاریخ استخراج**: 2026-05-14
**منبع**: vip.txt (26,349 خط)
**فایل‌های استخراج‌شده**: 92 فایل
