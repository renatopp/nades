import { createMemoryHistory, createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', component: () => import('./views/Home.vue') },
    { path: '/map/:mapId', component: () => import('./views/Map.vue') },
  ]
})

window.router = router

export default router