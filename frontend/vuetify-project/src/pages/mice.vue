<!-- ARCHIVO: src/pages/mice.vue -->
<template>
  <v-container>
    <v-data-iterator :items="mice" :items-per-page="8" search="">
      
      <!-- HEADER: Título y botón superior -->
      <template v-slot:header>
        <div class="d-flex justify-space-between align-center mb-4">
          <h2 class="text-h5">Ratones Populares</h2>
          <v-btn color="primary" variant="tonal" prepend-icon="mdi-plus">
            Añadir Ratón
          </v-btn>
        </div>
      </template>

      <!-- DEFAULT: El ciclo que dibuja las tarjetas -->
      <template v-slot:default="{ items }">
        <v-row>
          <v-col v-for="item in items" :key="item.raw.name" cols="12" sm="6" md="4" lg="3">
            <v-card elevation="2">
              <div class="bg-grey-lighten-4 d-flex justify-center align-center py-6">
                <!-- Aquí usamos un icono MDI en lugar de imagen para simplificar -->
                <v-icon size="80" color="grey-darken-2">mdi-mouse</v-icon>
              </div>
              
              <v-card-title class="text-subtitle-1 font-weight-bold">
                {{ item.raw.name }}
              </v-card-title>
              
              <v-card-text>
                <div class="d-flex justify-space-between">
                  <span>Precio</span>
                  <span class="font-weight-bold text-primary">${{ item.raw.price }}</span>
                </div>
              </v-card-text>
              
              <v-divider></v-divider>
              
              <v-card-actions class="d-flex justify-space-between px-4">
                <v-icon color="blue" class="cursor-pointer" @click="editMouse(item.raw)">mdi-pencil</v-icon>
                <v-icon color="red" class="cursor-pointer" @click="confirmDelete(item.raw.name)">mdi-delete</v-icon>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </template>

      <!-- FOOTER: Controles de Paginación -->
      <template v-slot:footer="{ page, pageCount, prevPage, nextPage }">
        <div class="d-flex align-center justify-end pa-4 mt-4">
          <div class="text-caption me-2">Página {{ page }} de {{ pageCount }}</div>
          <v-btn
            :disabled="page === 1"
            icon="mdi-chevron-left"
            variant="text"
            density="comfortable"
            @click="prevPage"
          ></v-btn>
          <v-btn
            :disabled="page >= pageCount"
            icon="mdi-chevron-right"
            variant="text"
            density="comfortable"
            @click="nextPage"
          ></v-btn>
        </div>
      </template>
      
    </v-data-iterator>

    <!-- NOTIFICACIÓN PEQUEÑA EN LA ESQUINA SUPERIOR DERECHA -->
    <v-snackbar
      v-model="snackbar"
      location="top right"
      color="success"
      :timeout="3000"
    >
      Ratón eliminado correctamente.
      <template v-slot:actions>
        <v-btn color="white" icon="mdi-close" variant="text" @click="snackbar = false"></v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import Swal from 'sweetalert2'

// Controla si la notificación de la esquina se muestra o no
const snackbar = ref(false)

const mice = ref([
  { name: 'Logitech G Pro X', price: 149.99 },
  { name: 'Razer DeathAdder V2', price: 69.99 },
  { name: 'Corsair Dark Core RGB', price: 89.99 },
  { name: 'SteelSeries Rival 3', price: 29.99 },
  { name: 'HyperX Pulsefire FPS Pro', price: 44.99 },
  { name: 'Zowie EC2', price: 69.99 },
  { name: 'Roccat Kone AIMO', price: 79.99 },
  { name: 'Logitech G903', price: 119.99 },
  { name: 'Cooler Master MM711', price: 49.99 },
  { name: 'Glorious Model O', price: 49.99 }
])

const editMouse = (item) => {
  console.log('Editar:', item.name)
}

const confirmDelete = (name) => {
  Swal.fire({
    title: '¿Desea eliminar este ratón?',
    text: name,
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Sí, eliminar',
    cancelButtonText: 'Cancelar'
  }).then((result) => {
    if (result.isConfirmed) {
      // Filtrar y eliminar el ratón del arreglo si el usuario da en OK
      mice.value = mice.value.filter(mouse => mouse.name !== name)
      // Mostrar la notificación web pequeña en la esquina
      snackbar.value = true
    }
  })
}
</script>