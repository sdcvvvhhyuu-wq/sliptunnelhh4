package com.slipnet.app

import android.content.Intent
import android.service.quicksettings.TileService

class QuickTileService : TileService() {
    override fun onTileAdded() {
        qsTile.state = android.service.quicksettings.Tile.STATE_INACTIVE
        qsTile.updateTile()
    }

    override fun onClick() {
        if (qsTile.state == android.service.quicksettings.Tile.STATE_ACTIVE) {
            qsTile.state = android.service.quicksettings.Tile.STATE_INACTIVE
            stopService(Intent(this, SlipNetVpnService::class.java))
        } else {
            qsTile.state = android.service.quicksettings.Tile.STATE_ACTIVE
            startService(Intent(this, SlipNetVpnService::class.java))
        }
        qsTile.updateTile()
    }
}
