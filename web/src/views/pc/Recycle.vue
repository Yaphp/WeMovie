<template>
    <el-row>
        <el-col :span="24">
            <div class="header">
                <div class="header-left">
                    <el-breadcrumb :separator-icon="ArrowRight">
                        <el-breadcrumb-item v-for="(bread,index) in breadcrumb" @click="changeDirectory(bread.id,index)">
                            <h3 :class="(index+1) === breadcrumb.length ? 'active-path':'normal-path'">{{ bread.name }}</h3>
                        </el-breadcrumb-item>
                    </el-breadcrumb>
                </div>

                <div class="header-right">
                    <div @click="searchVisible = true">
                        <el-icon :size="22">
                            <Search/>
                        </el-icon>
                    </div>
                </div>
            </div>

            <div class="nav">
                <div class="nav-left">
                    <p class="tips">{{ selectedCountText }}</p>
                </div>
                <div class="nav-right">
                    <el-dropdown @command="handleCommand">
                        <el-button type="primary" size="small" :icon="orderByTypeIcon" plain>
                            按{{ fieldMapping[orderByField] }}排序
                        </el-button>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item command="name" :icon="orderByField==='name' ? Select:''">文件名称
                                </el-dropdown-item>
                                <el-dropdown-item command="size" :icon="orderByField==='size' ? Select:''">文件大小
                                </el-dropdown-item>
                                <el-dropdown-item command="type" :icon="orderByField==='type' ? Select:''">文件类型
                                </el-dropdown-item>
                                <el-dropdown-item command="created_at" :icon="orderByField==='created_at' ? Select:''">
                                    创建时间
                                </el-dropdown-item>
                                <el-dropdown-item command="updated_at" :icon="orderByField==='updated_at' ? Select:''">
                                    修改时间
                                </el-dropdown-item>
                                <el-dropdown-item command="ascending" divided :icon="orderByType==='ascending' ? Select:''">升序
                                </el-dropdown-item>
                                <el-dropdown-item command="descending" :icon="orderByType==='descending' ? Select:''">
                                    降序
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>
            </div>

            <div class="list">
                <el-table ref="fileTable"
                          :data="fileList"
                          :default-sort="{prop: orderByField,order: orderByType}"
                          @select="selectedRow"
                          @select-all="selectedRow"
                          @row-click="handleClickRow"
                          :empty-text="tableEmptyText"
                          style="width: 100%"
                >
                    <el-table-column type="selection" width="50"/>
                    <el-table-column prop="name" label="名称">
                        <template #default="{ row }">
                            <el-dropdown ref="rightmenu" @visible-change="handleRightMenu" trigger="contextmenu">
                                <div class="file-name-icon">
                                    <div class="file-icon">
                                        <el-image loading="lazy" scroll-container=".el-scrollbar__wrap" :key="row.id" :src="getFileIcon(row)" class="file-icon" fit="cover">
                                            <template #placeholder>
                                                <el-icon size="30">
                                                    <Picture/>
                                                </el-icon>
                                            </template>
                                            <template #error>
                                                <img src="/images/loaderr.png" class="file-icon">
                                            </template>
                                        </el-image>
                                    </div>
                                    <div class="file-name">
                                        {{ row.name }}
                                    </div>
                                </div>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item :icon="Download" @click="handleDownloadFile(row)">下载</el-dropdown-item>
                                        <el-dropdown-item :icon="Edit" @click="handleRenameDialog(row)">重命名</el-dropdown-item>
                                        <el-dropdown-item :icon="CopyDocument" @click="handleMoveFile(row)">移动</el-dropdown-item>
                                        <el-dropdown-item :icon="RefreshRight" @click="handleRecoverFile(row)">恢复</el-dropdown-item>
                                        <el-dropdown-item :icon="Delete" @click="handleDeleteFile(row)">删除</el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </template>
                    </el-table-column>
                    <el-table-column prop="size" label="大小">
                        <template #default="{ row }">
                            <div class="file-size">
                                {{ getFileSize(row.size) }}
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="type" label="文件类型"></el-table-column>
                    <el-table-column prop="created_at" sortable label="创建时间"></el-table-column>
                    <el-table-column prop="updated_at" label="修改时间"></el-table-column>
                </el-table>
            </div>
        </el-col>
    </el-row>

    <div class="cl-dialog" v-show="toolbarVisible">
        <div class="toolbar">
            <el-tooltip
                class="box-item"
                effect="dark"
                content="移动"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/move.png" @click="handleMoveFile(null)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="下载"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/download.png" @click="handleDownloadFile(null)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="收藏"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/favorite.png" @click="handleFavoriteFile(null,1)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="取消收藏"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/unfavorite.png" @click="handleFavoriteFile(null,0)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="恢复"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/recover.png" @click="handleRecoverFile(null)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="删除"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/delete.png" @click="handleDeleteFile(null)">
            </el-tooltip>
        </div>
    </div>

    <FileName
        v-if="editDialogVisible"
        :title="editDialogTitle"
        :show="editDialogVisible"
        :file="editFile"
        @close="editDialogVisible = false"
        @confirm="handleRenameFile"
    >
        <template #icon>
            <img :src="getFileIcon(editFile)">
        </template>
    </FileName>

    <el-image-viewer
        v-if="previewImageShow"
        :hide-on-click-modal="true"
        :initial-index="previewImageIndex"
        :url-list="imgSrcList"
        @close="handleClosePreview"
    >
    </el-image-viewer>

    <FileSearch
        :show="searchVisible"
        @search="handleFileSearch"
        @close="searchVisible=false">
    </FileSearch>

    <FileShare
        v-if="shareVisible"
        :show="shareVisible"
        :file="shareFile"
        @close="shareVisible = false">
    </FileShare>

    <FileMove
        v-if="moveVisible"
        :show="moveVisible"
        :file="moveFile"
        @confirm="handleConfirmMove"
        @close="handleCancelMove">
    </FileMove>
