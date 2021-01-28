import {Role} from './base'
import Layout from '@/layout/Layout'
/*
控制台网关路由
*/
export const consoleRoutes = [
    {
        path: '/console',
        component: Layout,
        navigation:{
            name:"console",
            title: 'console',
            icon: 'mdi-home-floor-b',
        },
        role:Role.Master,
        default:true,
        children: [
            {
                path: 'project',
                component: () => import('@/views/console/project.vue'),
                meta: {title: 'project', icon: 'mdi-home-floor-b'},
            }
        ]
    },
    {
        path: '/console',
        component: Layout,
        navigation:{
            name:"console",
            title: 'console',
            icon: 'mdi-home-floor-b',
        },
        role:Role.Master,
        children: [
            {
                path: 'host',
                component: () => import('@/views/console/host.vue'),
                meta: {title: 'host', icon: 'mdi-home-floor-b'},
            }
        ]
    },
    {
        path: '/console',
        component: Layout,
        navigation:{
            name:"console",
            title: 'console',
            icon: 'mdi-home-floor-b',
        },
        role:Role.Master,
        children: [
            {
                path: 'cluster',
                component: () => import('@/views/console/cluster.vue'),
                meta: {title: 'cluster', icon: 'mdi-details'},
            }
        ]
    },
]