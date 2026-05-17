#!/bin/bash

echo "🔍 بررسی استخراج vip.txt"
echo "=================================="

# Count files by type
GO_FILES=$(find . -name "*.go" 2>/dev/null | wc -l)
KOTLIN_FILES=$(find . -name "*.kt" 2>/dev/null | wc -l)
SWIFT_FILES=$(find . -name "*.swift" 2>/dev/null | wc -l)
CONFIG_FILES=$(find . \( -name "*.yml" -o -name "*.yaml" -o -name "*.gradle.kts" -o -name "*.plist" -o -name "*.xml" \) 2>/dev/null | wc -l)

echo "📊 تعداد فایل‌ها:"
echo "  • Go files:        $GO_FILES"
echo "  • Kotlin files:    $KOTLIN_FILES"
echo "  • Swift files:     $SWIFT_FILES"
echo "  • Config files:    $CONFIG_FILES"

TOTAL=$((GO_FILES + KOTLIN_FILES + SWIFT_FILES + CONFIG_FILES))
echo "  ━━━━━━━━━━━━━━━━━"
echo "  • جمع:            $TOTAL فایل"

# Check critical files
echo ""
echo "✅ فایل‌های اساسی:"
[ -f ".github/workflows/build-all.yml" ] && echo "  ✓ CI/CD Workflow" || echo "  ✗ CI/CD Workflow"
[ -f "core/main.go" ] && echo "  ✓ Core main.go" || echo "  ✗ Core main.go"
[ -f "android/app/build.gradle.kts" ] && echo "  ✓ Android build" || echo "  ✗ Android build"
[ -f "core/go.mod" ] && echo "  ✓ Go modules" || echo "  ✗ Go modules"
[ -f ".gitignore" ] && echo "  ✓ .gitignore" || echo "  ✗ .gitignore"

# Directory structure
echo ""
echo "📁 ساختار پوشه‌ها:"
for dir in core android ios windows linux scanner-cli openwrt .github; do
    if [ -d "$dir" ]; then
        FILE_COUNT=$(find "$dir" -type f 2>/dev/null | wc -l)
        echo "  ✓ $dir ($FILE_COUNT فایل)"
    fi
done

echo ""
echo "=================================="
echo "✅ استخراج کامل شد!"

