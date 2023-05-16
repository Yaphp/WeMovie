<template>
    <el-row>
        <el-col :span="24">
            <div class="header">
                <div class="header-left">
                    <el-breadcrumb :separator-icon="ArrowRight">
                        <el-breadcrumb-item v-for="(bread,index) in breadcrumb">
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
                    <div class="tips">
                        <el-checkbox v-model="isSelectAll" @change="handleSelectAll" @click="handleSelectAll"/>
                        &nbsp;
                        {{ selectedCountText }}
                    </div>
                </div>
                <div class="nav-right">
                    <!--                    <el-dropdown @command="handleFileFilter">-->
                    <!--                        <el-icon size="26">-->
                    <!--                            <Operation/>-->
                    <!--                        </el-icon>-->
                    <!--                        <template #dropdown>-->
                    <!--                            <el-dropdown-menu>-->
                    <!--                                <el-dropdown-item command="image" :icon="queryParams.type==='image' ? Select:''">图片-->
                    <!--                                </el-dropdown-item>-->
                    <!--                                <el-dropdown-item command="video" :icon="queryParams.type==='video' ? Select:''">视频-->
                    <!--                                </el-dropdown-item>-->
                    <!--                            </el-dropdown-menu>-->
                    <!--                        </template>-->
                    <!--                    </el-dropdown>-->

                    <el-dropdown @command="handleShowMode">
                        <el-icon size="26">
                            <Grid/>
                        </el-icon>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item command="large" :icon="showMode==='large' ? Select:''">大图</el-dropdown-item>
                                <el-dropdown-item command="default" :icon="showMode==='default' ? Select:''">中图</el-dropdown-item>
                                <el-dropdown-item command="small" :icon="showMode==='small' ? Select:''">小图</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>

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

            <div class="file-image-list">
                <div :class="row.selected ? 'file-icon file-icon-selected-'+showMode:'file-icon show-mode-'+showMode"
                     :style="{marginRight: index === fileList.length-1 ? 'auto':''}"
                     v-for="(row,index) in fileList">
                    <div :class="row.selected ? 'file-selected':'file-selected'">
                        <el-checkbox v-model="row.selected" @change="handleCheckChange(row)" size="large"/>
                    </div>
                    <el-image loading="lazy"  @click="handleClickRow(row)" :class="'file-image show-mode-'+showMode" :key="row.id" :src="getFileIcon(row)" fit="cover">
                        <template #placeholder>
                            <el-icon size="30">
                                <Picture/>
                            </el-icon>
                        </template>
                        <template #error>
                            <img src="/images/loaderr.png" :class="'file-image show-mode-'+showMode">
                        </template>
                    </el-image>
                    <span v-if="queryParams.type === 'video'">
                        {{ row.name }}
                    </span>
                </div>
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
                content="分享"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/share.png" @click="handleShareFile(null)">
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
    CirclePlusFilled,
    CollectionTag,
    CopyDocument,
    Delete,
    Download,
    Edit,
    Folder,
    Plus,
    Search,
    Select,
    Share,
    Sort,
    SortDown,
    SortUp,
    UploadFilled,
    Picture,
    Operation,
    Grid
} from "@element-plus/icons-vue";
import {getFileIcon, getFileSize} from "utils/file.js";
import {upload, uploadChunk} from "api/upload";
import {fileUpdate} from "api/file.js";
import {fileAdd, fileDelete, fileList, fileFavorite} from "api/file";
import {shallowRef} from "vue";
import FileSearch from "components/pc/FileSearch.vue";
import FileName from "components/pc/FileName.vue";
import FileShare from "components/pc/FileShare.vue";
import FileMove from "components/pc/FileMove.vue";

