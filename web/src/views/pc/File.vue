<template>
    <el-row>
        <el-col :span="24">
            <div class="header">
                <div class="header-left">
                    <el-breadcrumb :separator-icon="ArrowRight">
                        <el-breadcrumb-item v-for="(bread,index) in breadcrumb"
                                            @click="changeDirectory(bread.id,index)">
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
                    <div>
                        <el-dropdown>
                            <el-icon :size="24">
                                <CirclePlusFilled/>
                            </el-icon>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item :icon="Plus" @click="addDialogVisible = true">新建文件夹
                                    </el-dropdown-item>
                                    <el-dropdown-item :icon="UploadFilled">
                                        <el-upload
                                            ref="upload"
                                            class="upload-demo"
                                            :http-request="handleChooseFile"
                                            :before-upload="beforeChooseFile"
                                            :show-file-list="false"
                                            v-model:file-list="uploadFileList"
                                            multiple
                                            :limit="5"
                                        >上传文件
                                        </el-upload>
                                    </el-dropdown-item>
                                    <el-dropdown-item :icon="Folder" @click.native.stop="handleChooseFolder">
                                        上传文件夹
                                    </el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
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

            <div class="file-list">
                <el-table ref="fileTable"
                          :data="fileList"
                          :default-sort="{prop: orderByField,order: orderByType}"
                          @select="selectedRow"
                          @select-all="selectedRow"
                          @row-click="handleClickRow"
                          @row-contextmenu="handleRightClickRow"
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
                                        <el-dropdown-item :icon="Download" @click="handleDownloadFile(row)">下载
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="Share" @click="handleShareFile(row)">分享
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="CollectionTag" v-if="row.favorite===0"
                                                          @click="handleFavoriteFile(row,1)">收藏
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="CollectionTag" v-if="row.favorite===1"
                                                          @click="handleFavoriteFile(row,0)">取消收藏
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="Edit" @click="handleRenameDialog(row)">重命名
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="CopyDocument" @click="handleMoveFile(row)">移动
                                        </el-dropdown-item>
                                        <el-dropdown-item :icon="Delete" @click="handleDeleteFile(row)">删除
                                        </el-dropdown-item>
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
        v-if="addDialogVisible"
        :title="addDialogTitle"
        :show="addDialogVisible"
        @close="addDialogVisible = false"
        @confirm="handleCreateFolder"
    ></FileName>

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

    <ProgressNotify
        v-for="(item,index) in uploadFileList"
        v-show="item.show"
        :style="{top: (85-(index*12))+'vh'}"
        :title="item.name"
        :percentage="item.percentage"
        :speed="item.speed"
        @completed="handleUploadCompleted(item)"
        @close="item.show=false">
    </ProgressNotify>

    <ProgressNotify
        v-show="mulitFileUpload"
        :title="folderUploadTitle"
        :percentage="folderUploadPercent"
        :speed="folderUploadSpeed"
        @close="mulitFileUpload=false">
    </ProgressNotify>

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
    Picture
} from "@element-plus/icons-vue";
import {fileTransferSpeed, getFileIcon, getFileSize} from "utils/file.js";
import {upload, uploadChunk} from "api/upload";
import {fileUpdate} from "api/file.js";
import {fileAdd, fileDelete, fileList, fileFavorite} from "api/file";
import {shallowRef} from "vue";
import ProgressNotify from "components/pc/ProgressNotify.vue";
import FileSearch from "components/pc/FileSearch.vue";
import FileName from "components/pc/FileName.vue";
import FileShare from "components/pc/FileShare.vue";
import FileMove from "components/pc/FileMove.vue";


