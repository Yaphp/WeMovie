<template>
    <van-nav-bar :title="activeTitle" fixed placeholder border safe-area-inset-top>
        <template #left>
            <div v-if="hasMultChoose">
                <div @click="handleChooseAll">{{ activeText }}</div>
            </div>
            <div v-else>
                <div :style="{fontSize:16+'px'}">
                    收藏夹
                </div>
            </div>
        </template>
        <template #right>
            <div class="nav-toolbar">
                <div class="nav-toolbar-item">
                    <van-icon name="search" size="22" @click="handleSearchJump"/>
                </div>
                <div class="nav-toolbar-item">
                    <el-dropdown @command="handleCommand" trigger="click">
                        <van-icon :name="orderByType" size="22"/>
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
                                <el-dropdown-item command="ascending" divided
                                                  :icon="orderByType==='ascending' ? Select:''">升序
                                </el-dropdown-item>
                                <el-dropdown-item command="descending" :icon="orderByType==='descending' ? Select:''">
                                    降序
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>
            </div>
        </template>
    </van-nav-bar>

    <van-skeleton :loading="loading" title :row="20"/>

    <van-list
        v-model:loading="loading"
        :finished="true"
        finished-text="没有更多了"
        @load="getFileList"
    >
        <div class="file-list">
            <div class="file-dir" v-if="pid!=='all'">
                <div class="file-dir-item" v-for="(bread,index) in breadcrumb" @click="changeDirectory(bread.id,index)">
                    {{ bread.name }}
                </div>
            </div>

            <div class="file-item" v-longpress="{handler:handleLongPress,data:file}" v-for="file in fileList" @click="handleClickRow(file)">
                <div class="file-left">
                    <van-checkbox v-if="hasMultChoose" class="van-checkbox-right-margin" v-model="file.checked" @click.stop="handleCheckFile(file)"/>
                    <!--                    <img class="file-icon" :src="getFileIcon(file)" :alt="file.name">-->
                    <van-image
                        class="file-icon"
                        lazy-load
                        fit="cover"
                        position="top"
                        :src="getFileIcon(file)"
                    >
                    </van-image>
                    <div class="file-name">
                        {{ getShortFileName(file.name) }}
                    </div>
                </div>
                <div class="file-right" @click.stop="handleShowActionSheet(file)">...</div>
            </div>
        </div>
    </van-list>

    <div class="bottom-fill-area"></div>

    <div class="cl-dialog" v-show="hasMultChoose">
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
                <img class="toolbar-icon" src="/images/favorite.png" @click="handleFavoriteFile(1)">
            </el-tooltip>

            <el-tooltip
                class="box-item"
                effect="dark"
                content="取消收藏"
                placement="top"
            >
                <img class="toolbar-icon" src="/images/unfavorite.png" @click="handleFavoriteFile(0)">
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

    <van-back-top right="10vw" bottom="15vh"></van-back-top>

    <van-overlay :show="previewImageShow" @click="previewImageShow = false">
        <div class="wrapper" @click="previewImageShow = false">
            <video class="wrapper-video" :src="previewImage" autoplay controls></video>
        </div>
    </van-overlay>

    <van-action-sheet class="action-sheet-bgcolor" safe-area-inset-bottom v-model:show="activeMenuVisible">
        <div class="header-fill"></div>
        <van-cell-group class="cell-group-margin" inset>
            <van-cell title="下载" @click="handleDownloadFile">
                <template #right-icon>
                    <van-icon name="down" class="action-icon"/>
                </template>
            </van-cell>
        </van-cell-group>

        <van-cell-group class="cell-group-margin" inset>
            <van-cell title="分享" @click="handleShareFile">
                <template #right-icon>
                    <van-icon name="share-o" class="action-icon"/>
                </template>
            </van-cell>
            <van-cell :title="activeFile.favorite === 1 ? '取消收藏':'收藏'" @click="handleFavoriteFile(activeFile.favorite === 1 ? 0:1)">
                <template #right-icon>
                    <van-icon :name="activeFile.favorite === 1 ? 'like':'like-o'" class="action-icon"/>
                </template>
            </van-cell>
        </van-cell-group>

        <van-cell-group class="cell-group-margin" inset>
            <van-cell title="重命名" @click="editDialogVisible = true">
                <template #right-icon>
                    <van-icon name="edit" class="action-icon"/>
                </template>
            </van-cell>
            <van-cell title="移动" @click="handleMoveFile">
                <template #right-icon>
                    <van-icon name="exchange" class="action-icon"/>
                </template>
            </van-cell>
            <van-cell title="移入回收站" @click="handleDeleteFile" title-style="color:red;">
                <template #right-icon>
                    <van-icon name="delete" color="red" class="action-icon"/>
                </template>
            </van-cell>
        </van-cell-group>
    </van-action-sheet>

    <van-dialog v-model:show="editDialogVisible" :title="editDialogTitle" show-cancel-button @confirm="handleConfirmRenameFile" @cancel="handleCancelRenameFile">
        <van-cell-group inset>
            <van-field class="field-bgcolor" v-model="activeFile.name" clearable/>
        </van-cell-group>
    </van-dialog>

    <van-dialog v-model:show="uploadDialogVisible" width="100px" :show-confirm-button="false">
        <van-circle
            v-model:current-rate="uploadPercent"
            :rate="100"
            :text="uploadPercent + '%'"
        />
    </van-dialog>

    <FileMove
        v-if="moveVisible"
        :show="moveVisible"
        :file="moveFile"
        @confirm="handleConfirmMove"
        @close="handleCancelMove">
    </FileMove>

    <FileShare
        v-if="shareVisible"
        :show="shareVisible"
        :file="shareFile"
        @close="shareVisible = false">
    </FileShare>
