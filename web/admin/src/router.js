import Vue from 'vue'
import Router from 'vue-router'
import Cartoon from './views/Cartoon.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'cartoon',
      component: Cartoon
    }
  ]
})
