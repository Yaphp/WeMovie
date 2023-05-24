import request from '@/utils/request'

//上传文件
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

export function uploadChunk(data) {
    return request({
        url: '/upload/chunk',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}