import request from '@/utils/request'

//用户列表
export function userList(data) {
    return request({
        url: '/user',
        method: 'get',
        params: data
    })
}

//添加用户
export function userAdd(data) {
    return request({
        url: '/user',
        method: 'post',
        data
    })
}

//修改用户
export function userUpdate(data) {
    return request({
        url: '/user',
        method: 'put',
        data
    })
}

//删除用户
export function userDelete(id) {
    return request({
        url: '/user',
        method: 'delete',
        params: {
            id
        }
    })
}