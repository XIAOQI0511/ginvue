import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '../store';
import Home from '../views/Home.vue';
import userRoutes from './module/user';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // 匿名函数回调，惰性加载
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});
router.beforeEach((to, from, next) => {
  // to是要去的路由，from是从哪个路由过来
  // 判断是否需要登录
  if (to.meta.auth) {
    if (store.state.userModule.token) {
      // 这里还要判断token的有效性，比如有没有过期，需要后台发放token的时候带上token
      // 的有效期，如果token无效，需要请求token
      next();
    } else {
      router.push({ name: 'login' });
    }
  } else {
    next();
  }
});
export default router;
