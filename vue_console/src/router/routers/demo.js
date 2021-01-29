import {Role} from './base'
import Layout from '@/layout/Layout'
/*
演示网关路由
*/
export const demoRoutes = [
    {
        path: '/demo',
        component: Layout,
        navigation:{
            name:"demo",
            title:'demo',
            icon: 'mdi-wall',
        },
        role:Role.Guester,
        default:true,
        children: [
            {
                path: 'index',
                component: () => import('@/views/demo/index.vue'),
                meta: {title: 'demo', icon: 'mdi-wall'},
            }
        ]
    },
    {
        path: '/demo/components',
        component: Layout,
        navigation:{
            name:"demo",
            title:'demo',
            icon: 'mdi-wall',
        },
        meta: {title: 'component', icon: 'mdi-wall'},
        role:Role.Guester,
        children: [
            {
                path: 'divider',
                component: () => import('@/views/demo/components/divider.vue'),
                meta: {title: 'divider', icon: 'mdi-wall'},
            },
            {
                path: 'message',
                component: () => import('@/views/demo/components/message.vue'),
                meta: {title: 'message', icon: 'mdi-wall'},
            }
        ]
    }
]