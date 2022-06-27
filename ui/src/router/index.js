import Vue from 'vue'
import VueRouter from 'vue-router'
import FormRegistro from '../components/FormRegistro.vue'
import FormLogin from "../components/FormLogin.vue"
import ListaVoos from "../components/ListaVoos.vue"
import AreaLogada from "../views/AreaLogada.vue"
import store from '@/store/index';
import DetalhesVoo from "@/components/DetalhesVoo"
import Mapa from "@/components/Mapa"
import Mapa3d from "@/components/Mapa3d"
Vue.use(VueRouter)


const routes = [
  {
    path: '/teste',
    name: 'teste',
    component: Mapa3d
  },  
  // {
  //   path: '/voo',
  //   name: 'voo',
  //   component: ListaVoos
  // },
  // {
  //   path: '/voo/:link',
  //   name: 'voo',
  //   component: DetalhesVoo,
  //   props: true
  // },
  // {
  //   path: '/registro',
  //   name: 'registro',
  //   component: FormRegistro
  // },
  // {
  //   path: '/login',
  //   name: 'login',
  //   component: FormLogin
  // },
  // {
  //   path: '/home',
  //   name: 'home',
  //   component: AreaLogada,
  // },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

router.beforeEach(async (to, from, next) => {
  // await store.dispatch('usuario/atualizaUsuario');
  // if (to.name != 'teste' && to.name != 'login' && to.name != 'registro' && !store.getters['usuario/usuario']) {
  //   console.log('nao autenticado')
  //   next('/login')

  // } else {
  //   next()
  // }
})