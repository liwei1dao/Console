import Vue from 'vue';
import VueI18n from 'vue-i18n'
import { getLanguge} from '@/utils/auth'

Vue.use(VueI18n)

const messages={
    en:{
      ... require('./lang/en.json'),
    },
    zh:{
      ... require('./lang/zh.json'),
    }
}

const i18n = new VueI18n({
    locale: getLanguge() ? getLanguge() :'zh', // 设置语种
    messages: messages
})
export default i18n