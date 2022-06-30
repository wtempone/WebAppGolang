<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <MenuPrincipal></MenuPrincipal>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Clube - Golang</v-toolbar-title>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import MenuPrincipal from './components/MenuPrincipal.vue';
import { mapActions, mapGetters } from "vuex";
import Map3DControls from './components/Map3DControls.vue';
import FormLogin from './components/FormLogin.vue';
import AreaLogada from './views/AreaLogada.vue';

export default {
  data: () => ({ drawer: null }),
  components: { Map3DControls, FormLogin, MenuPrincipal, AreaLogada }
}
</script>