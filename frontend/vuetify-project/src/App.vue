<template>
  <v-app>
    <!-- Panel Lateral -->
    <v-navigation-drawer v-model="drawer" app>
      <v-list>
        <v-list-item prepend-avatar="https://randomuser.me/api/portraits/men/85.jpg" title="John Smith" subtitle="Logged in"></v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-home" title="Home" to="/"></v-list-item>
        <v-list-item prepend-icon="mdi-book" title="Books" to="/books"></v-list-item>
        <v-list-item prepend-icon="mdi-mouse" title="Mice" to="/mice"></v-list-item>
        <v-list-item prepend-icon="mdi-qrcode-scan" title="Scanner" to="/scanner"></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <!-- Barra Superior -->
    <v-app-bar elevation="1">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <!-- Aquí se cargarán las demás páginas dinámicamente -->
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref } from 'vue'
const drawer = ref(true)
</script>