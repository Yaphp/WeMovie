import {createStore} from 'vuex'

export default createStore({
    state: {
        app: {
            "name": "WeMovie",
            "version": "1.0.0",
            "host": "http://" + window.location.hostname + ":80",
        },
        user: JSON.parse(sessionStorage.getItem("user")) || {},
        token: sessionStorage.getItem("token") || "",
        tabName: '首页',
        adminMenuData: [
            {
                entity: {
                    id: 0,
                    name: "/index",
                    icon: "el-icon-s-home",
                    alias: "首页"
                }
            },
            {
                entity: {
                    id: 1,
                    name: "/user",
                    icon: "el-icon-s-custom",
                    alias: "用户管理"
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
