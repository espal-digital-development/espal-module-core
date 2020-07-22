import { createApp } from 'vue'
import App from './App'
import { storageService, routerService } from './services'

// TODO :: for now it's here, will be moved later
storageService.keepALive = true

const app = createApp(App)
app.use(routerService._router)

app.mount('#app')
