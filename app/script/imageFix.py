import os
from standard.common.helper import getImageThumbnail

# 读取 dist/uploads/20230514下的所有图片
dir = '../../dist/uploads/20230514'

# 获取目录下的图片文件
files = os.listdir(dir)

# 遍历所有文件
print(files)

for i in files:
    # 获取文件名
    name = i.split('.')[0]

    # 获取文件后缀
    suffix = i.split('.')[1]

    if suffix not in ['gif']:
        continue

    if "thumbnail_" in i:
        # 删除缩略图
        os.remove(dir + '/' + i)

    if suffix in ['jpg', 'png', 'jpeg', 'webp', 'gif'] and "thumbnail_" not in i:
        print(i)
        # 生成缩略图
        getImageThumbnail(dir + '/' + i)
