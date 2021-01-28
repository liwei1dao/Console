import Vue from 'vue'
import Router from 'vue-router'

import {baseRoutes,consoleRoutes,devopsRoutes,demoRoutes} from './routers'
Vue.use(Router)

// export const constantRoutes = [
//     {
//         path: '/base',
//         component: LayoutN,
//         redirect: 'base/login',
//         children: [
//           {
//               path: 'login',
//               component: () => import('@/views/base/login.vue'),
//               meta: {title: 'login', icon: 'mdi-beer'},
//           },
//           {
//               path: 'register',
//               component: () => import('@/views/base/register.vue'),
//               meta: {title: 'register', icon: 'mdi-beer-outline'},
//           }
//         ]
//     },
//     {
//         path: '/console',
//         component: Layout,
//         navigation:{
//             name:"console",
//             title: 'console',
//             icon: 'mdi-home-floor-b',
//         },
//         role:Role.Master,
//         default:true,
//         children: [
//             {
//                 path: 'project',
//                 component: () => import('@/views/console/project.vue'),
//                 meta: {title: 'project', icon: 'mdi-home-floor-b'},
//             }
//         ]
//     },
//     {
//         path: '/console',
//         component: Layout,
//         navigation:{
//             name:"console",
//             title: 'console',
//             icon: 'mdi-home-floor-b',
//         },
//         role:Role.Master,
//         children: [
//             {
//                 path: 'host',
//                 component: () => import('@/views/console/host.vue'),
//                 meta: {title: 'host', icon: 'mdi-home-floor-b'},
//             }
//         ]
//     },
//     {
//         path: '/console',
//         component: Layout,
//         navigation:{
//             name:"console",
//             title: 'console',
//             icon: 'mdi-home-floor-b',
//         },
//         role:Role.Master,
//         children: [
//             {
//                 path: 'cluster',
//                 component: () => import('@/views/console/cluster.vue'),
//                 meta: {title: 'cluster', icon: 'mdi-details'},
//             }
//         ]
//     },
//     {
//         path: '/devops',
//         component: Layout,
//         navigation:{
//             name:"devops",
//             title: 'devops',
//             icon: 'mdi-wall',
//         },
//         role:Role.Master,
//         default:true,
//         children: [
//             {
//                 path: 'index',
//                 component: () => import('@/views/devops/index.vue'),
//                 meta: {title: 'devops', icon: 'mdi-wall'},
//             }
//         ]
//     }
// ]

export const constantRoutes = [].concat(baseRoutes,consoleRoutes,devopsRoutes,demoRoutes);

const createRouter = () => new Router({
// mode: 'history', // require service support
scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
})

const router = createRouter()
// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
}

export default router