import {createRouter, createWebHashHistory} from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/Login.vue'
import User from '@/views/User.vue'
import Index from '@/views/Index.vue'
const routes = [
    {
        path: '/home',
        name: 'home',
        component: Home,
        children: [
            {
                path: '/index',
                name: 'index',
                component: Index
            },
            {
                path: '/user',
                name: 'user',
                component: User
            }
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/',
        name: 'login',
        component: Login
    },
]

const router = createRouter({
    history: createWebHashHistory(process.env.BASE_URL),
    routes
})

export default router
