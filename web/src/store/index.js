import {createStore} from 'vuex'

export default createStore({
    state: {
        app: {
            "name": "WeMovie",
            "version": "1.0.0",
            "host": window.location.protocol + "://" + window.location.hostname + ":5000",
        },
        user: JSON.parse(sessionStorage.getItem("user")) || null,
        token: sessionStorage.getItem("token") || null,
        tabName: '首页',
        chunkSize: 1024 * 1024 * 10,
        rootMenuData: [],
        userMenuData: [
            {
                entity: {
                    id: 0,
                    name: "/files",
                    icon: "document",
                    alias: "文件"
                }
            },
            {
                entity: {
                    id: 1,
                    name: "/album",
                    icon: "image",
                    alias: "相册"
                }
            },
            {
                entity: {
                    id: 2,
                    name: "/favorite",
                    icon: "heart",
                    alias: "收藏夹"
                }
            },
            {
                entity: {
                    id: 3,
                    name: "/recycle",
                    icon: "trash",
                    alias: "回收站"
                }
            }
        ]
    },
    mutations: {
        SET_TOKEN: (state, token) => {
            state.token = token
            sessionStorage.setItem("token", token)
        },
        SET_USERINFO: (state, user) => {
            state.user = user
            sessionStorage.setItem("user", JSON.stringify(user))
        },
        REMOVE_INFO: (state) => {
            sessionStorage.setItem("token", '')
            sessionStorage.removeItem("user")
            state.user = {}
        }
    },
    getters: {
        getUser: state => {
            return state.user
        },
        getHost: state => {
            return state.app.host
        }
    },
    actions: {},
    modules: {}
})
