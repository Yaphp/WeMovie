import request from '@/utils/request'

//添加记录
export function upload(data) {
    return request({
        url: '/upload',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}