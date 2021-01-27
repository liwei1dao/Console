<template>
    <v-card >
        <v-toolbar
        color="primary"
        flat
        >
        <v-toolbar-title>{{$t('login')}}</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
            <v-form>
                <v-row justify="center">
                   <v-col cols="10">
                        <v-text-field
                            v-model="form.PhonOrEmail"
                            :rules="emailRules"
                            prepend-icon="mdi-email"
                            :label="$t('inputphoneoreamil')"
                            clearable
                        ></v-text-field>
                   </v-col>
                   <v-col cols="10" align-self>
                        <v-row 
                            align=center
                        >
                            <v-col cols="8">
                                <v-text-field
                                    v-model="form.Captcha"
                                    :rules="captchaRules"
                                    prepend-icon="mdi-email-check"
                                    :label="$t('inputVerification')"
                                    clearable
                                    required
                                ></v-text-field>
                            </v-col>
                            <v-col cols="4">
                                 <v-btn class="text-capitalize" rounded color="primary" @click="getcaptcha" :disabled="codeDisabled">{{Verification}}</v-btn>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey text-capitalize" to="/base/register" >{{$t('register')}}</v-btn>
            <v-btn color="primary text-capitalize" @click="handle_login">{{$t('login')}}</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
import { sendemailcaptcha } from '@/api/api'
import { paramsign } from '@/utils/md5'
export default {
  name: 'Login',
  components: {
  },
  data() {
    return {
      showPwd: false,
      codeDisabled:false,
      timer :null,
      countdown:60,
      form: {
        PhonOrEmail: "",
        Captcha: "",
      },
      loader: null,
      loading: false,
      getcaptchabutt: this.$i18n.t('getverification'),
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+/.test(v) || 'E-mail must be valid',
      ],
      captchaRules: [
        v => !!v || 'Captcha is required',
        v => v.length == 6 || 'Captcha must  6 characters',
      ],
    };
  },
  computed: {
      Verification:function(){
        if (this.timer == null){
          return this.$i18n.t('getverification')
        }else{
          return this.countdown
        }
      }
  },
  methods: {
    //获取验证码
    getcaptcha(){
      if (!this.form.PhonOrEmail) {
         this.$message.error("PhonOrEmail Is Null")
        return;
      }
      this.codeDisabled = true
      sendemailcaptcha(paramsign({Mailbox:this.form.PhonOrEmail,CaptchaType:0})).then(response => {
        const {message} = response
          this.$message.success(message)
          if (!this.timer) {//启动计时器
              this.timer = setInterval(() => {
              if (this.countdown > 0 && this.countdown <= 60) {
                this.countdown--;
                if (this.countdown == 0) {
                  clearInterval(this.timer);
                  this.countdown = 60;
                  this.timer = null;
                  this.codeDisabled = false;
                }
              }
            }, 1000)
          }
      }).catch(error => {
        this.$message.error(error.message)
        this.codeDisabled = false
      })
    },
    //登录
    handle_login() {
      if (!this.form.PhonOrEmail || !this.form.Captcha) {
         this.$message.error("PhonOrEmail Or Captcha Is Null")
        return;
      }

      // loginbycaptcha(paramsign({ PhonOrEmail: this.form.PhonOrEmail, Captcha: this.form.Captcha})).then(response => {
      //   this.$message.success("Success")
      //   const {data} = response
      //   this.$store.commit("Set_token",data.Token)
      //   this.$store.commit("Set_info",data.UserData)
      //   this.$router.push({ path:'/service' })
      // })

      this.loading = true
      this.$store.dispatch('user/loginbycaptcha',{PhonOrEmail: this.form.PhonOrEmail, Captcha: this.form.Captcha})
        .then(() => {
          this.loading = false
          this.$router.push({ path:'/service' })
        })
        .catch((error) => {
          this.loading = false
           this.$message.error(error.message)
        }
      )
    },
  },
  created() {},
};
</script>

<style>
</style>