<template>
  <v-app id="inspire">
    <!-- <v-navigation-drawer v-if="usuario" v-model="drawer" app>
      <MenuPrincipal></MenuPrincipal>
    </v-navigation-drawer> -->

    <v-app-bar app>
      <v-app-bar-nav-icon  @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Clube - Golang</v-toolbar-title>
    </v-app-bar>
    <Mapa3d></Mapa3d>
    <!-- <v-main style="margin-top: 50px;">
      <router-view />
    </v-main> -->
  </v-app>
</template>

<script>
import MenuPrincipal from './components/MenuPrincipal.vue';
import { mapActions, mapGetters } from "vuex";
import Mapa3d from './components/Mapa3d2.vue';

export default {
    components: { Mapa3d }
}
</script>