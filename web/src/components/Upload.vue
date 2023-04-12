<template>
    <!--添加图片的对话框-->
    <el-dialog
        title="上传图片"
        v-model="addDialogVisible"
        width="30%"
        @close="addDialogClosed"
    >
        <!--内容主体区域-->
        <el-form :model="addForm" label-width="150px" ref="addFormRef">
            <el-form-item label="车牌号" prop="carplate">
                <el-input v-model="addForm.carplate"></el-input>
            </el-form-item>
            <el-form-item label="图片" prop="picture">
                <el-upload
                    class="upload-demo"
                    drag
                    :http-request="handleUpload"
                    :on-preview="handlePreview"
                >
                    <el-icon class="el-icon--upload">
                        <upload-filled/>
                    </el-icon>
                    <div class="el-upload__text">
                        拖拽文件到此处 <em>或点击上传</em>
                    </div>
                    <!--                    <template #tip>-->
                    <!--                        <div class="el-upload__tip">-->
                    <!--                            jpg/png files with a size less than 500kb-->
                    <!--                        </div>-->
                    <!--                    </template>-->
                </el-upload>
            </el-form-item>
            <el-form-item label="上传时间" prop="upload_at">
                <el-date-picker
                    v-model="addForm.upload_at"
                    type="datetime"
                    placeholder="请选择上传时间"
                    format="YYYY-MM-DD HH:mm:ss"
                    value-format="YYYY-MM-DD HH:mm:ss"
                    :shortcuts="shortcuts"
                />
            </el-form-item>
            <el-form-item label="拍照时间" prop="photo_at">
                <el-date-picker
                    v-model="addForm.photo_at"
                    type="datetime"
                    placeholder="请选择拍照时间"
                    format="YYYY-MM-DD HH:mm:ss"
                    value-format="YYYY-MM-DD HH:mm:ss"
                    :shortcuts="shortcuts"
                />
            </el-form-item>
        </el-form>
        <!--底部按钮区域-->
        <span slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addState">确 定</el-button>
            </span>
    </el-dialog>


    <!--修改图片的对话框-->
    <el-dialog title="修改图片" v-model="editDialogVisible" width="30%">
        <!--内容主体区域-->
        <el-form :model="editForm" label-width="70px">
            <el-form-item label="车牌号" prop="carplate">
                <el-input v-model="editForm.carplate"></el-input>
            </el-form-item>
            <el-form-item label="行为图片" prop="picture">
                <el-upload
                    class="upload-demo"
                    drag
                    :http-request="handleEditUpload"
                    :on-preview="handleEditPreview"
                >
                    <el-icon v-if="!editForm.picture" class="el-icon--upload">
                        <upload-filled/>
                    </el-icon>

                    <el-image v-if="editForm.picture" style="width: 100%; height: auto" :src="this.$store.state.app.host+editForm.picture" :fit="contain"/>

                    <div class="el-upload__text">
                        拖拽文件到此处 <em>或点击上传</em>
                    </div>
                    <!--                    <template #tip>-->
                    <!--                        <div class="el-upload__tip">-->
                    <!--                            jpg/png files with a size less than 500kb-->
                    <!--                        </div>-->
                    <!--                    </template>-->
                </el-upload>
            </el-form-item>
            <el-form-item label="拍照时间" prop="photo_at">
                <el-date-picker
                    v-model="editForm.photo_at"
                    type="datetime"
                    placeholder="请选择拍照时间"
                    format="YYYY-MM-DD HH:mm:ss"
                    value-format="YYYY-MM-DD HH:mm:ss"
                    :shortcuts="shortcuts"
                />
            </el-form-item>
            <el-form-item label="上传时间" prop="upload_at">
                <el-date-picker
                    v-model="editForm.upload_at"
                    type="datetime"
                    placeholder="请选择拍照时间"
                    format="YYYY-MM-DD HH:mm:ss"
                    value-format="YYYY-MM-DD HH:mm:ss"
                    :shortcuts="shortcuts"
                />
            </el-form-item>
        </el-form>
        <!--底部按钮区域-->
        <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="editState">确 定</el-button>
        </span>
    </el-dialog>
</template>

<script>
import {Delete, Edit, UploadFilled} from "@element-plus/icons-vue";
import {upload} from "../api/upload.js";

