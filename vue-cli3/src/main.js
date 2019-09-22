import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import router from './router.js'
import store from './store.js'
new Vue({
    router,
    store,
    render: h => h(App),
}).$mount('#app')
