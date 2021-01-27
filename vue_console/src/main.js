import Vue from 'vue'
import App from './App.vue'
import i18n from './i18n';
import vuetify from './plugins/vuetify';
import store from './store'
import router from './router'
import './permission'
import '@/components/index.js'
import '@/components/base.js'
import '@/plugins/vee-validate.js'
import '@/styles/overrides.sass'


Vue.config.productionTip = false
new Vue({
  vuetify,
  store,
  i18n,
  router,
  render: h => h(App)
}).$mount('#app')
