import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VeeValidate ,{ ValidationProvider,ValidationObserver, Validator, localize} from 'vee-validate/dist/vee-validate.full.esm';
import pt_BR from 'vee-validate/dist/locale/pt_BR';
import vuetify from './plugins/vuetify'
import VueAxios from 'vue-axios'
import * as VueGoogleMaps from "vue2-google-maps" 
Vue.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyANtbO23v2yrapT-ASnjchh7_lRYYF3CxM",
    libraries: "places"
  }
});
Vue.component('ValidationProvider', ValidationProvider);
Vue.component('ValidationObserver', ValidationObserver);
Vue.component('Validator', Validator);
Vue.component('localize', localize);
Vue.component('VueAxios', VueAxios);

Vue.config.productionTip = false
localize('pt_BR', pt_BR);

new Vue({
  router,
  store,
  VeeValidate,
  vuetify,
  VueAxios,
  render: h => h(App)
}).$mount('#app')
