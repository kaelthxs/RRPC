import { createStore } from 'vuex';

export default createStore({
  state: {
    authVisible: false,
    user: null
  },
  mutations: {
    SET_AUTH_VISIBLE(state, value) {
      state.authVisible = value;
    },
    SET_USER(state, user) {
      state.user = user;
    }
  },
  actions: {
    toggleAuthForm({ commit }, value) {
      commit('SET_AUTH_VISIBLE', value);
    },
    loginUser({ commit }, userData) {
      commit('SET_USER', userData);
    }
  },
  getters: {
    isAuthVisible: (state) => state.authVisible,
    getUser: (state) => state.user
  }
});