</template>

<script>
import {getFileIcon, getFileSize, getShortFileName} from "utils/file.js";
import {
    Select
} from "@element-plus/icons-vue";
import {upload, uploadChunk} from "api/upload";
import {fileList, fileAdd, fileDelete, fileUpdate, fileFavorite} from "api/file";
import {showImagePreview} from 'vant';
import {showConfirmDialog} from 'vant';
import FileMove from "components/m/FileMove.vue";
import FileShare from "components/m/FileShare.vue"

export default {
    name: "File",
    components: {
        FileMove,
        FileShare
    },
    setup() {
        return {
            getFileIcon,
            getFileSize,
            getShortFileName,
            Select
        }
    },
    data() {
        return {
            activeIndex: 'files',
            activeText: '收藏夹',
            activeTitle: '',
            loading: true, // 控制加载动画是否显示
            addDialogVisible: false, // 控制添加文件对话框是否显示
            addDialogTitle: "新建文件夹",//修改文件名称对话框标题
            addFolderTitle: "新建文件夹",//修改文件名称对话框标题
            uploadDialogVisible: false, // 控制上传文件进度条是否显示
            pid: "all", // 当前路径父级id pid  默认为0 根目录
            breadcrumb: [{id: "all", name: "收藏夹"}], // 面包屑导航
            fileList: [], // 文件列表
            previewImage: [], // 当前预览图片地址
            previewImageShow: false, // 控制图片预览是否显示
            imgSrcList: [], // 图片预览列表
            selectedFile: [], // 选中的文件
            orderByField: "created_at", // 排序字段
            orderByType: "ascending", // 排序类型
            fieldMapping: {
                name: "文件名称",
                size: "文件大小",
                created_at: "创建时间",
                updated_at: "修改时间",
                type: "文件类型"
            }, // 排序字段对应的中文
            activeMenuVisible: false, // 控制底部菜单是否显示
            activeFile: {}, // 当前右边菜单选中的文件
            activeOldFile: {}, // 保存当前右边菜单选中的文件的原始数据
            editDialogTitle: "重命名",//修改文件名称对话框标题
            editDialogVisible: false, //展示修改文件名称对话框
            moveVisible: false, // 控制移动文件对话框是否显示
            moveFile: [], // 当前移动的文件
            shareVisible: false, // 控制分享文件对话框是否显示
            shareFile: [], // 当前分享的文件
        }
    },
    computed: {
        hasMultChoose() {
            return this.selectedFile.length > 0;
        },
        hasChoosedCount() {
            return this.selectedFile.length;
        }
    },
    watch: {
        selectedFile: {
            handler(newValue, oldValue) {
                if (newValue.length > 0) {
                    this.activeTitle = `已选中${newValue.length}项`;
                    if (newValue.length === this.fileList.length) {
                        this.activeText = "取消全选";
                    } else {
                        this.activeText = "全选";
                    }
                } else {
                    this.activeText = "收藏夹";
                    this.activeTitle = "";
                }
            },
            flush: 'post',
            deep: true
        },
    },
    directives: {
        longpress: {
            created(el, binding) {
                let timer = null;
                let startTime = 0;

                const handleStart = (e) => {
                    startTime = new Date().getTime();
                    timer = setTimeout(() => {
                        binding.value.handler(binding.value.data);
                    }, 500);
                };

                const handleEnd = () => {
                    clearTimeout(timer);
                    if (new Date().getTime() - startTime < 500) {
                        // 这里是点击事件 不做处理
                    }
                };
                el.addEventListener('touchstart', handleStart, {passive: true});
                el.addEventListener('touchend', handleEnd, {passive: true});
            },
        },
    },
    mounted() {
        //加载文件列表
        this.getFileList();
    },
    methods: {
        handleSearchJump() {
            this.$router.push('/search')
        },

        // 获取文件列表
        getFileList() {
            this.fileList = [];
            this.selectedFile = [];
            fileList({
                pid: this.pid,
                uid: this.$store.state.user.id,
                favorite: this.pid === 'all' ? 1 : 0
            }).then(res => {
                // 加载动画关闭
                this.loading = false;
                //设置文件数据
                this.fileList = res.data;
                this.imgSrcList = res.data.filter(item => item.type.indexOf('image') !== -1).map(item => item.path);
            });
        },

        //行点击事件
        handleClickRow(row) {
            // 多选状态下点击行 选中或取消选中
            if (this.hasMultChoose) {
                row.checked = !row.checked;
                if (row.checked) {
                    this.selectedFile.push(row);
                } else {
                    this.selectedFile = this.selectedFile.filter(item => item.id !== row.id);
                }
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
                showImagePreview({
                    images: this.imgSrcList,
                    startPosition: this.imgSrcList.indexOf(row.path),
                    // closeable: true,
                });
                return;
            }

            // 视频点击事件
            if (row.type.indexOf("video") > -1) {
                this.previewImage = row.path;
                this.previewImageShow = true;
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

        // 监听菜单的状态
        handleShowActionSheet(file) {
            this.activeMenuVisible = true;
            this.activeFile = file;
            this.activeOldFile = {...file};
        },

        //更改路径
        changeDirectory(id, index) {
            // 重置面包屑导航
            this.breadcrumb = this.breadcrumb.slice(0, index + 1);

            // 修改当前路径
            this.pid = id;

            // 修改工具栏状态
            this.activeMenuVisible = false;

            // 加载文件列表
            this.getFileList();
        },

        // 长按事件 将当前文件设为选中文件
        handleLongPress(file) {
            file.checked = true;
            this.selectedFile.push(file);
        },

        // 选择全部
        handleChooseAll() {
            if (this.activeText === '全选') {
                this.selectedFile = this.fileList;
                this.selectedFile.forEach(item => item.checked = true);
                this.activeText = '取消全选'
            } else {
                this.selectedFile.forEach(item => item.checked = false);
                this.selectedFile = [];
                this.activeText = '全选'
            }
        },

        // 监听文件选中状态
        handleCheckFile(file) {
            if (file.checked) {
                this.selectedFile.push(file);
            } else {
                this.selectedFile = this.selectedFile.filter(item => item.id !== file.id);
            }
        },

        // 文件排序
        handleCommand(field) {
            // 如果是排序字段 只改变排序方式 不改变排序字段 重新排序
            if (field === "ascending" || field === "descending") {
                this.orderByType = field;
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

        // 下载操作
        handleDownload(name, url) {
            let a = document.createElement('a');
            a.href = url;
            a.download = name;
            a.click();
            a.remove();
        },

        // 下载文件
        handleDownloadFile() {
            let files = this.activeFile && this.activeFile.type ? [this.activeFile] : this.selectedFile;

            // 判断选中项目数量
            if (files.length === 0) {
                this.$message.warning("请选择要下载的文件");
                return;
            }

            // 文件夹不支持下载
            if (files.some(item => item.type === 'folder')) {
                this.activeMenuVisible = false;
                this.$message.warning("文件夹不支持下载");
                return;
            }

            // 下载文件
            files.forEach(item => {
                let url = `${item.path}`;
                this.handleDownload(item.name, url)
            });

            this.activeMenuVisible = false;
        },

        // 分享文件
        handleShareFile() {
            let files = this.hasMultChoose ? this.selectedFile : [this.activeFile];
            if (files.length === 0) {
                this.$message.warning("请选择要分享的文件");
                this.activeMenuVisible = false;
                return;
            }
            this.activeMenuVisible = false;
            this.shareFile = files;
            this.shareVisible = true;
        },

        // 收藏文件
        handleFavoriteFile(status) {
            let files = this.activeFile && this.activeFile.type ? [this.activeFile] : this.selectedFile;

            if (files.length === 0) {
                this.$message.warning("请选择要收藏的文件");
                return;
            }

            let ids = files.map(item => item.id);
            fileFavorite({"id": ids.join(","), "status": status}).then(res => {
                if (res.code === 0) {
                    // 加载文件列表
                    this.fileList = [];
                    this.getFileList();
                    this.$message({
                        message: status === 1 ? "收藏成功" : "取消收藏成功",
                        type: "success",
                    });
                    this.activeMenuVisible = false;
                } else {
                    // 错误提示
                    this.$message.error(res.msg);
                }
            });
        },

        // 重命名
        handleConfirmRenameFile() {
            this.editDialogVisible = false;
            this.activeMenuVisible = false;
            fileUpdate(this.activeFile).then(res => {
                if (res.code === 0) {
                    this.$message({
                        message: "重命名成功",
                        type: "success",
                    });
                } else {
                    // 回复原来的文件名
                    this.activeFile.name = this.activeOldFile.name;

                    // 错误提示
                    this.$message.error(res.msg);
                }
            }).catch(err => {
                this.activeFile.name = this.activeOldFile.name;
                this.$message({
                    type: 'error',
                    message: err.msg
                });
            });
        },

        // 取消重命名
        handleCancelRenameFile() {
            this.activeFile.name = this.activeOldFile.name;
            this.editDialogVisible = false;
            this.activeMenuVisible = false;
            this.$message({
                type: 'info',
                message: '已取消'
            });
        },

        // 批量移动文件
        handleMoveFile() {
            let files = this.hasMultChoose ? this.selectedFile : [this.activeFile];
            if (files.length === 0) {
                this.$message.warning("请选择要移动的文件");
                return;
            }
            this.moveFile = files;
            this.moveVisible = true;
        },

        // 处理取消移动的操作
        handleCancelMove() {
            this.moveVisible = false;
            this.activeMenuVisible = false;
        },

        // 移动完成操作
        handleConfirmMove() {
            this.moveVisible = false;
            this.activeMenuVisible = false;
            this.pid = 0;
            this.getFileList();
        },

        // 删除文件
        handleDeleteFile() {
            let files = this.activeFile && this.activeFile.type ? [this.activeFile] : this.selectedFile;

            showConfirmDialog({
                title: '移入回收站',
                message: '将文件移入回收站?',
            }).then(() => {
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
                    this.activeMenuVisible = false;
                });
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });
            });
        },
    },
}
</script>

