package com.slipnet.app

import android.app.Application

class SlipNetApp : Application() {
    override fun onCreate() {
        super.onCreate()
        NetworkMonitor.start(this)
        AiOrchestrator.start()
        FallbackEngine.start()
    }
}
