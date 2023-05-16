// 导入axios实例
import request from '@/utils/request'

// 用户登录
export function userLogin(data) {
    return request({
        url: '/login',
        method: 'post',
        data: data,
    })
}

// 用户注册
export function userRegister(data) {
    return request({
        url: '/register',
        method: 'post',
        data: data,
    })
}