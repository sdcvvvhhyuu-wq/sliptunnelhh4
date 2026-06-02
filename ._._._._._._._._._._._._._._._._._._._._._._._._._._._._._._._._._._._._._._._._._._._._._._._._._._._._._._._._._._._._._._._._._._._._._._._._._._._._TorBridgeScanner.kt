package com.sliptunnel

import com.google.gson.Gson
import kotlinx.coroutines.*
import java.net.HttpURLConnection
import java.net.URL

object TorBridgeScanner {
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())
    private val _workingBridges = mutableListOf<String>()
    val workingBridges: List<String> get() = _workingBridges

    data class Relay(val fingerprint: String, val or_addresses: List<String>, val flags: List<String>?)

    fun startAutoScan() {
        scope.launch {
            while (isActive) {
                refreshBridges()
                delay(60 * 60 * 1000) // هر ساعت
            }
        }
    }

    suspend fun manualScan(): List<String> = withContext(Dispatchers.IO) {
        refreshBridges()
        _workingBridges.toList()
    }

    suspend fun importAndTest(lines: List<String>): List<String> = withContext(Dispatchers.IO) {
        val results = mutableListOf<String>()
        for (line in lines) {
            if (line.startsWith("obfs4 ") || line.startsWith("Bridge ")) {
                val parts = line.replace("Bridge ", "").split(" ")
                if (parts.size >= 3) {
                    val ip = parts[1]
                    val port = parts[2].toIntOrNull() ?: 443
                    if (tcpCheck(ip, port, 3000)) results.add(line)
                }
            }
        }
        _workingBridges.addAll(results)
        _workingBridges.toList()
    }

    private fun refreshBridges() {
        _workingBridges.clear()
        try {
            val json = URL("https://onionoo.torproject.org/details?type=relay&running=true&flag=Stable").readText()
            val data = Gson().fromJson(json, Map::class.java)
            val relays = data["relays"] as? List<Map<String, Any>> ?: return
            for (relay in relays.take(100)) {
                val addresses = relay["or_addresses"] as? List<String> ?: continue
                for (addr in addresses.take(2)) {
                    val clean = addr.replace("[", "").replace("]", "").split(":")
                    if (clean.size == 2) {
                        val ip = clean[0]
                        val port = clean[1].toIntOrNull() ?: continue
                        if (tcpCheck(ip, port, 2500)) {
                            _workingBridges.add("$ip:$port")
                        }
                    }
                }
            }
        } catch (_: Exception) { }
    }

    private fun tcpCheck(host: String, port: Int, timeout: Int): Boolean = try {
        java.net.Socket().use { it.connect(java.net.InetSocketAddress(host, port), timeout) }
        true
    } catch (_: Exception) { false }
}
