const userRoute = [
  {
    path: '/register',
    name: 'register',
    component: () => import(/* webpackChunkName: "about" */ '@/views/register/Register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "about" */ '@/views/login/Login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    // 给路由增加元信息，随便定义
    component: () => import(/* webpackChunkName: "about" */ '@/views/profile/Profile.vue'),
  },
];

export default userRoute;
