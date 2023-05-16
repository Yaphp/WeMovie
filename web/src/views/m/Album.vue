<template>
    <van-nav-bar :title="activeTitle" fixed placeholder border safe-area-inset-top>
        <template #left>
            <div v-if="hasMultChoose">
                <div @click="handleChooseAll">{{ activeText }}</div>
            </div>
            <div v-else>
                <div :style="{fontSize:16+'px'}">
                    相册
                </div>
            </div>
        </template>
        <template #right>
            <div class="nav-toolbar">
                <div class="nav-toolbar-item">
                    <van-icon name="search" size="22" @click="handleSearchJump"/>
                </div>
                <div class="nav-toolbar-item">
                    <el-dropdown @command="handleModeChage" trigger="click">
                        <van-icon name="apps-o" size="22"/>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item command="big" :icon="showMode==='big' ? Select:''">大图</el-dropdown-item>
                                <el-dropdown-item command="medium" :icon="showMode==='medium' ? Select:''">中图</el-dropdown-item>
                                <el-dropdown-item command="small" :icon="showMode==='small' ? Select:''">小图</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
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

    <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="getFileList"
    >
        <van-grid square :border="false" :clickable="true" :gutter="2" :column-num="columnNum">
            <van-grid-item
                v-longpress="{handler:handleLongPress,data:file}"
                v-for="(file,index) in fileList"
                :key="file.id"
                icon="photo-o"
                :class="{'active':file.checked}"
                @click="handleClickRow(file,index)"
            >
                <van-image
                    class="image-item"
                    lazy-load
                    fit="cover"
                    position="top"
                    :src="file.thumb"
                >
                </van-image>
            </van-grid-item>
        </van-grid>
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

    <van-back-top right="10vw" bottom="15vh"></van-back-top>

</template>

<script>
import {fileUpdate} from "api/file.js";
import {Select} from "@element-plus/icons-vue";
import {getFileIcon, getFileSize} from "utils/file.js";
import {fileDelete, fileList, fileFavorite} from "api/file";
import FileShare from "components/m/FileShare.vue";
import FileMove from "components/m/FileMove.vue";
import {showConfirmDialog, showImagePreview} from "vant";

export default {
    name: "File",
    components: {
        FileShare,
        FileMove
    },
    setup() {
        return {
            getFileIcon,
            getFileSize,
            Select
        }
    },
    data() {
        return {
            activeIndex: 'files',// 当前激活的导航栏索引
            activeText: '文件', // 当前激活的导航栏文本
            activeTitle: '', // 当前激活的导航栏标题
            loading: false, // 加载状态
            finished: false, // 文件加载是否完成
            page: 1,
            pid: "all", // 当前路径父级id pid  默认为0 根目录
            fileList: [], // 文件列表
            previewImage: "", // 当前预览图片地址
            previewImageList: [], // 图片预览列表
            previewImageShow: false, // 控制图片预览是否显示
            previewImageIndex: null, // 当前预览图片索引
            selectedFile: [], // 选中的文件
            selectedCountText: "", // 选中项统计文本
            orderByField: "created_at", // 排序字段
            orderByType: "ascending", // 排序类型
            fieldMapping: {
                name: "文件名称",
                size: "文件大小",
                created_at: "创建时间",
                updated_at: "修改时间",
                type: "文件类型"
            }, // 排序字段对应的中文
            queryParams: {
                type: "image", //图片或者视频
                keyword: "",
            }, // 文件搜索参数
            showMode: "medium", // 图片显示模式
            toolbarVisible: false, // 控制工具栏是否显示
            shareVisible: false, // 控制分享对话框是否显示
            shareFile: null, // 待分享文件
            moveVisible: false, // 控制分移动文件对话框是否显示
            moveFile: null, // 待移动文件
            isSelectAll: false, // 是否全选
        }
    },
    computed: {
        columnNum() {
            let modeMapping = {
                "big": 2,
                "medium": 3,
                "small": 4
            };
            return modeMapping[this.showMode]
        },
        hasMultChoose() {
            return this.selectedFile.length > 0;
        },
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
                    this.activeText = "相册";
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
    },
    methods: {
        // 搜索页面
        handleSearchJump() {
            this.$router.push('/search')
        },

        // 切换视图模式 大图/中图/小图
        handleModeChage(field) {
            this.showMode = field;
        },

        // 文件排序
        handleCommand(field) {
            // 如果是排序字段 只改变排序方式 不改变排序字段 重新排序
            if (field === "ascending" || field === "descending") {
                this.orderByType = field;
                this.fileList = [];
                this.page = 1;
                this.getFileList();
                return;
            }
            this.orderByField = field; // 如果不是排序字段 改变排序字段 重新排序
            this.fileList = [];
            this.page = 1;
            this.getFileList();
        },

        // 图片点击事件
        handleClickRow(row, index) {
            console.log(this.hasMultChoose)
            if (this.hasMultChoose) {
                row.checked = !row.checked;
                if (row.checked) {
                    this.selectedFile.push(row);
                } else {
                    this.selectedFile = this.selectedFile.filter(item => item.id !== row.id);
                }
                return;
            }

            showImagePreview({
                images: this.previewImageList,
                startPosition: index
            });
        },

        // 获取文件列表
        getFileList() {
            fileList({
                pid: this.pid,
                uid: this.$store.state.user.id,
                page: this.page,
                order_by: this.orderByField + " " + this.orderByType.replace("ending", ""),
                limit: 20,
                ...this.queryParams
            }).then(res => {
                if (res.data.length === 0) {
                    this.finished = true;
                }
                this.loading = false;
                this.page++; //页码自增
                this.fileList = this.fileList.concat(res.data);
                // this.activeTitle = `共 ${this.fileList.length} 项`;
                this.previewImageList = this.fileList.filter(item => item.type.indexOf('image') !== -1).map(item => item.path);
            });
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
            let files = this.selectedFile;

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
        handleShareFile() {
            let files = this.selectedFile;
            if (files.length === 0) {
                this.$message.warning("请选择要分享的文件");
                return;
            }
            this.shareFile = files;
            this.shareVisible = true;
        },

        // 收藏文件
        handleFavoriteFile(status) {
            let files = this.selectedFile;

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
                    // 修改文件收藏状态
                    files.forEach(item => {
                        item.favorite = status;
                    });
                } else {
                    // 错误提示
                    this.$message.error(res.msg);
                }
            });
        },

        // 批量移动文件
        handleMoveFile() {
            let files = this.selectedFile;
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
        },

        // 移动完成操作
        handleConfirmMove() {
            this.moveVisible = false;
            this.fileList = [];
            this.selectedFile = [];
            this.getFileList();
        },

        // 删除文件
        handleDeleteFile() {
            let files = this.selectedFile;

            showConfirmDialog({
                title: '移入回收站',
                message: '将文件移入回收站?',
            }).then(() => {
                let ids = files.map(item => item.id).join(',');
                fileDelete(ids).then(res => {
                    // console.log(res);
                    if (res.code === 0) {
                        this.fileList = [];
                        this.selectedFile = [];
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
}
</script>

<style scoped>
.nav-toolbar {
    display: flex;
    justify-content: space-between;
    width: 100px;
}

.image-item {
    aspect-ratio: 1;
    width: 100%;
    height: 100%;
}

.bottom-fill-area {
    width: 100%;
    height: 100px;
}

.active {
    border: 2px solid #409EFF;
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
</style>

