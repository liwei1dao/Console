import LayoutN from '@/layout/LayoutN'

export const Role = {
    Guester : 0,                    //游客
    Operator : 1,                   //运营
    Developer : 2,                  //开发
    Master : 3,                     //超级管理员
}

export const baseRoutes = [
    {
        path: '/base',
        component: LayoutN,
        redirect: 'base/login',
        children: [
          {
              path: 'login',
              component: () => import('@/views/base/login.vue'),
              meta: {title: 'login', icon: 'mdi-beer'},
          },
          {
              path: 'register',
              component: () => import('@/views/base/register.vue'),
              meta: {title: 'register', icon: 'mdi-beer-outline'},
          }
        ]
    }
]