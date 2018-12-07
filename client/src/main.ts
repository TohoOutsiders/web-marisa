import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

if (process.env.NODE_ENV === 'production') {
  console.log = () => {
    return false
  }
  Vue.config.productionTip = false
} else {
  Vue.config.devtools = true
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
