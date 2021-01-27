<template>
    <v-card >
        <v-toolbar
        color="primary"
        flat
        >
            <v-toolbar-title>{{$t('common.register')}}</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
            <v-form>
                <v-row justify="center">
                    <v-col cols="10">
                        <v-text-field
                            v-model="form.Email"
                            type="email"
                            :rules="emailRules"
                            prepend-icon="mdi-email"
                            :label="$t('common.InputPhoneOrEamil')"
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
                                    :label="$t('common.InputVerification')"
                                    clearable
                                    required
                                ></v-text-field>
                            </v-col>
                            <v-col cols="4">
                                 <v-btn rounded color="primary" @click="getcaptcha" :disabled="codeDisabled">{{Verification}}</v-btn>
                            </v-col>
                        </v-row>
                    </v-col>
                    <v-col cols="10">
                        <v-text-field
                            v-model="form.password"
                            :rules="passwordRules"
                            prepend-icon="mdi-lock-question"
                            :append-icon="showPwd1 ? 'mdi-eye-off' : 'mdi-eye'"
                            :type="showPwd1 ? 'text' : 'password'"
                            :label="$t('common.InputPassword')"
                            @click:append="showPwd1 = !showPwd1"
                            required
                        ></v-text-field>
                    </v-col>
                    <v-col cols="10">
                        <v-text-field
                            v-model="form.confirmpassword"
                            :rules="passwordRules"
                            prepend-icon="mdi-lock-question"
                            :append-icon="showPwd2 ? 'mdi-eye-off' : 'mdi-eye'"
                            :type="showPwd2 ? 'text' : 'password'"
                            :label="$t('common.InputConfirmPassword')"
                            @click:append="showPwd2 = !showPwd2"
                            required
                        ></v-text-field>
                    </v-col>
                </v-row>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey" to="/base/login" >{{$t('common.login')}}</v-btn>
            <v-btn color="primary" @click="handle_register">{{$t('common.register')}}</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
import { sendemailcaptcha } from '@/api/api'
import { register } from '@/api/user'
export default {
  name: 'Register',
  components: {
  },
  data() {
    return {
      codeDisabled:false,
      countdown : 60,
      timer: null,
      getcaptchabutt: this.$i18n.t('common.Verification'),
      showPwd1: false,
      showPwd2: false,
      form: {
        Nickname: "",
        Password: "",
        Email:"",
        Captcha:"",
      },
      loading: false,
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+/.test(v) || 'E-mail must be valid',
      ],
      captchaRules: [
        v => !!v || 'Captcha is required',
        v => v.length == 6 || 'Captcha must 6 characters',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
        v => v.length >= 6 || 'Password must be less than 6 characters',
      ],
    };
  },
  computed: {
      Verification:function(){
        if (this.timer == null){
          return this.$i18n.t('common.Verification')
        }else{
          return this.$i18n.t('common.VerificationR')+ "(" + this.countdown + ")";
        }
      }
  },
  methods: {
    checkemail(){
      return 
    },

    getcaptcha(){
      this.codeDisabled = true
      sendemailcaptcha({Mailbox:this.form.Email,CaptchaType:0}).then(response => {
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
    handle_register() {
      register(this.form).then(response => {
        const {message} = response
        this.$message.success(message)
      }).catch(error => {
        this.$message.error(error.message)
      })
    },
  },
  created() {},
};
</script>

<style lang="sass" scoped>

</style>