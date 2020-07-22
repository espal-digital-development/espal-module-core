import { createApp } from 'vue'
import App from './App'
import { storageService } from './services'

// TODO :: for now it's here, will be moved later
storageService.keepALive = true

createApp(App).mount('#app')
