import Vue from 'vue'
import VueRouter from 'vue-router'
import Hello from '../components/Hello.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Hello',
    component: Hello
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '../components/Login.vue')
  },
  {
    path: '/signup',
    name: 'Signup',
    component: () => import(/* webpackChunkName: "signup" */ '../components/Signup.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
