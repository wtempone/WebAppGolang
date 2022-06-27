<template>
  <v-card class="mx-auto" max-width="400">
    <v-list-item two-line>
      <v-list-item-content>
        <v-list-item-title class="text-h5">
          Cadastre-se
        </v-list-item-title>
        <v-list-item-subtitle>Digite os dados para criacção do seu usuário</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-card-text>
      <validation-observer ref="observer" v-slot="{ invalid }">
        <form @submit.prevent="submit">
          <validation-provider v-slot="{ errors }" name="nome" rules="required|max:100|min:3">
            <v-text-field v-model="formDados.nome" :counter="100" :error-messages="errors" label="Nome" required>
            </v-text-field>
          </validation-provider>

          <validation-provider v-slot="{ errors }" name="email" rules="required|email">
            <v-text-field v-model="formDados.email" :error-messages="errors" label="E-mail" required></v-text-field>
          </validation-provider>

          <validation-provider v-slot="{ errors }" name="senha" rules="required|min:4">
            <v-text-field v-model="formDados.senha" :counter="8" :error-messages="errors" label="Senha" type=password
              required>
            </v-text-field>
          </validation-provider>

          <validation-provider v-slot="{ errors }" rules="required" name="checkbox">
            <v-checkbox v-model="formDados.aceiteTermos" :error-messages="errors" value="1" label="Aceito termos de uso"
              type="checkbox" required>
            </v-checkbox>
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
      <router-link to="login" replace>
        Entrar com senha
      </router-link>


    </v-list-item>
  </v-card>

</template>

<script>
import FormLogin from "../components/FormLogin.vue"
import { login, registro } from "@/services/auth"
import { mapActions, mapGetters } from "vuex";

export default {
  data: () => ({
    formDados: {
      nome: '',
      email: '',
      senha: '',
      aceiteTermos: 0
    },
    formErro: '',
  }),
  components: { FormLogin },

  methods: {
    ...mapActions('usuario', ['atualizaUsuario']),

    submit() {
      this.$refs.observer.validate().then(() => {
        registro(this.formDados.nome, this.formDados.email, this.formDados.senha).then(() => {
          login(this.formDados.email, this.formDados.senha).then(() => {
            this.atualizaUsuario().then(() => {
              this.$router.push('/home')
            })
          })
        }).catch((err) => {
          this.formErro = err.response.data.message;
        })
      })
    },
    clear() {
      this.formDados.nome = ''
      this.formDados.email = ''
      this.formDados.senha = ''
      this.formDados.aceiteTermos = ''
      this.formErro = ''
      this.$refs.observer.reset()
    },
  },
}
</script>