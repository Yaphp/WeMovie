<template>
    <el-container class="home-container">
        <el-container>
            <!-- 侧边栏 -->
            <el-aside width="15%" class="el-aside">
                <h2 class="index-title">{{ this.$store.state.app.name }}</h2>

                <el-menu
                    :default-active="$route.path"
                    background-color="#252626"
                    text-color="#fff"
                    text-align="left"
                    active-text-color="#ffd04b"
                    style="border:1px solid #252626;"
                >
                    <Menu :navMenus="menuData"></Menu>
                </el-menu>
            </el-aside>

            <el-main>
                <el-page-header :icon="ArrowLeft" @back="this.$router.push('/index')">
                    <template #content>
                        <span class="text-large font-600 mr-3"> {{ this.$store.state.tagName || '首页' }} </span>
                    </template>
                    <template #extra>
                        <div class="flex items-center">
                            <span class="text-large font-600 mr-3" style="margin-right: 10px;"> Welcome, {{ this.$store.state.user.username || '未登录' }} </span>
                            <a @click="logout" style="cursor: pointer;margin-right:30px;color:rgba(9,163,229,0.76)">注销</a>
                        </div>
                    </template>
                </el-page-header>

                <!--路由占位符-->
                <router-view></router-view>

                <el-footer class="el-footer-container">
                    <div class="footer">
                        <div class="footer-left">
                            <span>© 2022 {{ this.$store.state.app.name }}</span>
                        </div>
                    </div>
                </el-footer>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
import Menu from "../components/Menu.vue";

export default {
    name: "Home",
    components: {
        Menu: Menu
    },
    data() {
        return {
            activeIndex: '/index',
            menuData: [],
        };
    },
    mounted() {
        this.menuData = this.$store.state.user.role === 0 ? this.$store.state.adminMenuData : this.$store.state.userMenuData
        // console.log(window.location.href)
        let start = window.location.href.lastIndexOf('/');
        let path = window.location.href.slice(start + 1);
        this.activeIndex = path;
        this.$router.push(path)
    },
    methods: {
        logout() {
            this.$store.commit('REMOVE_INFO');
            this.$router.push('/');
        }
    }
};
</script>

<style>

.index-title {
    text-align: center;
    margin: 20px;
    color: white;
    font-size: 18px;
}

.el-aside {
    background-color: #252626;
    height: calc(100vh + 50px);
    border: 0;

}

.el-main {
    --el-main-padding: 20px auto;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    background-color: #f0f2f5;
    padding: 10px 5px;
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
</style>