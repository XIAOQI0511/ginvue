import Vue from 'vue';
import Vuex from 'vuex';
import userModule from './module/user';

Vue.use(Vuex);

export default new Vuex.Store({
  // 开启严格模式
  strict: process.env.NODE_ENV !== 'production',
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    userModule,
  },
});