export default {
    name: "File",
    components: {
        Search,
        CirclePlusFilled,
        ProgressNotify,
        FileSearch,
        FileName,
        Share,
        FileShare,
        FileMove,
        Picture
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
            addDialogVisible: false, // 控制添加文件文件对话框是否显示
            addDialogTitle: "新建文件夹", // 控制添加文件对话框是否显示
            addFileName: "新建文件夹", // 添加文件夹表单名称
            editDialogVisible: false, // 控制修改文件文件对话框是否显示
            editDialogTitle: "重命名", // 控制修改文件对话框是否显示
            editFile: "", // 待修改文件对象
            pid: this.$route.query.pid || 0, // 当前路径父级id pid  默认为0 根目录
            breadcrumb: [
                {
                    id: 0,
                    name: "文件",
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
            hoverRow: null, // 鼠标悬浮的行
            rightMenuVisible: false, // 控制右键菜单是否显示
            mulitFileUpload: false, // 是否多文件、文件夹上传
            uploadFileList: [], // 上传文件列表
            folderUploadTitle: "", // 文件夹上传标题
            folderUploadCount: 0, // 文件夹下已上传文件数
            folderUploadTotal: 0, // 文件夹总文件数
            folderUploadPercent: 0, // 文件夹已上传文件百分比
            folderUploadSpeed: '', // 文件夹已上传文件上传速度
            folderUploadSize: 0, // 文件夹已上传文件大小
            folderUploadStart: '', // 文件夹上传开始时间
            searchVisible: false, // 控制搜索框是否显示
            queryParams: {
                type: this.$route.query.type || "",
                keyword: this.$route.query.keyword || "",
            }, // 文件搜索参数
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

        // 鼠标双击事件
        handleDblClickRow(row, event, column) {
        },

        // 鼠标右键
        handleRightClickRow(row, column, event) {
            // event.preventDefault();
        },

        // 创建目录
        async handleCreateFolder(file) {
            if (file.name === "") {
                this.$message.error("文件夹名称不能为空");
                return;
            }
            const res = await fileAdd({
                pid: this.pid,
                uid: this.$store.state.user.id,
                name: file.name,
                type: "folder",
                size: null,
                path: null,
            });

            console.log(res);
            if (res.code === 0) {
                this.addDialogVisible = false;// 关闭对话框
                this.addFileName = "新建文件夹";// 清空表单
                this.getFileList();// 重新加载文件列表
                // 提示
                this.$message({
                    message: "添加成功",
                    type: "success",
                });
                return res.data;
            } else {
                // 错误提示
                this.$message.error("添加失败");
                return false;
            }
        },

        // 选择上传文件夹
        async handleChooseFolder() {
            try {
                // 打开文件夹选择器
                const handle = await showDirectoryPicker();

                // 处理文件句柄
                const root = await this.handleFileHandle(handle);

                // 统计文件夹总文件数
                this.folderUploadTotal = await this.handleCountFolder(root);

                // 设置文件夹上传标题
                this.folderUploadTitle = `${root.name}下的${this.folderUploadTotal}个文件`;

                // 开启多文件上传状态
                this.mulitFileUpload = true;

                // 记录开始上传时间
                this.folderUploadStart = new Date().getTime();

                // 开始上传文件夹
                const upload = await this.handleUploadFolder(root, true);

                // 上传成功
                if (upload) {
                    this.pid = 0; // 重置pid
                    this.folderUploadTitle = ""; // 重置文件夹上传标题
                    this.folderUploadCount = 0; // 重置文件夹已上传文件数
                    this.folderUploadTotal = 0; // 重置文件夹总文件数
                    this.folderUploadPercent = 0; // 重置文件夹已上传文件数百分比
                    this.folderUploadStart = ""; // 重置文件夹上传开始时间
                    this.folderUploadSpeed = ""; // 重置文件夹上传速度
                    this.folderUploadSize = ""; // 重置文件夹上传大小
                    this.breadcrumb = [{id: 0, name: "文件"}]; // 重置面包屑导航
                    this.getFileList();// 重新加载文件列表
                    this.mulitFileUpload = false; // 重置多文件上传状态
                    this.$message({
                        message: "上传成功",
                        type: "success",
                    });
                } else {
                    this.$message.error("上传失败");
                }
            } catch (err) {
                console.debug(err)
                // 用户拒绝授权时给予提示
                this.$message.info('已取消');
            }
        },

        // 统计文件夹下总文件数
        async handleCountFolder(root) {
            let count = 0; // 文件总数

            // 子目录
            let subDirectory = [];

            // 遍历文件夹
            for await (const entry of root.values()) {
                if (entry.kind === "directory") {
                    subDirectory.push(entry);
                } else {
                    count++;
                }
            }

            // 遍历子目录
            for (let i = 0; i < subDirectory.length; i++) {
                count += await this.handleCountFolder(subDirectory[i]);
            }

            return count;
        },

        // 遍历文件夹 依次上传
        async handleUploadFolder(root, first = false) {
            let pid = this.pid;

            if (root.kind === "directory") {
                // 创建文件夹
                const res = await fileAdd({
                    pid: pid,
                    uid: this.$store.state.user.id,
                    name: root.name,
                    type: "folder",
                    size: null,
                    path: null,
                });

                // 判断是否创建成功
                if (res.code === 1) {
                    this.$message.error("创建文件夹失败");
                    return false;
                }

                // 获取新创建的文件夹id
                pid = res.data;

                // 首次上传切换根目录的pid
                if (first) {
                    this.pid = pid;
                }
            }

            // 子目录
            let subDirectory = [];

            for (let i = 0; i < root.children.length; i++) {
                if (root.children[i].kind === 'file') {
                    const file = await root.children[i].getFile();

                    try {
                        // 兼容handleUpload方法 必须包裹一层file
                        const res = await this.handleUpload({file: file}, pid);

                        if (res !== false) {
                            // 更新进度条
                            this.folderUploadCount += 1;
                            this.folderUploadSize += file.size;
                            this.folderUploadSpeed = fileTransferSpeed(this.folderUploadSize, new Date().getTime(), this.folderUploadStart);
                            this.folderUploadPercent = Math.floor(this.folderUploadCount / this.folderUploadTotal * 100);
                        }
                    } catch (e) {
                        // 错误提示
                        this.$message.error("上传错误:" + e.message);
                    }
                } else {
                    subDirectory.push(root.children[i]);
                }
            }

            // 先将文件上传完成 在对文件目录 递归上传
            for (let i = 0; i < subDirectory.length; i++) {
                await this.handleUploadFolder(subDirectory[i]);
            }

            return true;
        },

        // 处理文件句柄
        async handleFileHandle(handle) {
            if (handle.kind === 'file') {
                return handle;
            }
            handle.children = [];
            const iter = handle.entries();
            for await (const item of iter) {
                handle.children.push(await this.handleFileHandle(item[1]));
            }
            return handle;
        },

        beforeChooseFile() {
            // 限制文件上传数量
            // this.$message.warning(`请注意,最多同时上传5个文件`);

            // 多文件上传时禁止选择文件
            if (this.mulitFileUpload) {
                this.$message.warning('请等待文件夹上传完成');
                return false;
            }
        },

        // 选择上传文件
        async handleChooseFile(file) {
            this.mulitFileUpload = false;

            // 初始化上传进度条
            this.uploadFileList = this.uploadFileList.map((item, index) => {
                item.show = true;
                item.percentage = 0;
                return item;
            })

            try {
                const res = await this.handleUpload(file);

                if (res) {
                    // this.$refs.upload.handleRemove(file); // 清空已上传的文件
                    // 从上传列表中移除已上传的文件
                    this.uploadFileList = this.uploadFileList.filter((item, index) => {
                        // console.log(item.name, file.file.name)
                        return item.name !== file.file.name;
                    })
                    // console.log(this.uploadFileList)
                    this.$message({
                        message: "上传成功",
                        type: "success",
                    });
                    if (this.uploadFileList.length === 0) {
                        this.getFileList();// 重新加载文件列表
                    }
                } else {
                    this.$message.error("上传失败");
                }
            } catch (e) {
                // 错误提示
                this.$message.error("上传错误:" + e.message);
            }
        },

        // 上传文件至服务器
        async handleUpload(file, pid = '') {
            const formData = new FormData()

            // 判断文件大小超过超过1m 进行分片上传
            if (file.file.size > this.$store.state.chunkSize) {
                // 分片上传
                return await this.handleUploadChunk(file, pid);
            }

            formData.append('file', file.file) // 传入FILE文件
            formData.append('pid', pid || this.pid) // 传入父级id
            formData.append('uid', this.$store.state.user.id) // 传入用户id
            formData.append('name', file.file.name) // 传入文件名
            formData.append('type', file.file.type) // 传入文件类型
            formData.append('size', file.file.size) // 传入文件大小

            // 上传文件
            const res = await upload(formData);

            // 判断是否上传成功
            return res.code === 0;
        },

        async handleUploadChunk(file, pid = '') {
            let _this = this;

            // 按照1m分割文件
            let result;
            let path;//文件上传后分配的路径路径
            let fileIndex; //文件指针
            const chunkSize = this.$store.state.chunkSize; // 切片大小
            const chunkCount = Math.ceil(file.file.size / chunkSize);
            let startTime = new Date().getTime();// 开始上传时间

            for (let i = 0; i < chunkCount; i++) {
                let chunk = file.file.slice(i * chunkSize, (i + 1) * chunkSize);

                // 上传切片
                result = await uploadChunk({
                    name: file.file.name,
                    chunk: chunk,
                    chunk_number: i,
                    chunk_total: chunkCount,
                    path: path,
                });

                if (result.code === 0) {
                    // 保存切片路径
                    path = result.data;
                    //非文件夹上传模式 更新进度条
                    if (!_this.mulitFileUpload) {
                        fileIndex = _this.uploadFileList.findIndex((item) => {
                            return item.name === file.file.name;
                        });
                        if (fileIndex !== -1) {

                            // 偶数次上传计算上传速度
                            if (i % 2 === 0) {
                                let endTime = new Date().getTime(); // 计算上传速度
                                let totalChunkSize = i * chunkSize;//总上传量
                                _this.uploadFileList[fileIndex].speed = fileTransferSpeed(totalChunkSize, endTime, startTime); // 计算上传速度
                            }
                            _this.uploadFileList[fileIndex].percentage = parseInt((((i + 1) / chunkCount) * 100).toFixed(0));
                        }
                    }
                } else {
                    _this.$message.error(`${file.file.name}切片上传失败`);
                    result = false;
                    break;
                }
            }

            // 合并切片
            result = await uploadChunk({
                pid: pid || this.pid,
                uid: this.$store.state.user.id,
                name: file.file.name,
                type: file.file.type,
                size: file.file.size,
                path: path,
                merge: true
            });

            if (result.code === 0) {
                result = true;
            } else {
                _this.$message.error(`${file.file.name}合并文件失败`);
                result = false;
            }
            return result;
        },

        // 上传完成
        handleUploadCompleted(file) {
            this.$refs.upload.handleRemove(file); // 清空已上传的文件
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

        // 取消文件上传按钮的焦点状态
        unFocus() {
            const button = document.createElement('button');
            button.style.position = 'fixed';
            button.style.opacity = 0;
            document.body.appendChild(button);
            button.focus();
            button.remove();
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
            this.pid = 0;
            this.toolbarVisible = this.selectedFile.length > 0;
        },

        // 移动完成操作
        handleConfirmMove(file) {
            this.moveVisible = false;
            this.toolbarVisible = this.selectedFile.length > 0;
            this.pid = 0;
            this.selectedFile = [];
            this.breadcrumb = [{id: 0, name: "文件"}]; // 重置面包屑导航
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
                            message: "移入回收站成功",
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

.file-list {
    display: flex;
    justify-content: space-between;
    width: 96%;
    height: auto;
    margin: 30px auto 0 auto;
    align-items: center;
    min-height: 10vh;
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

