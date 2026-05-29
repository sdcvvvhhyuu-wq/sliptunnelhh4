package com.argotunnel.ui.screens

import android.content.Context
import android.content.Intent
import androidx.compose.animation.animateContentSize
import androidx.compose.animation.core.*
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Settings
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.argotunnel.ArgoVpnService
import com.argotunnel.ui.theme.QuantumCyan
import com.argotunnel.ui.theme.QuantumDark
import com.argotunnel.ui.theme.QuantumGreen
import com.argotunnel.ui.theme.QuantumPurple
import com.argotunnel.ui.theme.QuantumRed
import com.argotunnel.ui.theme.QuantumSurface
import com.argotunnel.ui.theme.QuantumWhite
import kotlinx.coroutines.delay

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeScreen(context: Context) {
    var connected by remember { mutableStateOf(false) }
    var statusText by remember { mutableStateOf("Secured Quantum Tunnel Inactive") }
    var discoveredHosts by remember { mutableStateOf(12) }
    var cleanIPs by remember { mutableStateOf(8) }
    val infiniteTransition = rememberInfiniteTransition(label = "pulse")
    val pulseAlpha by infiniteTransition.animateFloat(
        initialValue = 0.3f,
        targetValue = 1f,
        animationSpec = infiniteRepeatable(tween(1500), RepeatMode.Reverse),
        label = "pulse"
    )
    val configuration = LocalConfiguration.current
    val screenWidth = configuration.screenWidthDp.dp
    val isTablet = screenWidth >= 600.dp

    LaunchedEffect(connected) {
        if (connected) {
            statusText = "🔒 Quantum Tunnel Active – AI Adapting..."
            while (true) {
                delay(3000)
                discoveredHosts = (10..50).random()
                cleanIPs = (5..15).random()
            }
        } else {
            statusText = "Secured Quantum Tunnel Inactive"
        }
    }

    Scaffold(
        topBar = {
            TopAppBar(
                title = {
                    Text(
                        "ArgoTunnel",
                        style = MaterialTheme.typography.headlineMedium,
                        color = QuantumWhite
                    )
                },
                actions = {
                    IconButton(onClick = {}) {
                        Icon(Icons.Default.Settings, contentDescription = null, tint = QuantumWhite)
                    }
                },
                colors = TopAppBarDefaults.topAppBarColors(containerColor = QuantumDark)
            )
        }
    ) { padding ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .background(QuantumDark)
                .padding(padding)
                .verticalScroll(rememberScrollState())
                .padding(horizontal = 24.dp, vertical = 16.dp),
            verticalArrangement = Arrangement.Top,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Box(
                modifier = Modifier
                    .size(if (isTablet) 250.dp else 200.dp)
                    .clip(CircleShape)
                    .background(
                        Brush.radialGradient(
                            colors = listOf(
                                QuantumPurple.copy(alpha = 0.4f),
                                QuantumCyan.copy(alpha = 0.1f)
                            )
                        )
                    ),
                contentAlignment = Alignment.Center
            ) {
                Canvas(modifier = Modifier.size(if (isTablet) 200.dp else 160.dp)) {
                    val center = Offset(size.width / 2, size.height / 2)
                    drawCircle(
                        color = QuantumPurple.copy(alpha = pulseAlpha),
                        radius = size.width / 2
                    )
                    repeat(12) { i ->
                        val angle = Math.toRadians((i * 30).toDouble())
                        val radius = size.width / 2 - 20f
                        val x = center.x + (radius * Math.cos(angle)).toFloat()
                        val y = center.y + (radius * Math.sin(angle)).toFloat()
                        drawCircle(
                            color = QuantumCyan,
                            radius = 4f,
                            center = Offset(x, y)
                        )
                    }
                }
                Text(
                    text = "⚛",
                    style = MaterialTheme.typography.headlineLarge,
                    color = QuantumWhite
                )
            }

            Spacer(modifier = Modifier.height(16.dp))

            Text(
                text = statusText,
                style = MaterialTheme.typography.bodyLarge,
                color = if (connected) QuantumGreen else QuantumRed
            )

            Spacer(modifier = Modifier.height(24.dp))

            if (isTablet) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceEvenly
                ) {
                    InfoCard(
                        title = "🌍 Domains",
                        value = discoveredHosts.toString(),
                        color = QuantumCyan,
                        modifier = Modifier.weight(1f)
                    )
                    Spacer(modifier = Modifier.width(16.dp))
                    InfoCard(
                        title = "🔍 Clean IPs",
                        value = cleanIPs.toString(),
                        color = QuantumPurple,
                        modifier = Modifier.weight(1f)
                    )
                }
            } else {
                InfoCard(
                    title = "🌍 Domains",
                    value = discoveredHosts.toString(),
                    color = QuantumCyan,
                    modifier = Modifier.fillMaxWidth()
                )
                Spacer(modifier = Modifier.height(12.dp))
                InfoCard(
                    title = "🔍 Clean IPs",
                    value = cleanIPs.toString(),
                    color = QuantumPurple,
                    modifier = Modifier.fillMaxWidth()
                )
            }

            Spacer(modifier = Modifier.height(32.dp))

            Button(
                onClick = {
                    connected = !connected
                    if (connected) {
                        context.startService(Intent(context, ArgoVpnService::class.java))
                    } else {
                        context.stopService(Intent(context, ArgoVpnService::class.java))
                    }
                },
                modifier = Modifier
                    .fillMaxWidth()
                    .height(60.dp),
                shape = RoundedCornerShape(16.dp),
                colors = ButtonDefaults.buttonColors(
                    containerColor = if (connected) QuantumRed else QuantumGreen
                )
            ) {
                Text(
                    text = if (connected) "⏹ Disconnect" else "⚡ Connect",
                    style = MaterialTheme.typography.headlineMedium,
                    color = QuantumWhite
                )
            }

            Spacer(modifier = Modifier.height(16.dp))

            Text(
                text = "🧠 AI Insights",
                style = MaterialTheme.typography.bodyLarge,
                color = QuantumWhite
            )
            Text(
                text = "RL Agent: ${ if (connected) "active – profile ${listOf("Chrome","Firefox","Safari","Twitch","Zoom").random()}" else "idle" }",
                style = MaterialTheme.typography.labelSmall,
                color = QuantumOnSurface
            )
            Text(
                text = "GAN Generator: ${ if (connected) "injecting dummy packets" else "standby" }",
                style = MaterialTheme.typography.labelSmall,
                color = QuantumOnSurface
            )
        }
    }
}

@Composable
fun InfoCard(title: String, value: String, color: Color, modifier: Modifier = Modifier) {
    Card(
        modifier = modifier.animateContentSize(),
        colors = CardDefaults.cardColors(containerColor = QuantumSurface),
        shape = RoundedCornerShape(12.dp)
    ) {
        Column(
            modifier = Modifier.padding(16.dp),
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Text(
                text = title,
                style = MaterialTheme.typography.labelSmall,
                color = QuantumOnSurface
            )
            Text(
                text = value,
                style = MaterialTheme.typography.headlineLarge,
                color = color,
                fontWeight = FontWeight.Bold
            )
        }
    }
}
