import { loginbycaptcha,getuserinfo,loginout} from '@/api/user'
import { paramsign } from '@/utils/md5'
import { setToken,getToken,removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'

const state = {
  token : getToken(),
  userinfo : null
}

const mutations = {     
    Set_token:(state,token) => {
      state.token = token
    },
    Set_info: (state,userinfo) => {
      state.userinfo = userinfo
    },
}

const actions = {
  loginbycaptcha({commit},form) {
    console.log("loginbycaptcha:"+form)
    return new Promise((resolve, reject) => {
      loginbycaptcha(paramsign(form)).then(response => {
          const {data} = response
          commit("Set_token",data)
          setToken(data)
          resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  getuserinfo({commit}) {
    return new Promise((resolve, reject) => {
      console.log("loginbytoken:")
      getuserinfo().then(response => {
          const {data} = response
          commit("Set_info",data)
          setToken(data.Token)
          resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  loginout({commit}) {
    return new Promise((resolve, reject) => {
      loginout().then(() => {
        commit("Set_token","")
        commit('Set_info',null)
        removeToken()
        resetRouter()
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('Set_info', null)
      commit("Set_token","")
      removeToken()
      resolve()
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}