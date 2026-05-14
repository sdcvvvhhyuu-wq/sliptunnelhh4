package com.slipnet.app

import android.content.Context
import com.google.gson.Gson
import com.google.gson.reflect.TypeToken

data class Profile(
    val name: String,
    val tunnelType: String,
    val server: String,
    val port: Int,
    val sni: String? = null,
    val payload: String? = null
)

object ProfileManager {
    private const val PREFS_NAME = "slipnet_profiles"

    fun getProfiles(context: Context): MutableList<Profile> {
        val json = context.getSharedPreferences(PREFS_NAME, Context.MODE_PRIVATE)
            .getString("list", "[]") ?: "[]"
        return Gson().fromJson(json, object : TypeToken<MutableList<Profile>>() {}.type)
    }

    fun saveProfiles(context: Context, profiles: List<Profile>) {
        context.getSharedPreferences(PREFS_NAME, Context.MODE_PRIVATE)
            .edit().putString("list", Gson().toJson(profiles)).apply()
    }
}
