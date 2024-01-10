package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	"strconv"
	"strings"
	"weapp/app/model"
	"weapp/app/utils"
)

type FileController struct {
	BaseController
}

// Add 添加文件
func (con FileController) Add(c *gin.Context) {
	file := model.File{}

	err := c.ShouldBindJSON(&file)

	if err != nil {
		con.Error(c, "参数解析失败")
		return
	}

	file.CreatedAt = model.LocalTime(utils.GetDateTime())
	file.UpdatedAt = model.LocalTime(utils.GetDateTime())

	//fmt.Println(file)

	result := model.Db.Create(&file)

	if result.Error != nil {
		con.Error(c, result.Error.Error())
	} else {
		con.Success(c, file.Id)
	}
}

// Delete 删除文件
func (con FileController) Delete(c *gin.Context) {
	id := c.Query("id")
	ids := strings.Split(id, ",")

	// 是否删除本地文件 1删除 0不删除
	deleteFlag := c.DefaultQuery("delete_flag", "0")

	if deleteFlag == "0" {
		// 更新删除状态
		err := model.Db.Where("id in (?)", ids).Delete(&model.File{}).Error
		if err != nil {
			con.Error(c, "删除失败")
			return
		} else {
			con.Success(c, "删除成功")
			return
		}
	} else {
		// 批量取出文件路径
		var files []model.File

		err := model.Db.Where("id in (?)", ids).Find(&files).Error

		if err != nil {
			con.Error(c, "读取路径数据失败")
			return
		}

		// 遍历文件夹及子文件夹、文件
		var dirs []model.File
		dirs = model.GetDirAndFile(files, dirs)

		// 取出所有的file的id作为数组
		var fileIds []int
		for _, v := range dirs {
			fileIds = append(fileIds, v.Id)
		}

		var path string
		var thumbnailPath string

		// 遍历文件
		for _, v := range dirs {
			fmt.Println(v)
			if v.Type == "folder" {
				// 跳过
				continue
			}

			// 当前文件被引用次数
			var total int64

			// 查询文件是否被其他非当前目录的用户引用
			model.Db.Find(&model.File{}).Where("id not in ?", fileIds).Where("path = ?", v.Path).Count(&total)

			// 如果被引用次数大于0，跳过
			if total > 0 {
				continue
			}

			// 如果当前系统为windows，将路径中的/替换为\ // 拼接真实路径
			if runtime.GOOS == "windows" {
				path = utils.GetRootPath() + strings.Replace(v.Path, "/", "\\", -1)
			} else {
				path = utils.GetRootPath() + v.Path
			}

			// 删除文件
			err := os.Remove(path)
			if err != nil {
				fmt.Println("删除文件失败", err)
			}

			// 如果是图片，删除缩略图
			if strings.Contains(v.Thumb, ".jpg") || strings.Contains(v.Thumb, "png") || strings.Contains(v.Thumb, "jpeg") {

				if runtime.GOOS == "windows" {
					// 拼接缩略图路径
					thumbnailPath = utils.GetRootPath() + strings.Replace(v.Thumb, "/", "\\", -1)
				} else {
					thumbnailPath = utils.GetRootPath() + v.Thumb
				}

				// 删除缩略图
				err := os.Remove(thumbnailPath)
				if err != nil {
					fmt.Println("删除缩略图失败", err)
				}
			}
		}

		var file model.File

		//把ids遍历转为数字然后加入fileIds
		for _, v := range ids {
			idsInt, _ := strconv.Atoi(v)
			fileIds = append(fileIds, idsInt)
		}

		err = model.Db.Unscoped().Where("id in (?)", fileIds).Delete(&file).Error

		if err != nil {
			con.Error(c, "删除失败")
		} else {
			con.Success(c, "删除成功")
		}
	}
}

// Update 更新文件
func (con FileController) Update(c *gin.Context) {
	file := model.File{}

	var params = make(map[string]interface{})

	json.NewDecoder(c.Request.Body).Decode(&params)

	var ids interface{}

	// 如果params['id']是字符串
	if _, ok := params["id"].(string); ok {
		ids = strings.Split(params["id"].(string), ",")
	} else {
		ids = params["id"]
	}

	// 删除id
	delete(params, "id")

	result := model.Db.Debug().Model(&file).Where("id in (?)", ids).Updates(params)

	// 判断是否修改成功
	if result.RowsAffected > 0 {
		con.Success(c, "更新成功")
	} else {
		con.Error(c, result.Error.Error())
	}
}