export default {
    setup() {
        return {
            Edit,
            Delete
        }
    },
    components: {
        UploadFilled
    },
    data() {
        return {
            imageList: [], // 数据列表
            total: 0, // 数据总数
            options: [
                {key: "车牌", value: 'carplate', label: '车牌'},
                // {key: "行为图片", value: 'picture', label: '行为图片'},
                // {key: "行为模式", value: 'mode', label: '行为模式'},
                // {key: "速度", value: 'speed', label: '速度'},
                // {key: "备注", value: 'remark', label: '备注'},
                {key: "上传时间", value: 'upload_at', label: '上传时间'},
                {key: "拍摄时间", value: 'photo_at', label: '拍摄时间'},
            ],
            shortcuts: [
                {
                    text: '今天',
                    value: new Date(),
                },
                {
                    text: '昨天',
                    value: () => {
                        const date = new Date()
                        date.setTime(date.getTime() - 3600 * 1000 * 24)
                        return date
                    },
                },
                {
                    text: '一周前',
                    value: () => {
                        const date = new Date()
                        date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
                        return date
                    },
                },
            ],
            // 获取用户列表的参数对象
            queryInfo: {
                field: "",//查询字段
                keyword: "", // 查询参数
                pageNo: 1, // 当前页码
                pageSize: 10, // 每页显示条数
            },
            addDialogVisible: false, // 控制添加用户对话框是否显示
            addForm: {
                uid: this.$store.state.user.id,
                carplate: "",
                picture: "",
                upload_at: "",
                photo_at: "",
            },
            editDialogVisible: false, // 控制修改用户信息对话框是否显示
            editForm: {
                carplate: "",
                picture: "",
                upload_at: "",
                photo_at: "",
            },
            multipleSelection: [],
            ids: [],
            previewImage: false,
            previewImageUrl: "",
        };
    },
    created() {
        if(this.$store.state.user.role !==0){
            this.$notify({
                title: '错误',
                message: '您没有权限访问该页面',
                type: 'error',
                duration: 2000
            });
            this.$store.state.tagName = "用户管理";
            this.$router.push({path: '/user'})
        }else{
            this.getStateList();
        }
    },
    methods: {
        getStateList() {
            imageList(this.queryInfo)
                .then((res) => {
                    if (res.code === 200) {
                        //列表
                        this.imageList = res.data.list;
                        this.total = res.data.total || 0;
                    } else {
                        this.$message.error(res.msg);
                    }
                })
                .catch((err) => {
                    console.log(err);
                });
        },

        // 监听 pageSize 改变的事件
        handleSizeChange(newSize) {
            // console.log(newSize)
            this.queryInfo.pageSize = newSize;
            // 重新发起请求用户列表
            this.getStateList();
        },

        // 监听 当前页码值 改变的事件
        handleCurrentChange(newPage) {
            // console.log(newPage)
            this.queryInfo.pageNo = newPage;
            // 重新发起请求用户列表
            this.getStateList();
        },

        // 监听 添加对话框的关闭事件
        addDialogClosed() {
            // 表单内容重置为空
            this.$refs.addFormRef.resetFields();
        },

        // 监听 修改状态
        showEditDialog(info) {
            this.editDialogVisible = true;
            this.editForm = info;
        },

        //添加数据
        addState() {
            imageAdd(this.addForm)
                .then((res) => {
                    console.log(res)
                    if (res.code === 200) {
                        this.addDialogVisible = false;
                        this.getStateList();
                        this.$message({
                            message: "添加图片成功",
                            type: "success",
                        });
                    } else {
                        this.$message.error("添加图片失败");
                    }
                })
                .catch((err) => {
                    this.$message.error("添加图片异常");
                    console.log(err);
                });
        },

        //修改用户
        editState() {
            imageUpdate(this.editForm)
                .then((res) => {
                    if (res.code === 200) {
                        this.editDialogVisible = false;
                        this.getStateList();
                        this.$message({
                            message: "修改成功",
                            type: "success",
                        });
                    } else {
                        this.$message.error("修改失败");
                    }
                })
                .catch((err) => {
                    this.$message.error("修改异常");
                    console.log(err);
                });
        },

        // 根据ID删除对应的数据信息
        async removeStateById(id) {
            // 弹框 询问用户是否删除
            const confirmResult = await this.$confirm(
                "此操作将永久删除该数据, 是否继续?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            // console.log(confirmResult)
            if (confirmResult === "confirm") {
                //删除数据
                imageDelete(id)
                    .then((res) => {
                        if (res.code === 200) {
                            this.getStateList();
                            this.$message({
                                message: "删除成功",
                                type: "success",
                            });
                        } else {
                            this.$message.error("删除失败");
                        }
                    })
                    .catch((err) => {
                        this.$message.error("删除异常");
                        console.log(err);
                    });
            }
        },

        //批量选中事件处理
        handleSelectionChange(val) {
            console.log(val)
            this.multipleSelection = val;
        },

        //批量删除用户
        async batchDeleteState() {
            // 弹框 询问用户是否删除
            const confirmResult = await this.$confirm(
                "此操作将永久删除数据, 是否继续?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            if (confirmResult === "confirm") {
                //向被删除的ids赋值
                this.multipleSelection.forEach((item) => {
                    this.ids.push(item.id);
                });
                imageDelete(this.ids.join(","))
                    .then((res) => {
                        if (res.code === 200) {
                            this.getStateList();
                            this.$message({
                                message: "批量删除成功",
                                type: "success",
                            });
                        } else {
                            this.$message.error("批量删除失败");
                        }
                    })
                    .catch((err) => {
                        this.$message.error("批量删除异常");
                        console.log(err);
                    });
            }
        },

        handlePreview(file) {
            console.log(file);
            window.open(this.$store.state.app.host + this.addForm.picture)
        },

        handleUpload(file) {
            const formData = new FormData()
            formData.append('file', file.file) // 传入bpmn文件

            upload(formData).then((res) => {
                if (res.code === 200) {
                    this.$message({
                        message: "上传成功",
                        type: "success",
                    });

                    // 上传成功后，将上传返回值给表单
                    this.addForm.picture = res.data
                } else {
                    this.$message.error("上传失败");
                }
            })
                .catch((err) => {
                    this.$message.error("上传异常");
                    console.log(err);
                });
        },

        handleEditPreview(file) {
            console.log(file);
            window.open(this.$store.state.app.host + this.editForm.picture)
        },

        handleEditUpload(file) {
            console.log(file)
            const formData = new FormData()
            formData.append('file', file.file) // 传入bpmn文件

            upload(formData).then((res) => {
                if (res.code === 200) {
                    this.$message({
                        message: "上传成功",
                        type: "success",
                    });

                    // 上传成功后，将上传返回值给表单
                    this.editForm.picture = res.data
                } else {
                    this.$message.error("上传失败");
                }
            })
                .catch((err) => {
                    this.$message.error("上传异常");
                    console.log(err);
                });
        },
    },
};
</script>

<style>
.table-container {
    margin: 40px 20px;
}

.el-row {
    margin-bottom: 20px;
}

.el-col {
    border-radius: 4px;
}

.el-card {
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1) !important;
    height: 60pt;
}
</style>
