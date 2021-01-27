import Cookies from 'js-cookie'

/*
用户Token
*/
const TokenKey = 'ShareR-Token'
export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

/*
页面Languge
*/
const LangugeKey = 'ShareR-Languge'
export function getLanguge() {
  return Cookies.get(LangugeKey)
}
  
export function setLanguge(lang) {
  return Cookies.set(LangugeKey, lang)
}

export function removeLanguge() {
  return Cookies.remove(LangugeKey)
}

/*
页面Languge
*/
const DarkKey = 'ShareR-Dark'
export function getDark() {
  return Cookies.get(DarkKey)
}
  
export function setDark(dark) {
  return Cookies.set(DarkKey, dark)
}

export function removeDark() {
  return Cookies.remove(DarkKey)
}

