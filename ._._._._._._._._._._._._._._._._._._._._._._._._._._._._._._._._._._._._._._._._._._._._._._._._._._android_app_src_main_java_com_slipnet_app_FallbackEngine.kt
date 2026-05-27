package com.slipnet.app

import android.util.Log
import kotlinx.coroutines.*

object FallbackEngine {
    private const val TAG = "FallbackEngine"
    private var currentLevel = 0
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())
    private val levels = listOf(
        "direct_tls", "domain_fronting", "quic_masque", "doh",
        "shadowsocks", "dns_tunnel", "icmp_tunnel", "covert_http"
    )

    fun start() {
        scope.launch { while (isActive) { if (!NetworkMonitor.isConnected.value) escalate(); delay(5000) } }
    }

    fun escalate() {
        if (currentLevel < levels.lastIndex) {
            currentLevel++
            Log.w(TAG, "Escalating to ${levels[currentLevel]}")
        }
    }

    fun deescalate() {
        if (currentLevel > 0) {
            currentLevel--
            Log.i(TAG, "De-escalating to ${levels[currentLevel]}")
        }
    }

    fun currentStatus() = levels[currentLevel]
}
