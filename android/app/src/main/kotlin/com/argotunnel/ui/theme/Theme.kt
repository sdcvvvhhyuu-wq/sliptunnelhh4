package com.argotunnel.ui.theme

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.darkColorScheme
import androidx.compose.runtime.Composable

private val DarkColorScheme = darkColorScheme(
    primary = QuantumPurple,
    secondary = QuantumCyan,
    tertiary = QuantumGreen,
    background = QuantumDark,
    surface = QuantumSurface,
    onPrimary = QuantumWhite,
    onSecondary = QuantumDark,
    onTertiary = QuantumDark,
    onBackground = QuantumWhite,
    onSurface = QuantumOnSurface,
    error = QuantumRed
)

@Composable
fun ArgoTunnelTheme(content: @Composable () -> Unit) {
    MaterialTheme(
        colorScheme = DarkColorScheme,
        typography = Typography,
        content = content
    )
}
