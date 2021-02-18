<template>
  <div class="login">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="登录" class="mt-5">
          <b-form>
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
              <b-button variant="outline-primary" block @click="login">登录</b-button>
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

export default {
  // 绑定数据
  data() {
    return {
      user: {
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
    login() {
      // if (this.user.telephone.length !== 11) {
      //   // this.showTelephoneValidate = true;
      //   this.validation = false;
      //   return;
      // }
      // this.validation = true;
      console.log('login');
    },
  },

};
</script>

<style>

</style>
