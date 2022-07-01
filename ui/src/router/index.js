import Vue from 'vue'
import VueRouter from 'vue-router'
import FormRegistro from '../components/FormRegistro.vue'
import FormLogin from "../components/FormLogin.vue"
import store from '@/store/index';
import Map3DControls from "@/components/Map3DControls.vue"
import ListaVoos from "@/components/ListaVoos.vue"
Vue.use(VueRouter)
const routes = [
  {
    path: '/map3d',
    name: 'map3d',
    component: Map3DControls
  },
  {
    path: '/voos',
    name: 'voos',
    component: ListaVoos
  },
  {
    path: '/',
    name: '',
    component: Map3DControls
  },
  {
    path: '/login',
    name: 'login',
    component: FormLogin
  },
  {
    path: '/registro',
    name: 'registro',
    component: FormRegistro
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

router.beforeEach(async (to, from, next) => {
  await store.dispatch('usuario/atualizaUsuario');
  console.log("navegando: ",  to.name , store.getters['usuario/usuario'])
  if (to.name != '' && to.name != 'map3d' && to.name != 'login' && to.name != 'registro' && !store.getters['usuario/usuario']) {
    console.log('nao autenticado')
    next('/login')
  } else {
    next()
  }
})