import Vue from 'vue';
import Vuex from 'vuex';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';
import axios from 'axios';
import VueAxios from 'vue-axios';
import App from './App.vue';
import router from './router';
import store from './store';

// scss sytle
import './assets/sass/index.scss';

Vue.config.productionTip = false;

// install bootsrapvue
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
// vuelidate
Vue.use(Vuelidate);
// axios
Vue.use(VueAxios, axios);
Vue.use(Vuex);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
