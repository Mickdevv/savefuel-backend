import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import en from './locales/en.json'
import fr from './locales/fr.json'
import Aura from '@primeuix/themes/aura'
import App from './App.vue'
import router from './router'
import { createI18n } from 'vue-i18n'

const app = createApp(App)
const i18n = createI18n({
  locale: 'fr',
  fallbackLocale: 'fr',
  globalInjection: true,
  messages: {
    en,
    fr,
  },
})
app.use(i18n)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {},
  },
})
app.use(createPinia())
app.use(router)

app.mount('#app')
