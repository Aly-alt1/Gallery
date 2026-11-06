import { createRouter, createWebHistory } from 'vue-router'

import Home from '@/pages/index.vue'
import Books from '@/pages/books.vue'
import Mice from '@/pages/mice.vue'
import Scanner from '@/pages/scanner.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/books', component: Books },
  { path: '/mice', component: Mice },
  { path: '/scanner', component: Scanner },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router