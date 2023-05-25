<template>
    <el-dialog
        v-model="showDiag"
        title="移动文件"
        width="450px"
    >
        <div class="file-dir">
            <div class="file-dir-item" v-for="(bread,index) in breadcrumb" @click="changeDirectory(bread.id,index)">
                {{ bread.name }}
            </div>
        </div>
        <div class="file-move-list">
            <div class="file-move-item" v-if="addFolderVisible">
                <div class="file-left">
                    <img class="file-icon" :src="getFileIcon(newFile)">
                    <div class="file-name">
                        <el-input v-model="newFile.name" style="width:220px"></el-input>
                        <el-icon size="22" class="op-icon" color="#8999f9" @click="handleCreateFolder">
                            <SuccessFilled/>
                        </el-icon>
                        <el-icon size="22" class="op-icon" color="gray" @click="addFolderVisible = false">
                            <CircleCloseFilled/>
                        </el-icon>
                    </div>
                </div>
            </div>
            <div class="file-move-item" v-for="file in fileList" @click="handleClickRow(file)">
                <div class="file-left">
                    <img class="file-icon" :src="getFileIcon(file)" :alt="file.name">
                    <div class="file-name">
                        {{ getShortFileName(file.name) }}
                    </div>
                </div>
            </div>
        </div>
        <template #footer>
            <div class="dialog-footer">
                <div class="left">
                    <el-button type="primary" plain @click="addFolderVisible = true">新建文件夹</el-button>
                </div>
                <div class="right">
                    <el-button @click="handleClose">取消</el-button>
                    <el-button type="primary" @click="handleConfirm">移动到此处</el-button>
                </div>
            </div>
        </template>
    </el-dialog>
</template>

<script>
import {fileList, fileUpdate, fileAdd} from "api/file";
import {getFileIcon, getShortFileName} from "utils/file";
import {SuccessFilled, CircleCloseFilled} from "@element-plus/icons-vue";

export default {
    name: "FileMove",
    props: {
        show: Boolean,
        file: [Object, Array]
    },
    components: {
        SuccessFilled,
        CircleCloseFilled
    },
    setup() {
        return {
            getFileIcon,
            getShortFileName
        }
    },
    data() {
        return {
            showDiag: this.show,
            breadcrumb: [{id: 0, name: "全部文件"}], // 面包屑导航
            pid: 0, //移动后的父级目录
            files: this.file instanceof Array ? this.file : [this.file],
            fileList: [],
            newFile: {
                name: "新建文件夹",
                type: "folder"
            },
            addFolderVisible: false
        }
    },
    mounted() {
        this.handleGetFileList();
    },
    methods: {
        handleGetFileList() {
            // 获取文件列表
            this.fileList = [];
            let ids = this.files.map(file => file.id);
            fileList({
                pid: this.pid,
                uid: this.$store.state.user.id,
                type: "folder",
                exclude: ids.join(",")
            }).then(res => {
                this.fileList = res.data;
            });
        },
        handleClickRow(row) {
            // 文件夹点击事件
            if (row.type === "folder") {
                // 添加面包屑导航
                this.breadcrumb.push({
                    id: row.id,
                    name: row.name,
                });

                // 修改当前路径
                this.pid = row.id;

                // 加载文件列表
                this.handleGetFileList();
            }
        },
        //更改路径
        changeDirectory(id, index) {
            // 重置面包屑导航
            this.breadcrumb = this.breadcrumb.slice(0, index + 1);
            // 修改当前路径
            this.pid = id;
            // 加载文件列表
            this.handleGetFileList();
        },
        handleConfirm() {
            let ids = this.files.map(file => file.id);
            fileUpdate({
                "id": ids.join(","),
                "pid": this.pid,// 转为字符串
            }).then(res => {
                if (res.code === 0) {
                    this.$emit("confirm", {
                        pid: this.pid
                    });
                    this.$message.success("移动成功");
                } else {
                    this.$message.error(res.msg);
                }
            }).catch(error => {
                this.$message.error(error);
            });
        },
        // 创建目录
        async handleCreateFolder() {
            if (!this.newFile.name) {
                this.$message.error("文件夹名称不能为空");
                return;
            }
            const res = await fileAdd({
                pid: this.pid,
                uid: this.$store.state.user.id,
                name: this.newFile.name,
                type: "folder",
                size: null,
                path: null,
            });
            // console.log(res);
            if (res.code === 0) {
                this.addFolderVisible = false;
                this.newFile.name = "新建文件夹";
                // 成功提示
                this.$message({
                    message: "添加成功",
                    type: "success",
                });
                // 重新加载文件列表
                this.handleGetFileList();
                return res.data;
            } else {
                // 错误提示
                this.$message.error("添加失败");
                return false;
            }
        },
        handleClose() {
            this.$emit("close")
        },
    }
}
</script>

<style scoped>

.file-dir {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    flex-wrap: wrap;
    align-items: center;
    background-color: #ece6f3a1;
    padding: 10px;
}

.file-dir-item:not(:first-child):before {
    content: " > ";
}

.file-dir-item {
    margin-left: 5px;
    font-size: 12px;
}


.file-move-list {
    margin-top: 30px;
    width: 100%;
    height: 300px;
    overflow-y: scroll;
    padding: 10px;
    /*display: flex;*/
    /*flex-direction: column;*/
    /*justify-content: center;*/
    /*align-items: center;*/
}

.file-move-item {
    width: 100%;
    padding: 10px 0;
    border-bottom: 1px solid #ebeef5;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.file-move-item:hover {
    cursor: pointer;
    background-color: rgba(195, 196, 199, 0.52);
}


.file-move-item:active {
    background-color: rgba(195, 196, 199, 0.52);
}

.file-left {
    margin-left: 5px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
}

.file-icon {
    width: 40px;
    height: 40px;
    margin-right: 20px;
    object-fit: cover;
}

.file-name {
    display: flex;
    justify-content: center;
    align-items: center;
}

.op-icon {
    margin-left: 10px;
}

.dialog-footer {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
}
</style>