</template>

<script>
import {
    ArrowRight,
    CollectionTag,
    CopyDocument,
    Delete,
    Download,
    Edit,
    Plus,
    Search,
    Select,
    Share,
    Sort,
    SortDown,
    SortUp,
    Picture,
    RefreshRight
} from "@element-plus/icons-vue";
import {getFileIcon, getFileSize} from "utils/file.js";
import {fileRecover, fileUpdate} from "api/file.js";
import {fileDelete, fileList, fileFavorite} from "api/file";
import {shallowRef} from "vue";
import FileSearch from "components/pc/FileSearch.vue";
import FileName from "components/pc/FileName.vue";
import FileShare from "components/pc/FileShare.vue";
import FileMove from "components/pc/FileMove.vue";
import {showConfirmDialog} from "vant";


export default {
    name: "Recycle",
    components: {
        Share,
        Search,
        FileSearch,
        FileName,
        FileShare,
        FileMove,
        Picture
    },
    setup() {
        return {
            getFileIcon,
            getFileSize,
            Plus,
            ArrowRight,
            SortUp,
            SortDown,
            Sort,
            Select,
            Download,
            Share,
            CollectionTag,
            Delete,
            Edit,
            RefreshRight,
            CopyDocument
        }
    },
    data() {
        return {
            addDialogTitle: "新建文件夹", // 控制添加文件对话框是否显示
            addFileName: "新建文件夹", // 添加文件夹表单名称
            editDialogVisible: false, // 控制修改文件文件对话框是否显示
            editDialogTitle: "重命名", // 控制修改文件对话框是否显示
            editFile: "", // 待修改文件对象
            pid: "all", // 当前路径父级id pid  默认为0 根目录  all不设置pid
            breadcrumb: [
                {
                    id: "all",
                    name: "回收站",
                }
            ], // 面包屑导航
            fileList: [], // 文件列表
            previewImage: "", // 当前预览图片地址
            previewImageShow: false, // 控制图片预览是否显示
            previewImageIndex: null, // 当前预览图片索引
            imgSrcList: [], // 图片预览列表
            selectedFile: [], // 选中的文件
            selectedCountText: "", // 选中项统计文本
            toolbarVisible: false, // 控制工具栏是否显示
            orderByField: "created_at", // 排序字段
            orderByType: "ascending", // 排序类型
            orderByTypeIcon: shallowRef(SortUp), // 排序类型图标
            fieldMapping: {
                name: "文件名称",
                size: "文件大小",
                created_at: "创建时间",
                updated_at: "修改时间",
                type: "文件类型"
            }, // 排序字段对应的中文
            rightMenuVisible: false, // 控制右键菜单是否显示
            searchVisible: false, // 控制搜索框是否显示
            queryParams: {}, // 文件搜索参数
            tableEmptyText: "暂无数据", // 表格空数据文本
            shareVisible: false, // 控制分享对话框是否显示
            shareFile: null, // 待分享文件
            moveVisible: false, // 控制分移动文件对话框是否显示
            moveFile: null, // 待移动文件
        }
    },
    mounted() {
        //加载文件列表
        this.getFileList();
    },
    methods: {
        //行点击事件
        handleClickRow(row, event, column) {
            // 先把右键菜单关闭
            if (this.rightMenuVisible) {
                this.$refs.rightmenu.handleClose();
                return;
            }

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
                this.getFileList();
                return;
            }

            // 图片点击事件
            if (row.type.indexOf("image") > -1) {
                if (this.imgSrcList.length === 0) {
                    this.imgSrcList = this.fileList.filter(item => item.type.indexOf('image') !== -1).map(item => item.path);
                }
                this.previewImage = row.path;
                this.previewImageIndex = this.imgSrcList.indexOf(this.previewImage);
                this.previewImageShow = true;
                return;
            }

            // 视频点击事件
            if (row.type.indexOf("video") > -1) {
                window.open(row.path, "_blank");
                return;
            }

            // 音频点击事件
            if (row.type.indexOf("audio") > -1) {
                window.open(row.path, "_blank");
                return;
            }

            // 其他文件点击事件处理
            let suffix = row.name.substring(row.name.lastIndexOf('.') + 1).toLowerCase();

            // 文档点击事件
            if (["doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf", "txt", "md"].indexOf(suffix) > -1) {
                window.open(row.path, "_blank");
            }
        },

        // 监听右键菜单的状态
        handleRightMenu(state) {
            this.rightMenuVisible = state;
        },

        // 更改路径
        changeDirectory(id, index) {
            // 重置面包屑导航
            this.breadcrumb = this.breadcrumb.slice(0, index + 1);
            // 修改当前路径
            this.pid = id;
            // 修改工具栏状态
            this.toolbarVisible = false;
            // 加载文件列表
            this.getFileList();
        },

        // 获取文件列表
        getFileList() {
            // 重置文件列表
            this.fileList = [];
            this.imgSrcList = [];
            this.tableEmptyText = '加载中...';
            fileList({
                pid: this.pid,
                uid: this.$store.state.user.id,
                delete_flag: this.pid === 'all' ? 1 : 0,
                ...this.queryParams
            }).then(res => {
                this.tableEmptyText = '暂无数据';
                this.fileList = res.data;
                this.queryParams = {};// 重置查询参数
                // 把每个文件的path转换成绝对路径
                // this.fileList.forEach(item => {
                //     if (item.path) {
                //         item.path = `${this.$store.state.app.host}${item.path}`;
                //     }
                // });
                this.selectedCountText = `共 ${res.data.length} 项`;
            });
        },

        // 关闭图片预览
        handleClosePreview() {
            this.previewImageShow = false;
        },

        // 选中行
        selectedRow(selection, row) {
            // console.log(selection, row)
            this.selectedFile = selection;
            this.selectedCountText = selection.length === 0 ? `共 ${this.fileList.length} 项` : `已选择 ${selection.length} 项`;
            this.toolbarVisible = selection.length > 0;
        },

        // 文件排序
        handleCommand(field) {
            // 如果是排序字段 只改变排序方式 不改变排序字段 重新排序
            if (field === "ascending" || field === "descending") {
                this.orderByType = field;
                this.orderByTypeIcon = field === "ascending" ? shallowRef(SortUp) : shallowRef(SortDown);
                this.$refs.fileTable.sort(this.orderByField, field);
                return;
            }
            this.orderByField = field; // 如果不是排序字段 改变排序字段 重新排序
            this.$refs.fileTable.sort(field, this.orderByType);// 对文件排序
        },

        // 文件搜索
        handleFileSearch(params) {
            params.pid = "all"
            this.queryParams = params;
            this.searchVisible = false;
            this.$router.push({
                path: '/files',
                query: params
            });
        },

        // 移动文件
        handleMoveFile(file) {
            let files = file ? [file] : this.selectedFile;
            if (files.length === 0) {
                this.$message.warning("请选择要移动的文件");
                return;
            }
            this.toolbarVisible = false;
            this.moveFile = files;
            this.moveVisible = true;
        },

        // 处理取消移动的操作
        handleCancelMove() {
            this.moveVisible = false;
            this.toolbarVisible = true;
        },

        // 移动完成操作
        handleConfirmMove() {
            this.moveVisible = false;
            this.toolbarVisible = true
            this.pid = "all";
            this.getFileList();
        },

        // 下载操作
        handleDownload(name, url) {
            let a = document.createElement('a');
            a.href = url;
            a.download = name;
            a.click();
            a.remove();
        },

        // 下载文件
        handleDownloadFile(file) {
            let files = file && file.type ? [file] : this.selectedFile;

            // 判断选中项目数量
            if (files.length === 0) {
                this.$message.warning("请选择要下载的文件");
                return;
            }

            // 文件夹不支持下载
            if (files.some(item => item.type === 'folder')) {
                this.$message.warning("文件夹不支持下载");
                return;
            }

            // 下载文件
            files.forEach(item => {
                let url = `${item.path}`;
                this.handleDownload(item.name, url)
            });
        },

        // 分享文件
        handleShareFile(file) {
            let files = file ? [file] : this.selectedFile;
            if (files.length === 0) {
                this.$message.warning("请选择要分享的文件");
                return;
            }
            this.shareFile = files;
            this.shareVisible = true;
        },

        // 收藏文件
        handleFavoriteFile(file, status) {
            let files = file ? [file] : this.selectedFile;
            if (files.length === 0) {
                this.$message.warning("请选择要收藏的文件");
                return;
            }
            let ids = files.map(item => item.id);
            fileFavorite({"id": ids.join(","), "status": status}).then(res => {
                if (res.code === 0) {
                    // 加载文件列表
                    this.getFileList();
                    this.$message({
                        message: status === 1 ? "收藏成功" : "取消收藏成功",
                        type: "success",
                    });
                } else {
                    // 错误提示
                    this.$message.error(res.msg);
                }
            });
        },

        // 打开重名名对话框
        handleRenameDialog(file) {
            this.editDialogVisible = true;
            this.editFile = file;
            this.editDialogTitle = "重命名";
        },

        // 重命名
        handleRenameFile(file) {
            this.$confirm(`确认将名称修改为${file.name}?`, '重命名', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'info'
            }).then(() => {
                fileUpdate(file).then(res => {
                    this.editDialogVisible = false;
                    if (res.code === 0) {
                        this.getFileList();// 重新加载文件列表
                        this.$message({
                            message: "重命名成功",
                            type: "success",
                        });
                    } else {
                        // 错误提示
                        this.$message.error(res.msg);
                    }
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消'
                    });
                });
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });
            });
        },

        // 恢复文件
        handleRecoverFile(file) {
            let files = file ? [file] : this.selectedFile;

            this.$confirm('确认恢复文件?', '恢复文件', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'info'
            }).then(() => {
                let ids = files.map(item => item.id).join(',');
                fileRecover(ids).then(res => {
                    // console.log(res);
                    if (res.code === 0) {
                        this.getFileList();// 重新加载文件列表
                        this.$message({
                            message: "恢复成功",
                            type: "success",
                        });
                    } else {
                        // 错误提示
                        this.$message.error("恢复失败");
                    }
                    this.activeMenuVisible = false;
                });
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });
            });
        },

        // 删除文件
        handleDeleteFile(file) {
            let files = file ? [file] : this.selectedFile;
            this.$confirm('彻底删除文件?', '彻底删除', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.$notify({
                    title: '删除文件',
                    message: '正在执行删除任务,请稍后...',
                    type: 'info',
                });
                this.toolbarVisible = false;// 隐藏工具栏
                let ids = files.map(item => item.id).join(',');
                fileDelete(ids, 1).then(res => {
                    // console.log(res);
                    if (res.code === 0) {
                        this.getFileList();// 重新加载文件列表
                        this.$notify({
                            title: '删除文件',
                            dangerouslyUseHTMLString: true,
                            message: `删除成功`,
                            type: 'success',
                        });
                    } else {
                        // 错误提示
                        this.$message.error("删除失败");
                    }
                });
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });
            });
        },
    }
    ,
}
</script>

