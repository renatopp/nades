import './style.css'
import 'primevue/resources/themes/lara-dark-amber/theme.css'
import 'primeflex/primeflex.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice';

import App from './App.vue'
import router from './routes'

const app = createApp(App)
app.use(router)
app.use(PrimeVue)
app.use(ConfirmationService)

app.mount('#app')
