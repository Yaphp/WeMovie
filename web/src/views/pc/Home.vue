<template>
    <el-container class="home-container">
        <el-container>
            <!-- 侧边栏 -->
            <el-aside width="18%" class="el-aside">

                <div class="title-logo">
                    <div>
                        <img class="app-logo" src="/images/logo.png">
                    </div>
                    <h2 class="app-title">
                        {{ this.$store.state.app.name }}
                    </h2>
                </div>

                <el-menu
                    :default-active="$route.path"
                    background-color="#f5f5f6"
                    text-color="#black"
                    text-align="left"
                >
                    <Menu :navMenus="menuData"></Menu>
                </el-menu>


                <div class="footer-box">
                    <div class="user-profile">
                        <div class="user-image-box" @click="jumpProfile">
                            <div class="user-avatar">
                                <img class="user-image" :src="this.$store.state.user.avatar">
                            </div>
                            <div class="user-name">
                                <p class="username">{{ this.$store.state.user.username }}</p>
                                <p class="memory">
                                    存储总量: {{ getFileSize(this.$store.state.memory) || '0KB' }}
                                </p>
                            </div>
                        </div>

                        <div class="logout-icon-box">
                            <img class="logout-icon" @click="logout" src="/images/quit.png">
                        </div>
                    </div>
                </div>

            </el-aside>

            <el-main class="el-main-fix">
                <!--路由占位符-->
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
import Menu from "components/pc/Menu.vue";
import {getFileSize} from "utils/file.js";

export default {
    name: "Home",
    components: {
        Menu: Menu
    },
    setup() {
        return {
            getFileSize
        }
    },
    data() {
        return {
            activeIndex: '/index',
            menuData: [],
        };
    },
    mounted() {
        switch (this.$store.state.user.role) {
            case 0:
                this.menuData = this.$store.state.rootMenuData;
                break;
            case 1:
                this.menuData = this.$store.state.userMenuData;
                break;
        }
        // console.log(window.location.href)
        let start = window.location.href.lastIndexOf('/');
        let path = window.location.href.slice(start + 1);
        this.activeIndex = path;
        this.$router.push(path)
        document.title = this.$store.state.app.name;
    },
    methods: {
        jumpProfile() {
            this.$router.push('/profile');
        },
        logout() {
            this.$store.commit('REMOVE_INFO');
            this.$router.push('/');
        }
    }
};
</script>

<style>

.title-logo {
    margin: 25px 0;
    display: flex;
    justify-content: center;
    align-items: center;
}

.app-logo {
    width: 30px;
    height: 30px;
    mix-blend-mode: multiply
}

.app-title {
    text-align: center;
    margin: 10px;
    color: black;
    font-size: 18px;
}

.el-main-fix {
    background-color: white;
    height: 100vh !important;
}

.el-aside {
    background-color: #f5f5f6;
    height: 100vh;
    border: 0;
}

.footer-box {
    position: absolute;
    bottom: 0;
    width: 18%;
}

.user-profile {
    display: flex;
    justify-content: space-around;
    align-items: center;
    width: 100%;
    height: 60px;
    cursor: pointer;
    border-top: 1px solid #ebeef5;
    background-color: rgba(132, 133, 141, 0.08);
}

.user-image-box {
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.user-name {
    margin-left: 10px;
    font-size: 14px;
}

.logout-icon {
    margin-left: 10px;
    width: 20px;
    height: 20px;
    cursor: pointer;
}

.user-image {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    overflow: hidden;
}

.username {
    font-size: 16px;
    font-weight: bold;
}

.memory {
    font-size: 12px;
    color: #999;
}

.el-main {
    --el-main-padding: 20px auto;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background-color: white;
    padding: 10px 5px;
    height: auto;
}

.el-footer-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    position: fixed;
    bottom: 0;
    text-align: left;
    width: 100%;
}

.el-page-header__header {
    margin-left: 10px;
    margin-bottom: 10px;
}

.el-menu {
    border-right: 0;
}
</style>