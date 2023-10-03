import 'vite/modulepreload-polyfill'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// import * as bootstrap from './assets/bootstrap/js/bootstrap.esm.js';

import './assets/main.css'

const app = createApp(App)

app.use(router)

app.mount('#app')
