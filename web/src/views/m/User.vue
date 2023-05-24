<template>
    <div class="van-container">
        <van-row class="profile">
            <van-col span="8" class="avatar">
                <van-image round width="80" height="80" fit="cover" :src="this.$store.state.user && this.$store.state.user.avatar"/>
            </van-col>
            <van-col span="16" class="username">
                <div class="name">{{ this.$store.state.user && this.$store.state.user.username }}</div>
                <div class="memory">
                    存储总量: {{ getFileSize(this.$store.state.memory) || '0KB' }}
                </div>
            </van-col>
        </van-row>

        <van-cell-group class="cell-group-top" inset>
            <van-cell icon="manager" title="个人资料" is-link to="profile"/>
            <van-cell icon="like" title="收藏夹" is-link to="favorite"/>
            <van-cell icon="delete" title="回收站" is-link to="recycle"/>
        </van-cell-group>

        <van-cell-group inset>
            <van-cell icon="clear" title="退出登录" @click="logout" is-link to="/"/>
        </van-cell-group>
    </div>
</template>

<script>
import {getFileSize} from "utils/file.js";

export default {
    setup() {
        return {
            getFileSize
        }
    },
    data() {
        return {}
    },
    mounted() {
        this.init()
    },
    methods: {
        init() {
            console.log(!this.$store.state.user)
            if (!this.$store.state.user) {
                window.location.href = '#/login'
            }
        },
        logout() {
            window.sessionStorage.clear();
        }
    },
}
</script>

<style scoped>
.van-container {
    width: 100vw;
    height: 100vh;
    background-color: #f5f5f6;
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

.memory {
    width: 100%;
    font-size: 12px;
    color: #999;
}

.cell-group-top {
    margin-top: 40px;
    margin-bottom: 20px
}


</style>