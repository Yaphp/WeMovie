<template>
    <el-container>
        <el-header class="nav-header">
            <div class="left">
                <img src="/images/logo.png" class="logo">
                {{ this.$store.state.app.name }}
            </div>
            <div class="right">
                <el-dropdown v-if="this.$store.state.user">
                    <el-avatar>
                        <img :src="this.$store.state.user.avatar">
                    </el-avatar>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item :icon="HomeFilled" @click="handleDirectHome">主页</el-dropdown-item>
                            <el-dropdown-item :icon="SwitchButton" @click="handleLogout">注销</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
                <el-button v-else type="primary" @click="handleLogin">登录</el-button>
            </div>
        </el-header>

        <el-main v-if="!shareFileList" class="el-main-fix">
            <el-result
                v-if="errCode === 3"
                icon="error"
                title="共享文件不存在"
                sub-title="请检查共享链接是否正确"
            >
            </el-result>
            <el-result
                v-if="errCode === 2"
                icon="warning"
                title="共享文件已过期"
                sub-title="请联系文件分享者重新分享"
            >
            </el-result>
            <el-result
                v-if="errCode === 1"
                icon="info"
                title="提取码不正确"
                sub-title="请检查密码是否正确"
            ></el-result>
        </el-main>

        <el-main v-else class="el-main-fix">
            <el-col class="avatar-username-row">
                <div class="left-box">
                    <div class="avatar-box">
                        <el-avatar>
                            <img :src="shareUser.avatar">
                        </el-avatar>
                    </div>
                    <div class="username">
                        <div class="username-up">
                            <div class="file-dir">
                                <div class="file-dir-item" v-for="(bread,index) in breadcrumb"
                                     @click="changeDirectory(bread.id,index)">
                                    <span :class="(index+1) === breadcrumb.length ? 'active-path':'normal-path'">
                                    {{ bread.name }}
                                    </span>
                                </div>
                            </div>
                        </div>
                        <div class="username-down">
                            <span>{{ shareUser.username }} · {{
                                    shareInfo.expire === 1 ? '30天有效' : '永久有效'
                                }}</span>
                        </div>
                    </div>
                </div>
                <div class="right-box">
                    <el-button type="primary" @click="handleSaveFile">保存到我的云盘</el-button>
                </div>
            </el-col>

            <el-table
                :data="shareFileList"
                @select="selectedRow"
                @select-all="selectedRow"
                @row-click="handleClickRow"
                style="width: 100%"
            >
                <el-table-column type="selection" width="50"/>
                <el-table-column prop="name" label="名称">
                    <template #default="{ row }">
                        <div class="file-name-icon">
                            <div class="file-icon">
                                <el-image loading="lazy" scroll-container=".el-scrollbar__wrap" :key="row.id"
                                          :src="getFileIcon(row)" class="file-icon" fit="cover">
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
                    </template>
                </el-table-column>
                <el-table-column prop="created_at" label="创建时间"/>
                <el-table-column prop="size" label="文件大小">
                    <template #default="{ row }">
                        <div class="file-size">
                            {{ getFileSize(row.size) }}
                        </div>
                    </template>
                </el-table-column>
            </el-table>
        </el-main>
    </el-container>


    <FileSave
        v-if="moveVisible"
        :show="moveVisible"
        :file="selectedFile"
        @confirm="handleSaveFile"
        @close="handleCancelMove">
    </FileSave>
</template>

<script>
import {HomeFilled, SwitchButton, Picture} from "@element-plus/icons-vue";
import {shareList} from "api/share.js";
import {getFileSize, getFileIcon} from "utils/file.js";
import {fileList} from "api/file.js";
import FileSave from "components/pc/FileSave.vue";


export default {
    name: "Share",
    components: {
        Picture,
        FileSave
    },
    setup() {
        return {
            HomeFilled,
            SwitchButton,
            getFileSize,
            getFileIcon,
            Picture
        }
    },
    data() {
        return {
            shareUser: {},
            shareInfo: {},
            shareFileList: null,
            errCode: null,
            errMsg: null,
            pid: 0,
            breadcrumb: [
                {
                    id: 0,
                    name: "分享文件",
                }
            ], // 面包屑
            selectedFile: [], // 选中的文件
            moveVisible: false, // 移动弹窗
        }
    },
    mounted() {
        this.handleGetShareFile()
    },
    methods: {
        handleLogin() {
            this.$router.push({
                'path': '/',
                'query': {
                    'redirect': location.href
                }
            })
        },

        // 获取分享文件基础信息
        handleGetShareFile() {
            console.log(this.$route.query)
            shareList(this.$route.query).then(res => {
                if (res.code === 0) {
                    this.shareUser = res.data.user;
                    this.shareInfo = res.data.share;
                    this.shareFileList = res.data.files;
                } else {
                    this.errCode = res.code;
                    this.errMsg = res.msg;
                }
            }).catch(error => {
                console.log(error)
            })
        },

        handleDirectHome() {
            this.$router.push({
                'path': '/files'
            })
        },

        handleLogout() {
            this.$store.commit('REMOVE_INFO');
            this.$router.push('/');
        },

        handleSaveFile() {
            if (!this.$store.getters.getUser) {
                this.$message.warning('请先登录');
                return;
            }
            if (this.selectedFile.length === 0) {
                this.$message.warning('请选择文件');
                return;
            }
            this.moveVisible = true;
        },

        handleCancelMove() {
            this.moveVisible = false;
        },

        handleGetFileList() {
            // 获取文件列表
            this.fileList = [];
            fileList({
                pid: this.pid,
                uid: this.shareUser.id,
            }).then(res => {
                this.shareFileList = res.data;
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
            if (id === 0) {
                this.handleGetShareFile();
                return;
            }
            // 修改当前路径
            this.pid = id;
            // 加载文件列表
            this.handleGetFileList();
        },

        // 选中行
        selectedRow(selection, row) {
            // console.log(selection, row)
            this.selectedFile = selection;
        }
    }
}
</script>

<style>
.nav-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    background-color: #f5f5f6;
}

.nav-header .left {
    display: flex;
    align-items: center;
}

.nav-header .left > img {
    width: 50px;
    height: 50px;
    mix-blend-mode: multiply;
}

.el-table__row:hover {
    cursor: pointer;
}

.el-main-fix {
    width: 60%;
    margin: auto;
}

.avatar-username-row {
    display: flex;
    align-items: center;
    justify-content: space-between;

}

.left-box {
    display: flex;
    /*align-items: center;*/
    justify-content: flex-start;
}

.file-name-icon {
    display: flex;
    align-items: center;
    /*justify-content: center;*/
}


.el-checkbox__inner {
    --el-checkbox-border-radius: 50% !important;
    border-radius: 50% !important;
}

.file-icon {
    width: 30px;
    height: 30px;
    margin-right: 20px;
    object-fit: cover;
}

.username {
    text-align: left;
    margin-left: 10px;
}

.username-down {
    font-size: 12px;
    color: #909399;
}

.file-dir {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    flex-wrap: wrap;
    align-items: center;
}

.file-dir-item:not(:first-child):before {
    content: " > ";
    color: rgba(191, 194, 199, 0.69);
}

.file-dir-item {
    margin-left: 5px;
    font-size: 16px;
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
</style>