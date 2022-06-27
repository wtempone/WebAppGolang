import { obterUsuario } from "@/services/usuario"

const state = {
  usuario: null
}

const getters = {
  usuario(state) {
    return state.usuario;
  }
}

const actions = {
  async atualizaUsuario({ commit }) {
    try {
      const response = await obterUsuario();
      commit('atualizaUsuario', response.data);
    } catch (error) {
      console.log('Erro ao atualizar o usuario - store usuarios')
    }
  },
  async logout({ commit }) {
    try {
      commit('atualizaUsuario', null);
    } catch (error) {
      console.log('Erro ao fazer logoff - store usuarios')
    }
  }
}


const mutations = {
  atualizaUsuario(state, data) {
    console.log('atualizaUsuario', data)
    state.usuario = data;
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}