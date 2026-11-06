/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Types
import type { App } from 'vue'

// Plugins
import vuetify from './vuetify'
import router from '../router/index'

export function registerPlugins (app: App) {
  app
    .use(vuetify)
    .use(router) // Le decimos a la app que use las rutas
}