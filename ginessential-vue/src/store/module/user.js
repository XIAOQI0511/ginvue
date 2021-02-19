import storageService from '@/service/storageService';
import userService from '../../service/userService';

const userModule = {
  namespaced: true,
  state: {

    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO)
      ? JSON.parse(storageService.get(storageService.USER_INFO)) : null,
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // 更新state
      state.userInfo = userInfo;
    },
  },
  actions: {
    register(context, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        // 请求
      // const api = 'http://localhost:1016/api/auth/register';
      // this.axios.post(api, { ...this.user }).then((res) => {
        userService.register({ name, telephone, password }).then((res) => {
          // 保存token
          // this.$store.commit('userModule/SET_TOKEN', res.data.data.token);
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
          // storageService.set(storageService.USER_TOKEN, res.data.data.token);
          // userService.info().then((response) => {
        }).then((res) => {
          // 保存用户信息
          // this.$store.commit('userModule/SET_USERINFO', response.data.data.user);
          context.commit('SET_USERINFO', res.data.data.user);
          // storageService.set(storageService.USER_INFO,
          //   JSON.stringify(response.data.data.user));// 序列化

          // 跳转主页
          // this.$router.replace({ name: 'Home' });
          resolve(res);
        }).catch((err) => {
          reject(err);
          console.log('err:', err.response);
        });
      });
    },
    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        // 请求
        userService.login({ telephone, password }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user);
          // 跳转主页
          resolve(res);
        }).catch((err) => {
          reject(err);
          console.log('err:', err.response);
        });
      });
    },
    logout(context) {
      // 清除token
      context.commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清除用户信息
      context.commit('SET_USERINFO', '');
      storageService.set(storageService.USER_INFO, '');
      window.location.reload();
    },
  },
};

export default userModule;
