/**
 * 根据文件类型获取图标
 * @returns {*}
 * @param file
 */
function getFileIcon(file) {
    //文件类型
    let type = file.type;

    //文件名后缀
    let suffix = file.name.substring(file.name.lastIndexOf('.') + 1).toLowerCase();

    let icon;

    if (type === 'folder') {
        icon = '/images/folder.png'
    } else if (type.indexOf('image') >= 0) {
        icon = file.thumb;
    } else if (type.indexOf('video') >= 0) {
        icon = '/images/video.png'
    } else if (type.indexOf('audio') >= 0) {
        icon = '/images/audio.png'
    } else if (suffix === 'pdf') {
        icon = '/images/pdf.png'
    } else if (suffix === 'docx' || suffix === 'doc') {
        icon = '/images/docx.png'
    } else if (suffix === 'xlsx' || suffix === 'xls') {
        icon = '/images/excel.png'
    } else if (suffix === 'pptx') {
        icon = '/images/pptx.png'
    } else if (suffix === 'txt') {
        icon = '/images/txt.png'
    } else if (["zip", "rar", "7z", "gzip", "tar"].includes(suffix)) {
        icon = '/images/zip.png'
    } else {
        icon = '/images/file.png'; //默认文件图标
    }
    return icon
}

/**
 * 根据文件大小计算单位
 */
function getFileSize(size) {
    if (!size || size === "0") return '';
    let unit = 'B'
    if (size > 1024) {
        size = size / 1024
        unit = 'KB'
    }
    if (size > 1024) {
        size = size / 1024
        unit = 'MB'
    }
    if (size > 1024) {
        size = size / 1024
        unit = 'GB'
    }
    if (size > 1024) {
        size = size / 1024
        unit = 'TB'
    }
    if (size > 1024) {
        size = size / 1024
        unit = 'PB'
    }
    return size.toFixed(2) + unit
}

/**
 * 获取文件名的简写 将中间的部分替换成...
 */
function getShortFileName(name) {
    if (name.length <= 16) {
        return name;
    }
    let start = name.substring(0, 8)
    let end = name.substring(name.length - 8, name.length)
    return start + '...' + end
}


/**
 * 计算传输速度
 * @param size
 * @param end
 * @param start
 */
function fileTransferSpeed(size, end, start) {
    let speed = 0
    if (end && start) {
        let time = (end - start) / 1000
        speed = size / time
    }
    return getFileSize(speed);
}

export {getFileIcon, getFileSize, getShortFileName, fileTransferSpeed}