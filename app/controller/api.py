@route.route('/api/share', methods=["GET", "POST", "PUT", "DELETE"])
def share():
    """
    share 分享文件
    :return:
    """
    if request.method == "GET":
        hash_id = request.args.get("file")
        res = mysql().table("share").where({"hash_id": hash_id}).find()
        if not res:
            return jsonResponse(3, '分享不存在')

        # 判断分享是否过期
        created_at = res.get("created_at")  # 2023-01-01 00:00:00
        expire = res.get("expire")
        if expire == 1:
            # 当前时间 > 30天 + 创建时间
            if time.time() > (30 * 24 * 60 * 60 + time.mktime(time.strptime(str(created_at), "%Y-%m-%d %H:%M:%S"))):
                return jsonResponse(2, '分享已过期')

        # 校验密码
        if res.get("password") and res.get("password") != request.args.get("p"):
            return jsonResponse(1, '提取码不正确')

        # 分享用户信息
        user = mysql().table("user").where({"id": res.get("uid")}).find()
        result = {}
        result["user"] = user

        file_ids = res.get("fid").split(",")
        files = mysql().table("file").where({"id": ["IN", file_ids, "", "e"]}).select()
        result["share"] = res
        result["files"] = files
        return jsonResponse(0, '成功', result)
    elif request.method == "POST":
        fid = request.args.get("id")
        uid = request.args.get("uid")
        expire = request.args.get("expire")
        type = request.args.get("type")

        # 生成分享hash_id
        hash_id = md5(f'{fid}{uid}{expire}{type}')

        # 判断是否已经分享过
        res = mysql().table("share").where({"hash_id": hash_id}).find()
        if res:
            return jsonResponse(1, '分享已存在')

        # 判断是否需要密码
        if type == "2":
            # 生成4位随机密码 字母+数字
            password = ''.join(random.sample(string.ascii_letters + string.digits, 4))
        else:
            password = ''

        data = {
            "hash_id": hash_id,
            "fid": fid,
            "uid": uid,
            "expire": expire,
            "type": type,
            "password": password,
            "created_at": get_date_time(),
            "updated_at": get_date_time()
        }
        res = mysql().table("share").add(data)
        if res:
            url = f'{request.host_url}#/share?file={hash_id}&p={password}'
            return jsonResponse(0, '分享成功', url)
        return jsonResponse(1, '分享失败')


    elif request.method == "PUT":
        pass
    elif request.method == "DELETE":
        pass


@route.route('/api/file/save', methods=["GET", "POST", "PUT", "DELETE"])
def save():
    """
    保存到自己的文件夹
    :return:
    """
    if request.method == "GET":
        pass
    elif request.method == "POST":
        params = request.get_json()
        fid = params.get("id")
        uid = params.get("uid")
        pid = params.get("pid")

        # 判断是否已经保存过
        res = mysql().table("file").where({"fid": ["IN", fid.split(","), "", "e"], "uid": uid, "pid": pid}).select()

        if res:
            return jsonResponse(1, '文件已存在')

        # 获取文件信息
        res = mysql().table("file").where({"id": ["IN", fid.split(","), "", "e"]}).select()

        user = {}
        user['pid'] = pid
        user['uid'] = uid

        # 递归批量保存文件
        files = multFileSave(user, res, [])

        save_result = mysql().table("file").addAll(files)

        if save_result:
            return jsonResponse(0, '保存成功')
        return jsonResponse(1, '保存失败')

    elif request.method == "PUT":
        pass
    elif request.method == "DELETE":
        pass


def multFileSave(user, data, files=None):
    """
    批量递归保存文件
    :param user:
    :param files:
    :param data:
    :return:
    """
    for item in data:
        pid = item.get("id")
        if item.get("type") == "folder":
            subfile = mysql().table("file").where({"pid": pid}).select()

            # 创建新文件夹
            folder_id = mysql().table("file").add({
                "pid": user['pid'],
                "uid": user['uid'],
                "type": "folder",
                "name": item.get("name"),
                "created_at": get_date_time(),
                "updated_at": get_date_time()
            })

            item["pid"] = folder_id
            multFileSave(item, subfile, files)
        else:
            item["id"] = None
            item["pid"] = user['pid']
            item["uid"] = user['uid']
            item["created_at"] = get_date_time()
            item["updated_at"] = get_date_time()
            files.append(item)
    return files


def multFileRecord(path, files=None):
    """
    批量查询文件记录
    :param files:
    :param data:
    :return:
    """
    for item in path:
        files.append(item)
        if item.get("type") == "folder":
            pid = item.get("id")
            subfile = mysql().table("file").where({"pid": pid}).field("id,path,type,thumb").select()
            multFileRecord(subfile, files)
    return files


