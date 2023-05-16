<template>
    <van-row>
        <van-col span="2" class="back-button" @click="this.$router.back()">
            <van-icon name="arrow-left" size="16px" color="#999"/>
        </van-col>
        <van-col span="22">
            <van-search v-model="keyword" placeholder="输入关键字以搜索"
                        show-action
                        @keyup.enter="handleSearch()"
                        @search="handleSearch">
                <template #action>
                    <div @click="handleSearch">搜索</div>
                </template>
            </van-search>
        </van-col>
    </van-row>

    <van-tabs v-model:active="activeName" animated swipeable @change="handleSearchRequest">
        <van-tab v-for="ft in filetypes" :title="ft.label" :name="ft.value">
            <van-list
                v-model:loading="loading"
                :finished="true"
                finished-text="没有更多了"
                @load="handleSearchRequest"
            >
                <div class="file-list">
                    <div class="file-item" v-for="file in fileList" @click="handleClickRow(file)">
                        <div class="file-left">
                            <img class="file-icon" :src="getFileIcon(file)" :alt="file.name">
                            <div class="file-name">
                                {{ getShortFileName(file.name) }}
                            </div>
                        </div>
                        <div class="file-right"></div>
                    </div>
                </div>
            </van-list>
        </van-tab>
    </van-tabs>

    <van-overlay :show="previewImageShow" @click="previewImageShow = false">
        <div class="wrapper" @click="previewImageShow = false">
            <video class="wrapper-video" :src="previewImage" autoplay controls></video>
        </div>
    </van-overlay>
</template>

<script>
import {fileList} from "api/file.js";
import {showImagePreview} from 'vant';
import {getFileIcon, getFileSize, getShortFileName} from "utils/file.js";


export default {
    name: "Result",
    components: {},
    setup() {
        return {
            getFileIcon,
            getFileSize,
            getShortFileName
        }
    },
    data() {
        return {
            activeName: "all", // 文件类型
            keyword: this.$route.query.keyword,
            history: [], // 搜索历史
            filetypes: [
                {
                    "label": "全部",
                    "value": "all"
                }, {
                    "label": "文件夹",
                    "value": "folder"
                }, {
                    "label": "视频",
                    "value": "video"
                }, {
                    "label": "图片",
                    "value": "image"
                }, {
                    "label": "音频",
                    "value": "audio"
                }
            ], // 文件类型集合
            loading: false, // 是否正在加载
            fileList: [], // 文件列表
            previewImage: [], // 当前预览图片地址
            previewImageShow: false, // 控制图片预览是否显示
        }
    },
    mounted() {
        this.handleSearchRequest()
    },
    methods: {
        handleClickRow(row) {
            // 如果是文件夹，跳转到文件夹页面
            if (row.type === "folder") {
                this.$router.push({
                    path: '/files',
                    query: {
                        pid: row.id,
                    }
                })
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

        handleSearch() {
            // 保存搜索历史
            this.handleSetSearchHistory()

            // 发起搜索请求
            this.handleSearchRequest(this.keyword)
        },

        // 发起搜索请求
        handleSearchRequest() {
            this.fileList = [];
            this.selectedFile = [];
            fileList({
                uid: this.$store.state.user.id,
                keyword: this.keyword,
                pid: "all",
                type: this.activeName === 'all' ? '' : this.activeName
            }).then(res => {
                // 加载动画关闭
                this.loading = false;
                //设置文件数据
                this.fileList = res.data;
                this.imgSrcList = res.data.filter(item => item.type.indexOf('image') !== -1).map(item => item.path);
            });
        },

        // 保存搜索历史
        handleSetSearchHistory() {
            let history = localStorage.getItem("searchHistory")
            if (history) {
                let historyArr = JSON.parse(history)
                historyArr.unshift(this.keyword) // 插入到数组头部
                localStorage.setItem("searchHistory", JSON.stringify(historyArr))
            } else {
                localStorage.setItem("searchHistory", JSON.stringify([this.keyword]))
            }
        }
    }
}
</script>

<style scoped>
.back-button {
    display: flex;
    justify-content: center;
    align-items: center;
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
}

.file-right {
    margin-right: 5px;
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