<style>
.header {
    display: flex;
    justify-content: space-between;
    height: 25px;
    /*margin-top: 30px;*/
    width: 96%;
    margin: 30px auto 0 auto;
    align-items: center;
}

.nav {
    display: flex;
    justify-content: space-between;
    height: 30px;
    width: 96%;
    margin: 20px auto 0 auto;
    align-items: center;
}

.tips {
    font-size: 14px;
}

.list {
    display: flex;
    justify-content: space-between;
    /*height: 25px;*/
    /*margin-top: 30px;*/
    width: 96%;
    margin: 30px auto 0 auto;
    align-items: center;
}

.file-name-icon {
    display: flex;
    align-items: center;
    /*justify-content: center;*/
}

.file-icon {
    width: 30px;
    height: 30px;
    margin-right: 20px;
    object-fit: cover;
}

.el-table__row:hover {
    cursor: pointer;
}

.header-left {
    width: 50%;
    display: flex;
    justify-content: flex-start;
}

.header-right {
    width: 60px;
    display: flex;
    justify-content: space-between;
}

.normal-path {
    color: rgba(191, 194, 199, 0.69);
    cursor: pointer;
}

.normal-path:hover {
    color: #5a92da;
    cursor: pointer;
}

.active-path {
    color: #000000;
    cursor: pointer;
}

.el-checkbox__inner {
    border-radius: 50% !important;
}

.cl-dialog {
    position: absolute;
    top: 90vh;
    left: 50%;
    width: 180px;
    height: 36px;
    background-color: black;
    z-index: 2500;
    border-radius: 15px;
    display: flex;
    justify-content: space-around;
}

.toolbar {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
}

.toolbar-icon {
    width: 25px;
    height: 25px;
    /*margin: 5px;*/
    padding: 5px;
    cursor: pointer;
}

.toolbar-icon:hover {
    /*margin: 5px;*/
    background: rgba(255, 255, 255, 0.18);
}
</style>