@route.route('/api/file', methods=["GET", "POST", "PUT", "DELETE"])
def file():
    """
    file
    :return:
    """
    if request.method == "GET":
        page = request.args.get("page")
        limit = request.args.get("limit")
        keyword = request.args.get("keyword")
        type = request.args.get("type")
        id = request.args.get("id")
        pid = request.args.get("pid", 0)
        uid = request.args.get("uid", 0)
        order_by = request.args.get("order_by", "created_at asc")
        favorite = request.args.get("favorite", 0)  # 是否收藏
        exclude = request.args.get("exclude", "")  # 排除的文件夹id
        delete_flag = request.args.get("delete_flag", 0)

        query = {
            "delete_flag": delete_flag,
            "pid": pid,
            "uid": uid,
        }

        # 是否查询收藏
        if favorite:
            query["favorite"] = favorite

        # 全局搜索移除pid
        if pid == "all":
            query.pop("pid")

        # 搜索条件
        if type or keyword:
            if type:
                query["type"] = ["like", f'%{type}%', '', 'e']
            if keyword:
                query["name"] = ["like", f'%{keyword}%', '', 'e']
            if exclude:
                query["id"] = ["NOT IN", exclude.split(","), "", "e"]
            print(query)
            count = mysql().table('file').where(query).count()
            q = mysql().table('file')
            if page and limit:
                q = q.page(page, limit)
            data = q.where(query).order(order_by).select()
        elif id:
            data = mysql().table('file').where({"id": id}).find()
            return jsonResponse(0, 'success', data)
        elif pid:
            data = mysql().table('file').where(query).order(order_by).select()
            return jsonResponse(0, 'success', data)
        else:
            count = mysql().table('file').count()
            data = mysql().table('file').page(page, limit).order(order_by).select()
        return jsonResponse(0, 'success', data, count)
    elif request.method == "POST":
        data = request.get_json()
        data['created_at'] = get_date_time()
        data['updated_at'] = get_date_time()
        print(data)
        res = mysql().table('file').add(data)
        if res:
            return jsonResponse(0, '操作成功', res)
        else:
            print(mysql().showError())
            return jsonResponse(1, '操作失败')
    elif request.method == "PUT":
        data = request.get_json()
        ids = str(data["id"]).split(",")
        data.pop("id")  # 移除id
        res = mysql().table('file').where({"id": ["IN", ids, "", "e"]}).save(data)
        if res:
            return jsonResponse(0, '操作成功')
        else:
            return jsonResponse(1, '操作失败')
    elif request.method == "DELETE":
        id = request.args.get("id")

        # 是否删除本地文件 1删除 0不删除
        delete_flag = request.args.get("delete_flag", "0")

        if delete_flag == "0":
            res = mysql().table('file').where({"id": ["IN", id.split(","), "", "e"]}).save(
                {"delete_flag": 1, "updated_at": get_date_time()})
            if res:
                return jsonResponse(0, '删除成功')
            return jsonResponse(1, '删除失败')
        elif delete_flag == "1":
            # 批量取出文件路径
            path_list = mysql().table('file').where({"id": ["IN", id.split(","), "", "e"]}).field(
                "id,path,type,thumb").select()

            # 遍历文件夹及子文件夹、文件
            files = multFileRecord(path_list, [])
            files_ids = [str(item["id"]) for item in files]

            # 先从硬盘删除文件
            for file in files:
                if file["type"] == "folder":  # 为目录跳过
                    continue

                # 查询文件是否被其他用户引用
                count = mysql().table('file').where({"path": file["path"], "id": ["NOT IN", files_ids, "", "e"]}).count()

                # 如果被引用则跳过 不删除实体文件
                if count > 0:
                    continue
                try:
                    path = file['path'].replace("/", "\\")  # 处理间隔符
                    filepath = f'{current_app.root_path}/dist/{path}'  # 拼接文件路径
                    os.remove(filepath)  # 删除文件

                    if "image" in file["type"]:
                        thumb_path = file['thumb'].replace("/", "\\")  # 处理间隔符
                        thumb_path = f'{current_app.root_path}/dist/{thumb_path}'
                        os.remove(thumb_path)  # 删除缩略图
                except Exception as e:
                    print(e)
            res = mysql().table('file').where({"id": ["IN", files_ids, "", "e"]}).delete()
            if res:
                return jsonResponse(0, '删除成功')
            return jsonResponse(1, '删除失败')
        else:
            return jsonResponse(1, '请求参数错误')
    return jsonResponse(1, '请求类型错误')
