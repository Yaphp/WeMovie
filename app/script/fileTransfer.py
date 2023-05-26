from standard.common.helper import *
from threading import Thread  # 导入多线程


def fileTransfer(categorys, pid=0):
    for category in categorys:
        # 多线程创建文件夹
        Thread(target=createCategory, args=(category, pid)).start()


def createCategory(category, pid=0):
    # 创建文件夹
    folder_id = createFolder(category['title'], pid)

    if folder_id:
        print(f"{category['title']} 创建文件夹成功")

    # 查询此类别有无记录
    gallerys = mysql().table("we_gallery").where({"category": category['id']}).select()

    # 遍历所有图库
    for gallery in gallerys:
        # 一个gallery一个文件夹
        gallery_folder_id = createFolder(gallery['title'], folder_id)

        # 查询所有章节
        chapters = mysql().table("we_gallery_chapter").where({"gallery_id": gallery['id']}).select()

        # 遍历所有章节
        for chapter in chapters:
            # 查询所有图片
            sections = mysql().table("we_gallery_section").where({"chapter_id": chapter['id']}).select()
            for section in sections:
                # 保存图片
                data = {
                    'pid': gallery_folder_id,
                    'name': section['sort_id'],
                    'url': section['url']
                }
                file_id = createFile(data)

                if file_id:
                    print(f"{category['title']} - 第 {section['sort_id']} 张图片 保存成功")

    # 子类别
    subCategorys = mysql().table("we_gallery_category").where({"pid": category['id']}).select()

    fileTransfer(subCategorys, folder_id)


def createFolder(name, pid=0):
    folder = {
        'pid': pid,
        'name': name,
        'uid': 1,
        'thumb': '',
        'size': 0,
        'type': 'folder',
        'path': '',
        'favorite': 0,
        'delete_flag': 0,
        'created_at': get_date_time(),
        'updated_at': get_date_time()
    }

    # 添加文件夹 返回id
    return mysql().table("we_file").add(folder)


def createFile(data):
    # 获取文件后缀
    suffix = data['url'].split('.')[-1]
    file = {
        'pid': data['pid'],
        'name': data['name'],
        'uid': 1,
        'thumb': data['url'],
        'size': 0,
        'type': 'images/' + suffix,
        'path': data['url'],
        'favorite': 0,
        'delete_flag': 0,
        'created_at': get_date_time(),
        'updated_at': get_date_time()
    }
    return mysql().table("we_file").add(file)


if __name__ == '__main__':
    categorys = mysql().table("we_gallery_category").where({"pid": 0}).select()

    fileTransfer(categorys, 0)
