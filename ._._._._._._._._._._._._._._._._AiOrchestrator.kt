package com.sliptunnel

import android.util.Log
import kotlinx.coroutines.*
import java.net.InetSocketAddress
import java.net.Socket

object AiOrchestrator {
    private const val TAG = "AiOrchestrator"
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())
    private val scores = mutableMapOf("wireguard" to 100, "shadowsocks" to 95, "dns_tunnel" to 85, "icmp_tunnel" to 70, "naiveproxy" to 90, "slipstream" to 92)

    fun start() { scope.launch { while (isActive) { adapt(); delay(30_000) } } }

    private suspend fun adapt() = withContext(Dispatchers.IO) {
        val lat = measureLatency()
        val loss = (Math.random() * 5).toInt()
        if (lat > 400 || loss > 8) {
            boost("dns_tunnel", 15); boost("icmp_tunnel", 20); reduce("wireguard", 10)
        } else {
            boost("wireguard", 10); boost("slipstream", 15); reduce("dns_tunnel", 5)
        }
    }
    private fun boost(t: String, a: Int) { scores[t] = (scores[t] ?: 50) + a }
    private fun reduce(t: String, a: Int) { scores[t] = maxOf(10, (scores[t] ?: 50) - a) }
    private fun measureLatency(): Long = try {
        val s = System.currentTimeMillis()
        Socket().use { it.connect(InetSocketAddress("8.8.8.8", 53), 2000) }
        System.currentTimeMillis() - s
    } catch (_: Exception) { 999 }
    fun bestTunnel() = scores.maxByOrNull { it.value }?.key ?: "wireguard"
}
