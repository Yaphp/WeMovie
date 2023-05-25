import request from '@/utils/request'

//文件列表
export function fileList(data) {
    return request({
        url: '/file',
        method: 'get',
        params: data
    })
}

//添加文件
export function fileAdd(data) {
    return request({
        url: '/file',
        method: 'post',
        data
    })
}

//修改文件
export function fileUpdate(data) {
    return request({
        url: '/file',
        method: 'put',
        data
    })
}

//删除文件
export function fileDelete(id, delete_flag = 0) {
    return request({
        url: '/file',
        method: 'delete',
        params: {
            id: id,
            delete_flag: delete_flag
        }
    })
}

// 恢复文件
export function fileRecover(ids) {
    return request({
        url: '/file',
        method: 'put',
        data: {
            'id': ids,
            'delete_flag': 0
        }
    })
}

export function fileFavorite(data) {
    return request({
        url: '/file/favorite',
        method: 'put',
        params: data
    })
}

export function fileSave(data) {
    return request({
        url: '/share/save',
        method: 'post',
        data
    })
}