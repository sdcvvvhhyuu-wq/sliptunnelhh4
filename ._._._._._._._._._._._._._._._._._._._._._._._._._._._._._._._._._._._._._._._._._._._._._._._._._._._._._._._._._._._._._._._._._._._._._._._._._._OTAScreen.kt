package com.slipnet.app.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import com.slipnet.app.OTAUpdater
import kotlinx.coroutines.launch

@Composable
fun OTAScreen() {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    var status by remember { mutableStateOf("Click to check") }
    Button(onClick = {
        scope.launch {
            val url = OTAUpdater.checkForUpdates(context)
            if (url != null) { OTAUpdater.downloadAndInstall(context, url); status = "Downloading..." }
            else status = "Up to date"
        }
    }, modifier = Modifier.padding(16.dp)) { Text(status) }
}
