<template>
  <div class="register">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册" class="mt-5">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入您的名称（选填）"
              ></b-form-input>
            </b-form-group>
            <b-form-group label="手机号">
              <!-- model来自data(){}中定义-->
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="输入您的手机号"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合要求
              </b-form-invalid-feedback>
              <!-- <b-form-text id="password-help-block" text-variant="danger"
              v-if="showTelephoneValidate">
                手机号必须为11位
              </b-form-text> -->
              <!-- <b-form-invalid-feedback :state="validation">
                手机号必须为11位
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                手机号符合
              </b-form-valid-feedback> -->
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入您的密码"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group >
              <b-button variant="outline-primary" block @click="register">注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>

</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import userService from '@/service/userService';

export default {
  // 绑定数据
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      // 定义属性，控制是否显示
      // showTelephoneValidate: false,
      // validation: null, // 初始时不验证
    };
  },
  // $v对象在此定义后生成属性
  validations: {
    user: {
      name: {

      },
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    validateState(name) {
      // es6解构赋值
      // 初始化时$dirty=false，与表单交互时$dirty=true，返回error值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null; // 无交互时给state赋null
    },
    // 绑定方法提交注册
    register() {
      // 验证数据
      this.$v.user.$touch();// 将$dirty设为true
      if (this.$v.user.$anyError) {
        return;
      }
      // const api = 'http://localhost:1016/api/auth/register';
      // this.axios.post(api, { ...this.user }).then((res) => {
      userService.register(this.user).then((res) => {
        // 保存token
        this.$store.commit('userModule/SET_TOKEN', res.data.data.token);
        // storageService.set(storageService.USER_TOKEN, res.data.data.token);
        userService.info().then((response) => {
          // 保存用户信息
          this.$store.commit('userModule/SET_USERINFO', response.data.data.user);
          // storageService.set(storageService.USER_INFO,
          //   JSON.stringify(response.data.data.user));// 序列化

          // 跳转主页
          this.$router.replace({ name: 'Home' });
        });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
        console.log('err:', err.response);
      });
      // 请求api
      // if (this.user.telephone.length !== 11) {
      //   // this.showTelephoneValidate = true;
      //   this.validation = false;
      //   return;
      // }
      // this.validation = true;
      console.log('register');
    },
  },

};
</script>

<style>

</style>