export default {
    name: "File",
    components: {
        Search,
        CirclePlusFilled,
        FileSearch,
        FileName,
        Share,
        FileShare,
        FileMove,
        Picture,
        Operation,
        Grid
    },
    setup() {
        return {
            getFileIcon,
            getFileSize,
            Plus,
            UploadFilled,
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
            CopyDocument,
            Folder
        }
    },
    data() {
        return {
            editDialogVisible: false, // 控制修改文件文件对话框是否显示
            editDialogTitle: "重命名", // 控制修改文件对话框是否显示
            editFile: "", // 待修改文件对象
            pid: "all", // 当前路径父级id pid  默认为0 根目录
            breadcrumb: [
                {
                    id: 0,
                    name: "相册",
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
            queryParams: {
                type: "image", //图片或者视频
                keyword: "",
            }, // 文件搜索参数
            shareVisible: false, // 控制分享对话框是否显示
            shareFile: null, // 待分享文件
            moveVisible: false, // 控制分移动文件对话框是否显示
            moveFile: null, // 待移动文件
            isSelectAll: false, // 是否全选
            showMode: "default", // 显示模式
        }
    },
    mounted() {
        //加载文件列表
        this.getFileList();
    },
    methods: {
        //行点击事件
        handleClickRow(row) {
            // 先把右键菜单关闭
            if (this.rightMenuVisible) {
                this.$refs.rightmenu.handleClose();
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
            }
        },

        // 监听右键菜单的状态
        handleRightMenu(state) {
            this.rightMenuVisible = state;
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
                ...this.queryParams
            }).then(res => {
                this.fileList = res.data;
                this.selectedCountText = `共 ${res.data.length} 项`;
            });
        },

        // 关闭图片预览
        handleClosePreview() {
            this.previewImageShow = false;
        },

        // 选中行
        handleCheckChange(row) {
            let selection = this.fileList.filter(item => item.selected);
            this.selectedFile = selection;
            this.selectedCountText = selection.length === 0 ? `共 ${this.fileList.length} 项` : `已选择 ${selection.length} 项`;
            this.toolbarVisible = selection.length > 0;
        },

        // 选中所有行
        handleSelectAll() {
            this.isSelectAll = !this.isSelectAll;
            this.fileList.forEach(item => {
                item.selected = this.isSelectAll;
            });
            this.selectedFile = this.isSelectAll ? this.fileList : [];
            this.selectedCountText = this.selectedFile.length === 0 ? `共 ${this.fileList.length} 项` : `已选择 ${this.selectedFile.length} 项`;
            this.toolbarVisible = this.selectedFile.length > 0;
        },

        // 文件筛选
        handleFileFilter(type) {
            this.queryParams.type = type;
            this.queryParams.keyword = "";
            this.getFileList();
        },

        // 展示模式
        handleShowMode(mode) {
            this.showMode = mode;
        },

        // 文件排序
        handleCommand(field) {
            // 如果是排序字段 只改变排序方式 不改变排序字段 重新排序
            if (field === "ascending" || field === "descending") {
                this.orderByType = field;
                this.orderByTypeIcon = field === "ascending" ? shallowRef(SortUp) : shallowRef(SortDown);
                this.fileList.sort((a, b) => {
                    if (field === "ascending") {
                        return a[this.orderByField] > b[this.orderByField] ? 1 : -1;
                    } else {
                        return a[this.orderByField] < b[this.orderByField] ? 1 : -1;
                    }
                });
                return;
            }
            this.orderByField = field; // 如果不是排序字段 改变排序字段 重新排序
            this.fileList.sort((a, b) => {
                if (this.orderByType === "ascending") {
                    return a[this.orderByField] > b[this.orderByField] ? 1 : -1;
                } else {
                    return a[this.orderByField] < b[this.orderByField] ? 1 : -1;
                }
            });
        },

        // 文件搜索
        handleFileSearch(params) {
            params.pid = "all"
            this.queryParams.keyword = params.keyword;
            this.searchVisible = false;
            this.getFileList();
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
        handleConfirmMove(file) {
            this.moveVisible = false;
            this.toolbarVisible = true
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
                    // this.getFileList();
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

        // 删除文件
        handleDeleteFile(file) {
            let files = file ? [file] : this.selectedFile;
            this.$confirm('将文件移入回收站?', '移入回收站', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.toolbarVisible = false;// 隐藏工具栏
                let ids = files.map(item => item.id).join(',');
                fileDelete(ids).then(res => {
                    // console.log(res);
                    if (res.code === 0) {
                        this.getFileList();// 重新加载文件列表
                        this.$message({
                            message: "删除成功",
                            type: "success",
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

<style scoped>
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

.nav-right {
    width: 250px;
    display: flex;
    justify-content: space-evenly;
}

.tips {
    font-size: 15px;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    height: 30px;
    width: 120px;
}

.file-image-list {
    display: flex;
    justify-content: space-between;
    width: 96%;
    flex-direction: row;
    height: auto;
    flex-wrap: wrap;
    margin: 30px auto 0 auto;
    align-items: center;
    overflow: hidden;
}

.file-icon {
    width: 80px;
    height: 80px;
    margin: 12px;
    object-fit: cover;
}


.file-icon-selected > .file-image {
    transform: scale(0.5);
    transition: all 0.3s ease-in-out;
}

.file-icon:hover > .file-selected {
    opacity: 1;
}


.file-image {
    width: 80px;
    height: 80px;
    margin: 6px;
    cursor: pointer;
}

.show-mode-large {
    width: 160px;
    height: 160px;
}

.file-icon-selected-large {
    width: 200px;
    height: 200px;
    border: 3px solid #409EFF;
}

.show-mode-default {
    width: 80px;
    height: 80px;
}

.file-icon-selected-default {
    width: 120px;
    height: 120px;
    border: 3px solid #409EFF;
}

.show-mode-small {
    width: 40px;
    height: 40px;
}

.file-icon-selected-small {
    width: 80px;
    height: 80px;
    border: 3px solid #409EFF;
}

.file-selected {
    position: relative;
    top: 25px;
    width: 15px;
    height: 15px;
    left: 15px;
    z-index: 999;
    opacity: 0;
}

.file-selected-hover {
    cursor: pointer;
    opacity: 1;
}

.el-table__row:hover {
    cursor: pointer;
}

.el-image > img {
    object-fit: cover;
    object-position: top !important;
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

