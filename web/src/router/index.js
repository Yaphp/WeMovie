import {createRouter, createWebHashHistory} from 'vue-router'
import {isMobileDevice} from "../utils/os.js";

const pcRoutes = [
    {
        path: '/home',
        name: 'home',
        component: () => import('@/views/pc/Home.vue'),
        children: [
            {
                path: '/index',
                name: 'index',
                component: () => import('@/views/pc/Index.vue')
            },
            {
                path: '/profile',
                name: 'Profile',
                component: () => import('@/views/pc/Profile.vue')
            },
            {
                path: '/files',
                name: 'File',
                component: () => import('@/views/pc/File.vue')
            },
            {
                path: '/album',
                name: 'Album',
                component: () => import('@/views/pc/Album.vue')
            },
            {
                path: '/favorite',
                name: 'Favorite',
                component: () => import('@/views/pc/Favorite.vue')
            }, {
                path: '/recycle',
                name: 'Recycle',
                component: () => import('@/views/pc/Recycle.vue')
            },
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/pc/Login.vue')
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/pc/Register.vue')
    },
    {
        path: '/',
        name: 'login',
        component: () => import('@/views/pc/Login.vue')
    },
    {
        path: '/share',
        name: 'Share',
        component: () => import('@/views/pc/Share.vue')
    }

]

const mRoutes = [
    {
        path: '/home',
        name: 'home',
        component: () => import('@/views/m/Home.vue'),
        children: [
            {
                path: '/files',
                name: 'File',
                component: () => import('@/views/m/File.vue')
            },
            {
                path: '/album',
                name: 'Album',
                component: () => import('@/views/m/Album.vue')
            },
            {
                path: '/user',
                name: 'user',
                component: () => import('@/views/m/User.vue')
            },
            {
                path: '/profile',
                name: 'Profile',
                component: () => import('@/views/m/Profile.vue')
            },
            {
                path: '/favorite',
                name: 'Favorite',
                component: () => import('@/views/m/Favorite.vue')
            },
            {
                path: '/recycle',
                name: 'Recycle',
                component: () => import('@/views/m/Recycle.vue')
            }
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/m/Login.vue')
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/m/Register.vue')
    },
    {
        path: '/search',
        name: 'Search',
        component: () => import('@/views/m/Search.vue')
    },
    {
        path: '/result',
        name: 'Result',
        component: () => import('@/views/m/Result.vue')
    },
    {
        path: '/share',
        name: 'Share',
        component: () => import('@/views/m/Share.vue')
    },
    {
        path: '/',
        name: 'login',
        component: () => import('@/views/m/Login.vue')
    }
]

let routes;
if (isMobileDevice()) {
    routes = mRoutes
} else {
    routes = pcRoutes
}

// console.log(routes)

const router = createRouter({
    history: createWebHashHistory(process.env.BASE_URL),
    routes: routes
})

export default router
