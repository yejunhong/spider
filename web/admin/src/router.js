import Vue from 'vue'
import Router from 'vue-router'
import Cartoon from './views/Cartoon.vue'
import Resource from './views/Resource.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/resource',
      name: 'resource',
      component: Resource
    },
    {
      path: '/cartoon',
      name: 'cartoon',
      component: Cartoon
    },
    {
      path: '/',
      redirect: '/cartoon'
    }
  ]
})