<style scoped>
.file-list {
    width: 100%;
    height: 100%;
    overflow: auto;
    padding: 10px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.file-dir {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
    align-items: center;
    background-color: #ece6f3a1;
    padding: 10px;
}

.file-dir-item:not(:first-child):before {
    content: " > ";
}

.file-dir-item {
    margin-left: 5px;
}


.file-item {
    width: 100%;
    padding: 10px 0;
    border-bottom: 1px solid #ebeef5;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.file-item:active {
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
    object-position: top;
}

.file-right {
    margin-right: 5px;
}

.action-sheet-bgcolor {
    background-color: #f7f8fa;
}

.header-fill {
    width: 100%;
    height: 20px;
}

.cell-group-margin {
    margin-bottom: 10px;
}

.action-icon {
    font-size: 16px;
    line-height: inherit;
}

.field-bgcolor {
    background-color: #eaecf1;
    margin: 25px auto;
}

.van-checkbox-right-margin {
    margin-right: 10px;
}

.cl-dialog {
    position: fixed;
    top: 78vh;
    left: calc(50% - 170px);
    width: 340px;
    height: 50px;
    background-color: black;
    z-index: 2000;
    border-radius: 15px;
    display: flex;
    justify-content: space-around;
}

.toolbar {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    width: 100%;
}

.toolbar-icon {
    width: 30px;
    height: 30px;
    /*margin: 5px;*/
    padding: 5px;
    cursor: pointer;
}

.toolbar-icon:hover {
    /*margin: 5px;*/
    background: rgba(255, 255, 255, 0.18);
}

.nav-toolbar {
    display: flex;
    justify-content: space-between;
    width: 50px;
}

.van-uploader-width, .van-uploader__wrapper, .van-uploader__input-wrapper {
    width: 100%;
}

.bottom-fill-area {
    width: 100%;
    height: 100px;
}

.wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
}

.wrapper-video {
    width: 100%;
    height: 300px;
}
</style>

