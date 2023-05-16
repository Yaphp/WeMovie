import request from '@/utils/request'

export function shareAdd(data) {
    return request({
        url: '/share',
        method: 'post',
        params: data
    })
}

export function shareDelete(id) {
    return request({
        url: '/share',
        method: 'delete',
        params: {
            id
        }
    })
}

export function shareUpdate(data) {
    return request({
        url: '/share',
        method: 'put',
        data
    })
}

export function shareList(data) {
    return request({
        url: '/share',
        method: 'get',
        params: data
    })
}