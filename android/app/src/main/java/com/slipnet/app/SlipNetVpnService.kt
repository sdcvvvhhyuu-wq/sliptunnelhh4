package com.slipnet.app

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Intent
import android.net.VpnService
import android.os.ParcelFileDescriptor
import androidx.core.app.NotificationCompat
import kotlinx.coroutines.*

class SlipNetVpnService : VpnService() {
    private var vpnInterface: ParcelFileDescriptor? = null
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        startForeground(1, createNotification())
        startVpn()
        return START_STICKY
    }

    private fun startVpn() {
        val builder = Builder()
            .setSession("SlipNet Quantum")
            .addAddress("10.0.0.2", 32)
            .addRoute("0.0.0.0", 0)
            .addDnsServer("1.1.1.1")
            .addDnsServer("8.8.8.8")
            .setMtu(1280)
            .setBlocking(true)
        vpnInterface = builder.establish()
        scope.launch {
            // in a real scenario, native Go core is invoked here
            while (isActive) {
                delay(60000)
            }
        }
    }

    private fun createNotification(): Notification {
        val channelId = "slipnet_vpn"
        val nm = getSystemService(NOTIFICATION_SERVICE) as NotificationManager
        nm.createNotificationChannel(
            NotificationChannel(channelId, "SlipNet VPN", NotificationManager.IMPORTANCE_LOW)
        )
        val pendingIntent = PendingIntent.getActivity(
            this, 0, Intent(this, MainActivity::class.java),
            PendingIntent.FLAG_IMMUTABLE
        )
        return NotificationCompat.Builder(this, channelId)
            .setContentTitle("SlipNet Active")
            .setContentText("Quantum‑safe, AI‑morphed traffic")
            .setSmallIcon(android.R.drawable.ic_lock_lock)
            .setContentIntent(pendingIntent)
            .setOngoing(true)
            .build()
    }

    override fun onDestroy() {
        scope.cancel()
        vpnInterface?.close()
        super.onDestroy()
    }
}
