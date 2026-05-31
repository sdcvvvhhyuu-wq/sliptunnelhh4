package com.slipnet.app

import android.app.DownloadManager
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.IntentFilter
import android.net.Uri
import android.os.Environment
import android.widget.Toast
import androidx.core.content.FileProvider
import com.google.gson.Gson
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import java.io.File
import java.net.URL

data class Release(val tag_name: String, val assets: List<Asset>)
data class Asset(val browser_download_url: String, val name: String)

object OTAUpdater {
    suspend fun checkForUpdates(context: Context): String? = withContext(Dispatchers.IO) {
        try {
            val url = URL("https://api.github.com/repos/sdcvvvhhyuu-wq/argotunnel/releases/latest")
            val json = url.readText()
            val release = Gson().fromJson(json, Release::class.java)
            val currentVersion = "v" + context.packageManager.getPackageInfo(context.packageName, 0).versionName
            if (release.tag_name != currentVersion) {
                release.assets.find { it.name.endsWith(".apk") }?.browser_download_url
            } else null
        } catch (e: Exception) {
            null
        }
    }

    fun downloadAndInstall(context: Context, downloadUrl: String) {
        val dm = context.getSystemService(Context.DOWNLOAD_SERVICE) as DownloadManager
        val request = DownloadManager.Request(Uri.parse(downloadUrl))
            .setTitle("SlipNet Update")
            .setDescription("Downloading latest version...")
            .setDestinationInExternalPublicDir(Environment.DIRECTORY_DOWNLOADS, "SlipNet_update.apk")
            .setNotificationVisibility(DownloadManager.Request.VISIBILITY_VISIBLE_NOTIFY_COMPLETED)
        val downloadId = dm.enqueue(request)

        val onComplete = object : BroadcastReceiver() {
            override fun onReceive(context: Context, intent: Intent) {
                val id = intent.getLongExtra(DownloadManager.EXTRA_DOWNLOAD_ID, -1)
                if (id == downloadId) {
                    val file = File(
                        Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS),
                        "SlipNet_update.apk"
                    )
                    if (file.exists()) {
                        val apkUri = FileProvider.getUriForFile(context, "${context.packageName}.provider", file)
                        val installIntent = Intent(Intent.ACTION_VIEW).apply {
                            setDataAndType(apkUri, "application/vnd.android.package-archive")
                            flags = Intent.FLAG_GRANT_READ_URI_PERMISSION or Intent.FLAG_ACTIVITY_NEW_TASK
                        }
                        context.startActivity(installIntent)
                    }
                }
            }
        }
        context.registerReceiver(onComplete, IntentFilter(DownloadManager.ACTION_DOWNLOAD_COMPLETE))
    }
}
