import axios from 'axios'
import store from '@/store'
import router from '@/router'
import Message from '@/components/Message/index.js'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
    timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
    config => {
        if (store.getters.token) {
            config.headers['X-Token'] = getToken()
        }
        return config
    },
    error => {
        console.log(error) // for debug
        return Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    response => {
        const res = response.data
        console.log('response:' + JSON.stringify(response.baseURL) + "|res:" + JSON.stringify(res))
        if (res.code !== 0) {
            if (res.code == 16 ){
                async() => {
                    await store.dispatch('user/resetToken')
                    router.push({ path:'/base'})
                }
            }else{
                Message.error(res.message || 'Error')
            }
            return Promise.reject(new Error(res.message || 'Error'))
        } else {
            return res
        }
    },
    error => {
        Message.error(error.message)
        return Promise.reject(error)
    }
)
  
export default service