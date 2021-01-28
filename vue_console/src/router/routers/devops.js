import {Role} from './base'
import Layout from '@/layout/Layout'
/*
运营网关路由
*/
export const devopsRoutes = [
    {
        path: '/devops',
        component: Layout,
        navigation:{
            name:"devops",
            title: 'devops',
            icon: 'mdi-wall',
        },
        role:Role.Master,
        default:true,
        children: [
            {
                path: 'index',
                component: () => import('@/views/devops/index.vue'),
                meta: {title: 'devops', icon: 'mdi-wall'},
            }
        ]
    }
]