<template>
    <div>
        <v-list-item>
            <v-list-item-content>
                <v-list-item-title class="text-h6">
                    Clube
                </v-list-item-title>
                <v-list-item-subtitle>
                    Teste conhecimentos Golang
                </v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list-item v-if="usuario">
            <v-list-item-avatar>
                <v-img></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
                <v-list-item-title v-html="usuario.nome"></v-list-item-title>
                <v-list-item-subtitle v-html="usuario.email"></v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>
        <v-list dense nav>
            <router-link v-for="item in items" :to="item.link" replace>
                <v-list-item :key="item.title" link>
                    <v-list-item-icon>
                        <v-icon>{{ item.icon }}</v-icon>
                    </v-list-item-icon>

                    <v-list-item-content>
                        <v-list-item-title>{{ item.title }}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </router-link>
        </v-list>
        <div class="pa-2">
            <v-btn @click="login" block>
                {{ usuario ? "Sair" : "Login" }}
            </v-btn>
        </div>
    </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import { logout } from '@/services/auth';

export default {
    data() {
        return {
            items: [
                // { title: 'Home', icon: 'mdi-home', link: "/home" },
                // { title: 'Home', icon: 'mdi-home', link: "/home" },
                { title: 'Voos', icon: 'mdi-paragliding', link: "/voos" },
            ],
            right: null,
        }
    },
    computed: {
        ...mapGetters('usuario', ['usuario'])
    },
    methods: {
        ...mapActions('usuario', ['atualizaUsuario', 'logout']),
        login() {
            if (this.usuario) {
                logout().then(() => {
                    this.logout()
                    this.$router.push("/login")
                })
            } else {
                this.$router.push("/login")
            }
        }
    }
}
</script>