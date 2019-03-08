import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);
Vue.config.productionTip = false

// http://127.0.0.1:4321/cartoon/resource

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