// Index 获取文件
func (con FileController) Index(c *gin.Context) {
	page := c.Query("page")                                 // 页码
	limit := c.Query("limit")                               // 每页条数
	keyword := c.DefaultQuery("keyword", "")                // 关键词
	fileType := c.DefaultQuery("type", "")                  // 文件类型
	id := c.DefaultQuery("id", "")                          // 文件id
	pid := c.DefaultQuery("pid", "0")                       // 父级id
	uid := c.DefaultQuery("uid", "0")                       // 用户id
	favorite := c.DefaultQuery("favorite", "")              // 是否收藏
	exclude := c.DefaultQuery("exclude", "")                // 排除的文件夹id
	deleteFlag := c.DefaultQuery("delete_flag", "0")        // 删除标记
	orderBy := c.DefaultQuery("order_by", "created_at asc") // 排序

	var pageNo int
	var pageSize int

	if page != "" && limit != "" {
		pageNo = utils.StrToInt(page)    // 把params["page"]转换成int
		pageSize = utils.StrToInt(limit) // 把params["limit"]转换成int

		// 默认分页
		if pageNo <= 0 {
			pageNo = 1
		}

		// 默认每页10条
		if pageSize <= 0 {
			pageSize = 10
		}
	} else {
		pageNo = -1
		pageSize = -1
	}

	// 声明查询条件
	query := map[string]interface{}{
		"pid": pid,
		"uid": uid,
	}

	// 非搜索全部时指定pid
	if pid == "all" {
		delete(query, "pid")
	}

	// 是否收藏查询条件
	if favorite != "" {
		query["favorite"] = favorite
	}

	// 文件对象
	var file []model.File

	// 总条数
	var total int64

	// 创建文件模型
	fileModel := model.Db.Debug().Where(query)

	// 是否查询已删除的数据
	if deleteFlag == "1" {
		fileModel = fileModel.Unscoped()
	}

	// 搜索条件
	if fileType != "" || keyword != "" {
		// 按文件类型搜索
		if fileType != "" {
			fileModel.Where("type LIKE ?", "%"+fileType+"%")
		}

		// 按文件名搜索
		if keyword != "" {
			fileModel.Where("name LIKE ?", "%"+keyword+"%")
		}

		// 排除的文件夹
		if exclude != "" {
			ids := strings.Split(exclude, ",")
			fileModel.Where("id NOT IN (?)", ids)
		}

		// 计算总条数
		fileModel.Find(&file).Count(&total)

		// 判断是否分页
		if pageNo > 0 && pageSize > 0 {
			fileModel.Offset((pageNo - 1) * pageSize).Limit(pageSize)
		}
		// 查询排序
		fileModel.Order(orderBy).Find(&file)
	} else if id != "" {
		file := model.File{}      // 单文件重新声明
		fileModel.Find(&file, id) // 查询单文件
	} else if pid != "" {
		// 根据pid查询
		fileModel.Where(query).Order(orderBy).Find(&file).Count(&total)
		if pageNo > 0 && pageSize > 0 {
			fileModel.Offset((pageNo - 1) * pageSize).Limit(pageSize)
		}
		fileModel.Find(&file)
	} else {
		fileModel.Count(&total) // 计算总条数
		if pageNo > 0 && pageSize > 0 {
			fileModel.Offset((pageNo - 1) * pageSize).Limit(pageSize)
		}
		fileModel.Order(orderBy).Find(&file)
	}

	con.SuccessWithCount(c, file, total)
}

// Favorite 收藏/取消收藏文件
func (con FileController) Favorite(c *gin.Context) {
	id := c.Query("id")           // 文件id
	ids := strings.Split(id, ",") // 文件ids
	status := c.Query("status")   // 收藏状态

	data := make(map[string]interface{})

	data["favorite"] = status
	data["updated_at"] = model.LocalTime(utils.GetDateTime())

	err := model.Db.Model(&model.File{}).Where("id in (?)", ids).Updates(data).Error

	if err != nil {
		con.Error(c, "收藏失败")
	} else {
		con.Success(c, "收藏成功")
	}
}
