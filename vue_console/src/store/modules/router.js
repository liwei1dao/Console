import { constantRoutes } from '@/router'
/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes,role) {
    const result = {}
    routes.forEach(route => {
        const tmp = { ...route}
        if (tmp.navigation && role >= tmp.role) {
            if (!Object.prototype.hasOwnProperty.call(result,tmp.navigation.name)){
                var temp = {
                    name : tmp.navigation.name,
                    title: tmp.navigation.title,
                    icon: tmp.navigation.icon,
                    default:0,
                    routes:[]
                }
                result[tmp.navigation.name] = temp
            }
            result[tmp.navigation.name].routes.push(tmp)
            if (tmp.default){
                result[tmp.navigation.name].default = result[tmp.navigation.name].routes.length-1
            }
        }
    })
    return result
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function defRoutes(routes) {
    for (var key in routes) {
        var value = routes[key];
        for(var i in value.routes){
            if (value.routes[i].default){
                return value.routes[i]
            }
        }
    }
    return null
}



const state = {
    routes : null,
    currtagroute : null,
}

const mutations = {
    Set_routes: (state, routes) => {
        state.routes = routes
    },
    Set_currtagroute: (state, currtagroute) => {
        state.currtagroute = currtagroute
    },
}

const actions = {
    initroutes({commit},role) {
        return new Promise((resolve, reject) => {
            var routes = filterAsyncRoutes(constantRoutes,role)
            commit("Set_routes",routes)
            var defroutes = defRoutes(routes)
            if (defroutes == null){
                reject(new Error('没有找到可供显示的网关组'))
                return
            }
            commit("Set_currtagroute",defroutes)
            resolve()
        })
    },
    setcurrtagroute({commit},currtagroute){
        commit("Set_currtagroute",currtagroute)
    }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}