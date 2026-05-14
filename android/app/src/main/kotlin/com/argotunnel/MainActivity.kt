package com.argotunnel

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Surface
import androidx.compose.ui.Modifier
import com.argotunnel.ui.screens.HomeScreen
import com.argotunnel.ui.theme.ArgoTunnelTheme
import com.argotunnel.ui.theme.QuantumDark

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            ArgoTunnelTheme {
                Surface(modifier = Modifier.fillMaxSize(), color = QuantumDark) {
                    HomeScreen(context = this)
                }
            }
        }
    }
}
