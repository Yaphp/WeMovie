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
            <van-row class="profile">
                <van-col span="4" class="avatar">
                    <van-image round width="50" height="50" fit="cover" :src="shareUser.avatar"/>
                </van-col>
                <van-col span="12" class="username">
                    <div class="name">{{ shareUser.username }}</div>
                    <div class="expire">
                        {{ shareInfo.expire === 1 ? '30天有效' : '永久有效' }}
                    </div>
                </van-col>
                <van-col span="8" class="savebtn">
                    <van-button type="primary" size="small" round @click="handleSaveFile">保存文件</van-button>
                </van-col>
            </van-row>

            <van-divider></van-divider>

            <van-list
                v-model:loading="loading"
                :finished="true"
                finished-text="没有更多了"
                @load="shareFileList"
            >
                <div class="file-list">
                    <div class="file-dir-path" v-if="pid!==0">
                        <div class="file-dir-path-item" v-for="(bread,index) in breadcrumb" @click="changeDirectory(bread.id,index)">
                            {{ bread.name }}
                        </div>
                    </div>

                    <div class="file-item" v-longpress="{handler:handleLongPress,data:file}" v-for="file in shareFileList" @click="handleClickRow(file)">
                        <div class="file-left">
                            <van-checkbox v-if="hasMultChoose" class="van-checkbox-right-margin" v-model="file.checked" @click.stop="handleCheckFile(file)"/>
                            <van-image
                                class="file-icon"
                                lazy-load
                                fit="cover"
                                position="top"
                                :src="getFileIcon(file)"
                            >
                            </van-image>
                            <div class="file-name">{{ getShortFileName(file.name) }}</div>
                        </div>
                    </div>
                </div>
            </van-list>

            <div class="header-fill"></div>
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
import {getFileSize, getFileIcon, getShortFileName} from "utils/file.js";
import {fileList} from "api/file.js";
import FileSave from "components/m/FileSave.vue";


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
            getShortFileName,
            Picture
        }
    },
    data() {
        return {
            loading: true, // 控制加载动画是否显示
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
    computed: {
        hasMultChoose() {
            return this.selectedFile.length > 0;
        },
        hasChoosedCount() {
            return this.selectedFile.length;
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
    methods: {
        // 跳转登录
        handleLogin() {
            this.$router.push({
                'path': '/',
                'query': {
                    'redirect': location.href
                }
            })
        },

        // 长按事件 将当前文件设为选中文件
        handleLongPress(file) {
            file.checked = true;
            this.selectedFile.push(file);
        },

        // 获取分享文件基础信息
        handleGetShareFile() {
            // console.log(this.$route.query)
            this.shareFileList = [];
            this.selectedFile = [];

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

        // 监听文件选中状态
        handleCheckFile(file) {
            if (file.checked) {
                this.selectedFile.push(file);
            } else {
                this.selectedFile = this.selectedFile.filter(item => item.id !== file.id);
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

.el-main-fix {
    width: 100%;
    margin: auto;
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

.file-dir-path {
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

.file-dir-path-item:not(:first-child):before {
    content: " > ";
    color: rgba(191, 194, 199, 0.69);
}

.file-dir-path-item {
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

.header-fill {
    width: 100%;
    height: 50px;
}

.profile {
    padding-top: 20px
}

.avatar {
    display: flex;
    justify-content: center;
    align-items: center;
}

.username {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    font-size: 22px;
    font-weight: bold;
}

.name {
    width: 100%;
}

.expire {
    width: 100%;
    font-size: 12px;
    color: #999;
}

.van-row {
    flex-wrap: unset;
}

.van-checkbox-right-margin {
    margin-right: 10px;
}

.savebtn {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>