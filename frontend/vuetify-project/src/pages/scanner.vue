<template>
  <v-container>
    <v-text-field 
      v-model="scannedResult" 
      label="Código escaneado" 
      readonly
      append-inner-icon="mdi-qrcode-scan"
      @click:append-inner="dialog = true"
    ></v-text-field>

    <v-dialog v-model="dialog" max-width="500">
      <v-card>
        <v-card-title>Escanear código</v-card-title>
        <v-card-text>
          <qrcode-stream v-if="dialog" @decode="onDecode"></qrcode-stream>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="dialog = false">Cerrar</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import { QrcodeStream } from 'vue-qrcode-reader'

const dialog = ref(false)
const scannedResult = ref('')

const onDecode = (result) => {
  scannedResult.value = result
  dialog.value = false
}
</script>