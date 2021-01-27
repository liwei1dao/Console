import { getLanguge, setLanguge ,getDark,setDark} from '@/utils/auth'

const state = {
    appname:"Lego",
    currlanguage: getLanguge() ? getLanguge()  :'zh',
    dark: (getDark() == "true") ? true : false,
 }
 
 const mutations = {
  Set_currlanguage: (state,lang) => {
    state.currlanguage = lang
    setLanguge(lang)
  },
  Set_currdark: (state,dark) => {
      state.dark = dark
      setDark(dark)
  },
 }
 
 const actions = {
  setcurrlanguage({ commit }, lang) {
    commit('Set_currlanguage',lang)
  },
  setcurrdark({ commit }, dark) {
      commit('Set_currdark',dark)
  }
 }
 
 export default {
   namespaced: true,
   state,
   mutations,
   actions
 }