<template>
  <v-card class="mt-n120 dialog-form mx-auto" elevation="24" max-width="400">
    <v-list-item two-line>
      <v-list-item-content>
        <v-list-item-title class="text-h5">
          Login
        </v-list-item-title>
        <v-list-item-subtitle>Digite os dados para entrar com seu usuário</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-card-text>
      <validation-observer ref="observer" v-slot="{ invalid }">
        <form @submit.prevent="submit">
          <validation-provider v-slot="{ errors }" name="email" rules="required|email">
            <v-text-field v-model="formDados.email" :error-messages="errors" label="E-mail" required></v-text-field>
          </validation-provider>

          <validation-provider v-slot="{ errors }" name="senha" rules="required|min:4">
            <v-text-field v-model="formDados.senha" :counter="8" :error-messages="errors" label="Senha" type=password
              required>
            </v-text-field>
          </validation-provider>
          <p class="error--text" v-if="formErro">
            {{ formErro }}
          </p>
          <v-btn class="mr-4" type="submit" :disabled="invalid">
            enviar
          </v-btn>
          <v-btn @click="clear">
            limpar
          </v-btn>
        </form>
      </validation-observer>
    </v-card-text>

    <v-list-item>
      <router-link to="registro" replace>
        Cadastre-se
      </router-link>

    </v-list-item>
  </v-card>

</template>

<script>

import FormRegistro from './FormRegistro.vue'
import { login } from '@/services/auth';
import { mapActions, mapGetters } from "vuex";

export default {
  data: () => ({
    formDados: {
      email: '',
      senha: '',
    },
    carregando: false,
    formErro: '',
  }),
  components: { FormRegistro },
  methods: {
    ...mapActions('usuario', ['atualizaUsuario']),
    submit() {
      this.$refs.observer.validate().then(() => {
        this.formErro = ''
        login(this.formDados.email, this.formDados.senha)
          .then(() => {
            console.log('passou login)')
            this.atualizaUsuario().then(() => {
              this.$router.push('/map3d')
            })
          }).catch((err) => {
            this.formErro = err.response.data.message;
          })
      })
    },
    clear() {

      this.formDados.email = ''
      this.formDados.senha = ''
      this.formErro = ''
      this.$refs.observer.reset()
    },

  },
}
</script>
<style lang="scss">
.dialog-form { 
  margin-top: 120px!important;
}
</style>