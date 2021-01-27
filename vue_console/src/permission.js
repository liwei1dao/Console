import router from './router'
import NProgress from 'nprogress';
import store from './store'
import Message from '@/components/Message/index.js'
import { getToken } from '@/utils/auth'
import path from 'path'

NProgress.configure({
  template: `<div class="bar" role="bar" style="background:#1DE2C3;"><div class="peg"></div></div>
  <div class="spinner" role="spinner"><div class="spinner-icon"></div></div>`,
  showSpinner: false,
});

router.beforeEach(async(to, from, next) => {
    const token = getToken()
    const userinfo = store.getters.userinfo
    NProgress.start();
    if (to.path.indexOf('/base') !== -1) {
      if (token){
        if (userinfo){
          next(store.getters.configure.path)
          NProgress.done()
        }else{
          try {
            await store.dispatch('user/getuserinfo')
            console.log("userinfo:"+store.getters.userinfo);
            await store.dispatch('router/initroutes',store.getters.userinfo.UserRole?store.getters.userinfo.UserRole:0)
            if (!to.matched.length) {
              next(resolvePath(store.getters.currtagroute.path,store.getters.currtagroute.children[0].path))
            }else{
              next({ ...to, replace: true })
            }
          }catch (error) {
            Message.error(error.message)
            await store.dispatch('user/resetToken')
            next(`/base`)
            NProgress.done()
          }
        }
      }else{
        if (!to.matched.length) {
          next(resolvePath(store.getters.currtagroute.path,store.getters.currtagroute.children[0].path))
        }else{
          next()
          NProgress.done()
        }
      }
    }else{
      if (token && userinfo){
        if (!to.matched.length) {
          next(resolvePath(store.getters.currtagroute.path,store.getters.currtagroute.children[0].path))
        }else{
          next()
          NProgress.done()
        }
      }else{
        if (token){
          try {
            await store.dispatch('user/getuserinfo')
            await store.dispatch('router/initroutes',store.getters.userinfo.UserRole?store.getters.userinfo.UserRole:0)
            if (!to.matched.length) {
              next(resolvePath(store.getters.currtagroute.path,store.getters.currtagroute.children[0].path))
            }else{
              next({ ...to, replace: true })
            }
          }catch (error) {
            Message.error(error.message)
            await store.dispatch('user/resetToken')
            next(`/base`)
            NProgress.done()
          }
        }else{
          next(`/base`)
          NProgress.done()
        }
      }
    }
})

router.afterEach(() => {
  NProgress.done();
});

function resolvePath(ppath,vpath) {
  return path.resolve(ppath, vpath)
}