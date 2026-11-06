<!-- ARCHIVO: src/pages/books.vue -->
<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="books"
      item-key="title"
      class="elevation-1"
    >
      <!-- Cabecera de la tabla con el título y botón -->
      <template v-slot:top>
        <v-toolbar flat color="transparent">
          <v-toolbar-title>Libros Populares</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          
          <!-- MODAL (DIÁLOGO) PARA AGREGAR O EDITAR UN LIBRO -->
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ props }">
              <v-btn color="primary" variant="tonal" prepend-icon="mdi-plus" v-bind="props" @click="resetForm">
                Añadir Libro
              </v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="text-h5">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="12">
                      <v-text-field v-model="editedItem.title" label="Título"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-text-field v-model="editedItem.author" label="Autor"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-select v-model="editedItem.genre" :items="['Ficción', 'No-Ficción', 'Distópico', 'Sci-Fi']" label="Género"></v-select>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-text-field v-model="editedItem.year" label="Año" type="number"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-text-field v-model="editedItem.pages" label="Páginas" type="number"></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue-darken-1" variant="text" @click="close">Cancelar</v-btn>
                <v-btn color="blue-darken-1" variant="text" @click="save">Guardar</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>

      <!-- Slot para personalizar la columna de Acciones -->
      <template v-slot:item.actions="{ item }">
        <v-icon size="small" class="me-2" color="blue" @click="editBook(item)">
          mdi-pencil
        </v-icon>
        <v-icon size="small" color="red" @click="deleteBook(item)">
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>

    <!-- NOTIFICACIÓN PEQUEÑA EN LA ESQUINA SUPERIOR DERECHA -->
    <v-snackbar
      v-model="snackbar"
      location="top right"
      color="success"
      :timeout="3000"
    >
      Libro eliminado correctamente.
      <template v-slot:actions>
        <v-btn color="white" icon="mdi-close" variant="text" @click="snackbar = false"></v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref, computed } from 'vue'
import Swal from 'sweetalert2'

// Controla si la ventana emergente está abierta o cerrada
const dialog = ref(false)

// Controla si la notificación de la esquina se muestra o no
const snackbar = ref(false)

// Cabeceras traducidas
const headers = [
  { title: 'Título', align: 'start', key: 'title' },
  { title: 'Autor', key: 'author' },
  { title: 'Género', key: 'genre' },
  { title: 'Año', key: 'year' },
  { title: 'Páginas', key: 'pages' },
  { title: 'Acciones', key: 'actions', sortable: false },
]

// Lista de libros con géneros ajustados
const books = ref([
  { title: 'Matar a un ruiseñor', author: 'Harper Lee', genre: 'Ficción', year: 1960, pages: 281 },
  { title: '1984', author: 'George Orwell', genre: 'Distópico', year: 1949, pages: 328 },
  { title: 'El Gran Gatsby', author: 'F. Scott Fitzgerald', genre: 'Ficción', year: 1925, pages: 180 },
  { title: 'Sapiens', author: 'Yuval Noah Harari', genre: 'No-Ficción', year: 2011, pages: 443 },
  { title: 'Dune', author: 'Frank Herbert', genre: 'Sci-Fi', year: 1965, pages: 412 }
])

// Variables para manejar la edición
const editedIndex = ref(-1)
const editedItem = ref({ title: '', author: '', genre: '', year: null, pages: null })
const defaultItem = { title: '', author: '', genre: '', year: null, pages: null }

// Título dinámico para el modal
const formTitle = computed(() => {
  return editedIndex.value === -1 ? 'Añadir Libro' : 'Editar Libro'
})

// Limpiar el formulario al agregar uno nuevo
const resetForm = () => {
  editedItem.value = Object.assign({}, defaultItem)
  editedIndex.value = -1
}

// Cargar los datos del libro seleccionado al formulario
const editBook = (item) => {
  editedIndex.value = books.value.indexOf(item)
  editedItem.value = Object.assign({}, item)
  dialog.value = true
}

// Confirmar y borrar el libro
const deleteBook = (item) => {
  const index = books.value.indexOf(item)
  
  Swal.fire({
    title: '¿Desea eliminar el libro?',
    text: item.title,
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Sí, eliminar',
    cancelButtonText: 'Cancelar'
  }).then((result) => {
    if (result.isConfirmed) {
      books.value.splice(index, 1)
      // Mostrar la notificación web pequeña en la esquina
      snackbar.value = true
    }
  })
}

// Cerrar el modal
const close = () => {
  dialog.value = false
  resetForm()
}

// Guardar los cambios
const save = () => {
  if (editedIndex.value > -1) {
    Object.assign(books.value[editedIndex.value], editedItem.value)
  } else {
    books.value.push(Object.assign({}, editedItem.value))
  }
  close()
}
</script>