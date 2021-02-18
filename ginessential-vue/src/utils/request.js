import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
  // process是nodejs维护的全局变量，描述当前nodejs的进程信息
  // env保存当前node进程所有的环境变量
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 1000 * 5,
  // 会返回上一次的token
  // headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

// Add a request interceptor
// 拦截后动态配置token
service.interceptors.request.use((config) => {
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  // Do something before request is sent
  return config;
}, (error) => {
  // Do something with request error
  return Promise.reject(error);
});

export default service;
