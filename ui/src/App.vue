<template>
  <v-app id="inspire">
    <!-- <v-navigation-drawer v-if="usuario" v-model="drawer" app>
      <MenuPrincipal></MenuPrincipal>
    </v-navigation-drawer> -->

    <v-app-bar app>
      <v-app-bar-nav-icon  @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Clube - Golang</v-toolbar-title>
    </v-app-bar>
    <SingleMapa3d></SingleMapa3d>
    <!-- <v-main style="margin-top: 50px;">
      <router-view />
    </v-main> -->
  </v-app>
</template>

<script>
import MenuPrincipal from './components/MenuPrincipal.vue';
import { mapActions, mapGetters } from "vuex";
import SingleMapa3d from './components/SingleMapa3d.vue';

export default {
    components: { SingleMapa3d }
}
</script>