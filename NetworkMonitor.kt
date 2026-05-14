package com.sliptunnel

import android.content.Context
import android.net.ConnectivityManager
import android.net.Network
import android.net.NetworkCapabilities
import android.net.NetworkRequest
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow

object NetworkMonitor {
    private val _connected = MutableStateFlow(true)
    val isConnected: StateFlow<Boolean> = _connected

    fun start(ctx: Context) {
        val cm = ctx.getSystemService(Context.CONNECTIVITY_SERVICE) as ConnectivityManager
        val req = NetworkRequest.Builder().addCapability(NetworkCapabilities.NET_CAPABILITY_INTERNET).addCapability(NetworkCapabilities.NET_CAPABILITY_VALIDATED).build()
        cm.registerNetworkCallback(req, object : ConnectivityManager.NetworkCallback() {
            override fun onAvailable(network: Network) { _connected.value = true; FallbackEngine.deescalate() }
            override fun onLost(network: Network) { _connected.value = false }
            override fun onUnavailable() { _connected.value = false }
        })
    }
}
